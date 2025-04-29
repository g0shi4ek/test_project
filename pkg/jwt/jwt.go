package jwt

import (
	"time"

	"github.com/g0shi4ek/test_project/config"
	"github.com/golang-jwt/jwt"
)

func getRole(email string) string {
	if email == "admin@gmail.com" {
		return "admin"
	}
	return "user"
}

func CreateNewToken(email string, cfg *config.Config) (string, error) {
	key := []byte(cfg.SecretKey)

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": email,                            // Subject (user identifier)
		"iss": "app",                            // Issuer
		"aud": getRole(email),                   // Audience (user role)
		"exp": time.Now().Add(time.Hour).Unix(), // Expiration time
		"iat": time.Now().Unix(),                // Issued at
	})

	tokenString, err := claims.SignedString(key)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
