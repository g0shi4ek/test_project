package main

import (
	"context"
	"log"

	"github.com/g0shi4ek/test_project/config"
	"github.com/g0shi4ek/test_project/internal/app"
	"github.com/gin-gonic/gin"
)

func main() {

	ctx := context.Background()

	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Cannot load configuraton, %v", err)
		return
	}

	err = app.NewAuthServer(ctx, cfg)
	if err != nil{
		log.Fatalf("Cannot start gRPC auth server, %v", err)
	}

	startService := gin.Default()
	


}
