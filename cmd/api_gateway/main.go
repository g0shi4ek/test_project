package main

import (
	"context"
	"log"
	"os/signal"
	"syscall"

	"github.com/g0shi4ek/test_project/cmd/api_gateway/cache"
	"github.com/g0shi4ek/test_project/cmd/clients"
	"github.com/g0shi4ek/test_project/config"
	"github.com/g0shi4ek/test_project/internal/app"
	"github.com/g0shi4ek/test_project/pkg/middleware"
	"github.com/g0shi4ek/test_project/pkg/redis"
	"github.com/gin-gonic/gin"
)

func main() {
	ctx, stop := signal.NotifyContext(
		context.Background(),
		syscall.SIGTERM,
		syscall.SIGQUIT,
	)
	defer stop()

	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Cannot load configuraton, %v", err)
		return
	}

	redis, err := redis.ConnectRedis(context.Background(), cfg)
	if err != nil{
		log.Printf("Failed to create redis client, %v", err)
		return 
	}
	tokenCache := cache.NewTokenRepository(redis)

	go func() {
		if err := app.NewAuthServer(redis, ctx, cfg); err != nil {
			log.Fatalf("Cannot start gRPC auth server, %v", err)
		}
	}()

	client, err := clients.NewAuthClient(tokenCache, ctx, cfg)
	if err != nil {
		log.Fatalf("Cannot start gRPC auth client, %v", err)
	}

	router := gin.Default()

	api := router.Group("/api")
	api.POST("/register", client.RegisterUser)
	api.POST("/login", client.LoginUser)
	api.Use(middleware.AuthMiddleware(ctx, client))

	go func() {
		if err := router.Run(":" + cfg.ApiPort); err != nil {
			log.Fatalf("Cannot start gin server, %v", err)
		}
	}()

	<-ctx.Done()

	log.Println("Shutting Down Servers")
}
