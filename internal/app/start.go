package app

import (
	"log"
	"net"

	pb "github.com/g0shi4ek/test_project/gen/go/auth_grpc.pb.go"
	"github.com/g0shi4ek/test_project/internal/service"
	"google.golang.org/grpc"
)


func NewAuthServer(){
	listener, err := net.Listen("tcp", ":50051")
	if err != nil{
		log.Print("failed to listen", err)
	}
	grcpServer := grpc.NewServer()
	server := service.Authservice{}

	pb.RegisterAuthServer(grcpServer, server)

	err = grcpServer.Serve(listener)
	if err != nil{
		log.Fatalf("Error starting server", err)
	}
}

