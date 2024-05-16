package controller

import (
	"errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"github.com/mystpen/Pet-API/config"
	"github.com/mystpen/Pet-API/internal/dto"
	"github.com/mystpen/Pet-API/internal/model"
	"github.com/mystpen/Pet-API/internal/validator"
)

type userService interface {
	RegisterUser(*dto.RegistrationRequest) error
	GetRegisteredUser(*dto.LogInRequest) (*model.User, error)
	CreateToken(*model.User) string
}

type Controller struct {
	service userService
	redis   *redis.Client
}

func NewController(service userService, redisClient *redis.Client) *Controller {
	return &Controller{
		service: service,
		redis:   redisClient,
	}
}

func (c *Controller) Routes(r *gin.Engine, cfg *config.Config) {
	r.POST("signup", c.Signup)
	r.POST("signin", c.Signin)
}

func (c *Controller) Signup(ctx *gin.Context) {
	var request dto.RegistrationRequest
	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	validate := validator.New()
	err = validate.Struct(request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = c.service.RegisterUser(&request)
	if err != nil {
		if errors.Is(err, model.ErrDuplicateEmail) {
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": "a user with this email address already exists"})
			return
		} else {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()}) // TODO: add some err
			return
		}
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "User successful registered"})
}

func (c *Controller) Signin(ctx *gin.Context) {
	var request dto.LogInRequest
	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	validate := validator.New()
	err = validate.Struct(request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := c.service.GetRegisteredUser(&request)
	if err != nil {
		switch {
		case errors.Is(err, model.ErrRecordNotFound):
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "incorrect email or password"})
			return
		default:
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	}

	// create token session
	token := c.service.CreateToken(user)

	// set redis with time limit
	c.redis.Set(token, 1, time.Hour)

	// header add
	ctx.Writer.Header().Add("Authorization", "Basic "+token)
}
