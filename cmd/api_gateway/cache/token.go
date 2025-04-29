package cache

import (
	"time"

	"github.com/go-redis/redis"
)

type TokenRepository interface {
	SetToken(string) error
	IsValidToken(string) (bool, error)
}

type TokenRepo struct {
	client *redis.Client
}

func NewTokenRepository(r *redis.Client) TokenRepository {
	return &TokenRepo{
		client: r,
	}
}

func (r *TokenRepo) SetToken(token string) error {
	return r.client.Set("valid_token:" + token,"true", time.Hour).Err()
}

func (r *TokenRepo) IsValidToken(token string) (bool, error) {
	exists, err := r.client.Exists("valid_tokens:"+token).Result()
    return exists == 1, err
}
