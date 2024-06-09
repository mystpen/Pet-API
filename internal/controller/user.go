package controller

import (
	"errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"github.com/mystpen/Pet-API/internal/dto"
	"github.com/mystpen/Pet-API/internal/model"
	"github.com/mystpen/Pet-API/internal/validator"
)

type userService interface {
	RegisterUser(*dto.RegistrationRequest) error
	GetRegisteredUser(*dto.LogInRequest) (*model.User, error)
	CreateToken(*model.User) string
}

type docService interface {
}

type Controller struct {
	userService userService
	docService  docService
	redis       *redis.Client
}

func NewController(userservice userService, docervice docService, redisClient *redis.Client) *Controller {
	return &Controller{
		userService: userservice,
		docService:  docervice,
		redis:       redisClient,
	}
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
	err = c.userService.RegisterUser(&request)
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

	user, err := c.userService.GetRegisteredUser(&request)
	if err != nil {
		switch {
		case errors.Is(err, model.ErrRecordNotFound), errors.Is(err, model.ErrNoMatch):
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "incorrect email or password"})
			return
		default:
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	}

	// create token session
	token := c.userService.CreateToken(user)

	// set redis with time limit
	c.redis.Set(token, 1, time.Minute)

	ctx.JSON(http.StatusOK, gin.H{"message": "User successful logged in"})
}

func (c *Controller) TestHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "Test auth"})
}
