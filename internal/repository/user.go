package repository

import (
	"context"
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
	exists, err := r.checkUserExists(ctx, username)
	if err != nil {
		return ErrCannotCheckUserExist
	} else if exists {
		return ErrUserAlreadyExists
	}

	_, err = r.Pool.Exec(ctx, "INSERT INTO users (username, password) VALUES ($1, $2)", username, password)
	if err != nil {
		return ErrCannotCreateUser
	}
	return nil
}

func (r *UserRepository) checkUserExists(ctx context.Context, username string) (bool, error) {
	var exists bool
	err := r.Pool.QueryRow(ctx, "SELECT EXISTS(SELECT 1 FROM users WHERE username = $1)", username).Scan(&exists)
	if err != nil {
		return false, ErrCannotCheckUserExist
	}
	return exists, nil
}

func (r *UserRepository) GetUser(ctx context.Context, username, password string) (entity.User, error) {
	var user entity.User
	err := r.Pool.QueryRow(ctx, "SELECT id, username FROM users WHERE username = $1 AND password = $2", username, password).Scan(&user.Id, &user.Username)
	if err != nil {
		return entity.User{}, ErrCannotGetUser
	}
	return user, nil
}
