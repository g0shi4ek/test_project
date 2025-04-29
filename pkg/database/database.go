package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/g0shi4ek/test_project/config"
	"github.com/jackc/pgx/v5/pgxpool"
)

// мб singleton сделать

func NewUserPool(ctx context.Context, cfg *config.Config) (*pgxpool.Pool, error) {
	connStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.AuthDbHost, cfg.AuthDbPort, cfg.AuthDbUser, cfg.AuthDbPassword, cfg.AuthDbName, cfg.AuthDbSSL,
	)

	config, err := pgxpool.ParseConfig(connStr)
	if err != nil {
		log.Fatalf("Cannot parse config, %v", err)
	}
	config.MaxConns = 10
	config.MaxConnLifetime = time.Hour * 2 // вынести в конфиг

	log.Printf("Connected to database, %s", cfg.AuthDbName)

	return pgxpool.NewWithConfig(ctx, config)
}
