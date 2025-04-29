package repository

import (
	"context"

	"github.com/g0shi4ek/test_project/internal/models"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user *models.User) error
	GetUserByEmail(ctx context.Context, email string) (*models.User, error)
}

type UserRepo struct {
	pg *pgxpool.Pool
}

func NewUserRepository(pg *pgxpool.Pool) UserRepository {
	return &UserRepo{pg: pg}
}

func (r *UserRepo) CreateUser(ctx context.Context, user *models.User) error {
	query := "INSERT INTO Users (email, password_hash) VALUES ($1, $2) RETURNING user_id"
	return r.pg.QueryRow(ctx, query, user.Email, user.PasswordHash).Scan(&user.UserUUID)
}

func (r *UserRepo) GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	var user models.User
	query := "SELECT user_id, email, password_hash FROM Users WHERE email = $1"
	err := r.pg.QueryRow(ctx, query, email).Scan(&user.UserUUID, &user.Email, &user.PasswordHash)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
