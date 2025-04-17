package models

import (
	pb "github.com/g0shi4ek/test_project/gen/go/auth_grpc.pb.go"
)

type AuthRequest struct {
	email    string
	password string
}

type AuthResp struct {
	userId string
}

func NewAuthRequest(req *pb.RegisterRequest) *AuthRequest {
	return &AuthRequest{
		email:    req.Email,
		password: req.Password,
	}
}

// брать юзера?

type User struct {
	userUUID string
	email    string
	password string
}
