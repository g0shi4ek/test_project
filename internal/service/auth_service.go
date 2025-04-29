package service

import (
	"context"
	"fmt"
	"log"

	"github.com/g0shi4ek/test_project/config"
	pb "github.com/g0shi4ek/test_project/gen/go"
	"github.com/g0shi4ek/test_project/internal/models"
	"github.com/g0shi4ek/test_project/internal/repository"
	"github.com/g0shi4ek/test_project/pkg/jwt"
)

type AuthService struct {
	pb.UnimplementedAuthServer
	UserRepo  repository.UserRepository
	Config    *config.Config
}

func NewAuthService(userRepo repository.UserRepository, cfg *config.Config) *AuthService {
	return &AuthService{
		UnimplementedAuthServer: pb.UnimplementedAuthServer{},
		UserRepo:                userRepo,
		Config:                  cfg,
	}
}

func (s *AuthService) RegisterUser(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResp, error) {
	hash, err := HashPassword(req.GetPassword())
	if err != nil {
		return nil, err
	}
	user := models.User{
		UserUUID:     "",
		Email:        req.GetEmail(),
		PasswordHash: hash,
	}
	log.Println("New request", req.Email)
	if err := s.UserRepo.CreateUser(ctx, &user); err != nil {
		log.Printf("Cannot create user, %v", err)
		return nil, fmt.Errorf("cannot create user, %w", err)
	}

	return &pb.RegisterResp{
		UserId: user.UserUUID,
	}, nil
}

func (s *AuthService) LoginUser(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResp, error) {

	email, password := req.GetEmail(), req.GetPassword()
	log.Println("New request", req.Email)
	newUser, err := s.UserRepo.GetUserByEmail(ctx, email)
	if err != nil {
		log.Printf("Cannot find user with email %s, %v", email, err)
		return nil, fmt.Errorf("cannot find user, %w", err)
	}

	validPassword := newUser.PasswordHash

	if !CheckPassword(validPassword, password) {
		log.Printf("Wrong password")
		return nil, fmt.Errorf("wrong password %s", password)
	}

	jwtToken, err := jwt.CreateNewToken(email, s.Config)
	if err != nil {
		log.Printf("Cannot create token")
		return nil, fmt.Errorf("cannot create token %w", err)
	}

	return &pb.LoginResp{
		Token: jwtToken,
	}, nil
}
