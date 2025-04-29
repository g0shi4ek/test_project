package app

import (
	"context"
	"log"
	"net"

	"github.com/g0shi4ek/test_project/config"
	pb "github.com/g0shi4ek/test_project/gen/go"
	"github.com/g0shi4ek/test_project/internal/repository"
	"github.com/g0shi4ek/test_project/internal/service"
	"github.com/g0shi4ek/test_project/pkg/database"
	"github.com/go-redis/redis"
	"google.golang.org/grpc"
)

func NewAuthServer(cache * redis.Client,  ctx context.Context, cfg *config.Config) error {
	listener, err := net.Listen("tcp", ":"+cfg.AuthPort)
	if err != nil {
		log.Printf("Failed to listen, %v", err)
		return err
	}
	grpcServer := grpc.NewServer()

	db, err := database.NewUserPool(context.Background(), cfg)
	if err != nil {
		log.Printf("Failed to create pool, %v", err)
		return err
	}
	defer db.Close()

	authRepo := repository.NewUserRepository(db)
	authService := service.NewAuthService(authRepo, cfg)

	pb.RegisterAuthServer(grpcServer, authService)

	log.Printf("Server is listening on port %s...", cfg.AuthPort)
	if err := grpcServer.Serve(listener); err != nil {
		log.Printf("Server dropped with error: %v", err)
	}

	return nil
}
