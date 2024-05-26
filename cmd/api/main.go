package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"

	"github.com/mystpen/Pet-API/config"
	"github.com/mystpen/Pet-API/internal/controller"
	"github.com/mystpen/Pet-API/internal/redis"
	"github.com/mystpen/Pet-API/internal/repository"
	"github.com/mystpen/Pet-API/internal/repository/user"
	"github.com/mystpen/Pet-API/internal/service"
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

	repo := user.NewUserRepository(db)
	service := service.NewUserService(repo)
	controller := controller.NewController(service, redisClient)

	controller.Routes(server, cfg)

	err = server.Run(fmt.Sprintf(":%v", cfg.Port))
	if err != nil {
		log.Fatal(err)
	}
}
