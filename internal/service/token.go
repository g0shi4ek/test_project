package service

import (
	"time"

	"github.com/g0shi4ek/test_project/config"
	"github.com/g0shi4ek/test_project/internal/models"
	"github.com/golang-jwt/jwt"
)

func getRole(user models.User) string {
	if user.Email == "admin@gmail.com" {
		return "admin"
	}
	return "user"
}

func CreateNewToken(user models.User, cfg *config.Config) (string, error) {
	key := []byte(cfg.SecretKey)

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.Email,                       // Subject (user identifier)
		"iss": "app",                            // Issuer
		"aud": getRole(user),                    // Audience (user role)
		"exp": time.Now().Add(time.Hour).Unix(), // Expiration time
		"iat": time.Now().Unix(),                // Issued at
	})

	tokenString, err := claims.SignedString(key)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
