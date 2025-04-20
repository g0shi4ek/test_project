package app

import (
	"context"
	"log"
	"net"

	"github.com/g0shi4ek/test_project/config"
	pb "github.com/g0shi4ek/test_project/gen/go"
	"github.com/g0shi4ek/test_project/internal/adapter/database"
	"github.com/g0shi4ek/test_project/internal/repository"
	"github.com/g0shi4ek/test_project/internal/service"
	"google.golang.org/grpc"
)

func NewAuthServer(ctx context.Context, cfg *config.Config) error{
	listener, err := net.Listen("tcp", cfg.AuthPort)
	if err != nil {
		log.Printf("failed to listen, %v", err)
		return err
	}
	grcpServer := grpc.NewServer()

	db, err := database.NewUserPool(context.Background(), cfg)
	if err != nil {
		log.Printf("failed to create pool, %v", err)
		return err
	}
	defer db.Close()

	authRepo := repository.NewUserRepository(db)
	authService := service.NewAuthUseCases(pb.UnimplementedAuthServer{}, authRepo, cfg)

	pb.RegisterAuthServer(grcpServer, &authService.Service)

	err = grcpServer.Serve(listener)
	if err != nil {
		log.Printf("Error starting server, %v", err)
		return err
	}
	log.Printf("gRPC server start, port: %s", cfg.AuthPort)
	return nil
}
