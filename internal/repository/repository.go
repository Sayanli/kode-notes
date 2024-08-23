package repository

import (
	"context"
	"kode-notes/internal/entity"
)

type User interface {
	CreateUser(ctx context.Context, user entity.User) error
	GetNotes(ctx context.Context, userId int) ([]entity.Note, error)
}

type Repository struct {
	User
}

func NewRepository() *Repository {
	return &Repository{
		User: NewUserRepository(),
	}
}
