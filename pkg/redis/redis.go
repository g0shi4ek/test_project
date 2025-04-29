package redis

import (
	"context"
	"log"
	"strconv"

	"github.com/g0shi4ek/test_project/config"
	"github.com/go-redis/redis"
)

func ConnectRedis(ctx context.Context, cfg *config.Config) (*redis.Client, error) {
	db, _ := strconv.Atoi(cfg.RedisName)
	client := redis.NewClient(&redis.Options{
		Addr: "localhost:" + cfg.RedisPort,
		DB:   db,
	})

	_, err := client.Ping().Result()
	if err != nil {
		log.Println("Cannot connected to redis")
		return nil, err
	}
	log.Println("Connected to redis")
	return client, nil
}
