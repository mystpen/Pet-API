package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"

	"github.com/mystpen/Pet-API/config"
	"github.com/mystpen/Pet-API/internal/controller"
	"github.com/mystpen/Pet-API/internal/redis"
	"github.com/mystpen/Pet-API/internal/repository"
	docrepo "github.com/mystpen/Pet-API/internal/repository/document"
	"github.com/mystpen/Pet-API/internal/repository/user"
	"github.com/mystpen/Pet-API/internal/service"
	docservice "github.com/mystpen/Pet-API/internal/service/document"
	"github.com/mystpen/Pet-API/pkg"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Println(err, nil)
	}

	// Connect to DB
	db, err := pkg.OpenDB(*cfg)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	redisClient, err := redis.NewRedisClient(cfg)
	if err != nil {
		log.Println(err)
		return
	}

	server := gin.Default()

	err = repository.Init(db)
	if err != nil {
		log.Println(err)
		return
	}

	userRepo := user.NewUserRepository(db)
	docRepo := docrepo.NewDocumentRepository(db)
	userService := service.NewUserService(userRepo)
	docService := docservice.NewDocumentService(docRepo)

	controller := controller.NewController(userService, docService, redisClient)

	controller.Routes(server, cfg)

	err = server.Run(fmt.Sprintf(":%v", cfg.Port))
	if err != nil {
		log.Fatal(err)
	}
}
