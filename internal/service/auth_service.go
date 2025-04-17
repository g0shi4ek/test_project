package service

import (
	"context"
	"log"

	pb "github.com/g0shi4ek/test_project/gen/go/auth_grpc.pb.go"
	"github.com/g0shi4ek/test_project/internal/models"
)

type Authservice struct {
	pb.UnimplementedAuthServer
}

func (s *Authservice) RegisterUser(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResp, error) {
	newUser := models.NewAuthRequest(req)
	log.Println(newUser.email)
	return &pb.RegisterResp{
		UserId: "1",
	}, nil
}

func (s *Authservice) LoginUser(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResp, error) {
	newUser := models.NewAuthRequest(req)
	log.Println(newUser.password)
	return &pb.RegisterResp{
		Token: "f6rg;kjsfn",
	}, nil
}
