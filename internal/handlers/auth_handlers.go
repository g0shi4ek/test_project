package handlers

import (
	"log"
	"net/http"

	pb "github.com/g0shi4ek/test_project/gen/go"
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context, client pb.AuthClient) string {
	var input struct{ Email, Password string }
	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request format"})
		return ""
	}
	log.Printf("Login User %s, %s", input.Email, input.Password)

	token, err := client.LoginUser(c.Request.Context(), &pb.LoginRequest{
		Email:    input.Email,
		Password: input.Password,
	})
	if err != nil {
		log.Printf("server error %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input data"})
		return ""
	}

	c.Header("Authorization", token.Token)
	c.JSON(http.StatusOK, gin.H{"token": token})
	return token.Token
}

func Register(c *gin.Context, client pb.AuthClient) {
	var input struct{ Email, Password string }
	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request format"})
		return
	}
	log.Printf("Register New User %s, %s", input.Email, input.Password)

	userId, err := client.RegisterUser(c.Request.Context(), &pb.RegisterRequest{
		Email:    input.Email,
		Password: input.Password,
	})
	if err != nil {
		log.Printf("server error %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input data"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"uuid": userId})
}
