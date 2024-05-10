package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/mystpen/Pet-API/config"
)

type userService interface{

}

type Controller struct{
	service userService
}

func NewController(service userService) *Controller{
	return &Controller{service: service}
}

func (c *Controller) Routes(r *gin.Engine, cfg *config.Config){
	
}