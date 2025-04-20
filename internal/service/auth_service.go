package service

import (
	"context"
	"fmt"
	"log"

	"github.com/g0shi4ek/test_project/config"
	pb "github.com/g0shi4ek/test_project/gen/go"
	"github.com/g0shi4ek/test_project/internal/models"
	"github.com/g0shi4ek/test_project/internal/repository"
)

type AuthUseCases struct {
	Service  pb.UnimplementedAuthServer
	UserRepo repository.UserRepository
	Config   *config.Config
}

func NewAuthUseCases(service pb.UnimplementedAuthServer, userRepo repository.UserRepository, cfg *config.Config) *AuthUseCases {
	return &AuthUseCases{
		Service:  service,
		UserRepo: userRepo,
		Config:   cfg,
	}
}

func (s *AuthUseCases) RegisterUser(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResp, error) {
	user := models.User{
		UserUUID: "",
		Email:    req.GetEmail(),
		Password: req.GetPassword(),
	}
	log.Println("New request", req.Email)
	err := s.UserRepo.CreateUser(ctx, user)
	if err != nil {
		log.Printf("Cannot create user, %v", err)
		return nil, fmt.Errorf("cannot create user, %w", err)
	}

	return &pb.RegisterResp{
		UserId: user.UserUUID,
	}, nil
}

func (s *AuthUseCases) LoginUser(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResp, error) {
	user := models.User{
		UserUUID: "",
		Email:    req.GetEmail(),
		Password: req.GetPassword(),
	}
	log.Println("New request", req.Email)
	newUser, err := s.UserRepo.GetUserByEmail(ctx, user.Email)
	if err != nil {
		log.Printf("Cannot find user with email %s, %v", user.Email, err)
		return nil, fmt.Errorf("cannot find user, %w", err)
	}

	if newUser.Password != user.Password {
		log.Printf("Wrong password")
		return nil, fmt.Errorf("wrong password %s", user.Password)
	}

	jwtToken, err := CreateNewToken(user, s.Config)
	if err != nil {
		log.Printf("Cannot create token")
		return nil, fmt.Errorf("cannot create token %w", err)
	}

	return &pb.LoginResp{
		Token: jwtToken,
	}, nil
}
