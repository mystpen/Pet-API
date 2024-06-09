package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/mystpen/Pet-API/config"
)

func (c *Controller) Routes(r *gin.Engine, cfg *config.Config) {
	users := r.Group("/user")
	users.POST("signup", c.Signup)
	users.POST("signin", c.Signin)

	doc := r.Group("/document")
	doc.GET("test", c.BasicAuthMiddleware(), c.TestHandler)
}
