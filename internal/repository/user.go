package repository

import (
	"context"
	"fmt"
	"kode-notes/internal/entity"

	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepository struct {
	*pgxpool.Pool
}

func NewUserRepository(pg *pgxpool.Pool) *UserRepository {
	return &UserRepository{
		pg,
	}
}

func (r *UserRepository) CreateUser(ctx context.Context, username, password string) error {
	_, err := r.Pool.Exec(ctx, "INSERT INTO users (username, password) VALUES ($1, $2)", username, password)
	if err != nil {
		return fmt.Errorf("repository - CreateUser - r.Pool.Exec: %w", err)
	}
	return nil
}

func (r *UserRepository) GetUser(ctx context.Context, username, password string) (entity.User, error) {
	var user entity.User
	err := r.Pool.QueryRow(ctx, "SELECT id, username FROM users WHERE username = $1 AND password = $2", username, password).Scan(&user.Id, &user.Username)
	if err != nil {
		return entity.User{}, fmt.Errorf("repository - GetUser - r.Pool.QueryRow: %w", err)
	}
	return user, nil
}
