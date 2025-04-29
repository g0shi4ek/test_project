package clients

import (
	"context"
	"log"

	"github.com/g0shi4ek/test_project/cmd/api_gateway/cache"
	"github.com/g0shi4ek/test_project/config"
	pb "github.com/g0shi4ek/test_project/gen/go"
	"github.com/g0shi4ek/test_project/internal/handlers"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type AuthClient struct {
	Client pb.AuthClient
	TokenRepo cache.TokenRepository
}

func NewAuthClient(tokenRepo cache.TokenRepository,  ctx context.Context, cfg * config.Config) (*AuthClient, error){
	cl, err := NewAuthpbClient(ctx, cfg)
	if err != nil{
		return nil, grpc.ErrServerStopped
	}
	return &AuthClient{
		Client: cl,
		TokenRepo: tokenRepo, 
	}, nil
}

func NewAuthpbClient(ctx context.Context, cfg * config.Config) (pb.AuthClient, error) {
	conn, err := grpc.NewClient(
		"localhost:" + cfg.AuthPort,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	log.Printf("Start grpc client on %s", cfg.AuthPort)
	return pb.NewAuthClient(conn), nil
}

func (cl *AuthClient) LoginUser(c *gin.Context) {
	token := handlers.Login(c, cl.Client)
	if token != ""{
		cl.TokenRepo.SetToken(token)
	}
}

func (cl *AuthClient) RegisterUser(c *gin.Context) {
	handlers.Register(c, cl.Client)
}
