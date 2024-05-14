package controller

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mystpen/Pet-API/config"
	"github.com/mystpen/Pet-API/internal/dto"
	"github.com/mystpen/Pet-API/internal/model"
	"github.com/mystpen/Pet-API/internal/validator"
)

type userService interface{
	RegisterUser(*dto.RegistrationRequest) error
}

type Controller struct{
	service userService
}

func NewController(service userService) *Controller{
	return &Controller{service: service}
}

func (c *Controller) Routes(r *gin.Engine, cfg *config.Config){
	r.POST("signup", c.Signup)
	r.POST("signin", c.Signin)
}

func (c *Controller) Signup(ctx *gin.Context){
	var request dto.RegistrationRequest
	err := ctx.ShouldBindJSON(&request)
	if err != nil{
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	validate := validator.New()
	err = validate.Struct(request)
	if err != nil{
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = c.service.RegisterUser(&request)
	if err != nil{
		if errors.Is(err, model.ErrDuplicateEmail){
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": "a user with this email address already exists"})
			return
		} else{
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()}) //TODO: add some err
			return
		}
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "User successful registered"})
}

func (c *Controller) Signin(ctx *gin.Context){

}