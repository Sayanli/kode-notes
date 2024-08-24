package repository

import (
	"context"
	"kode-notes/internal/entity"
	"kode-notes/pkg/postgres"
)

type User interface {
	CreateUser(ctx context.Context, username, password string) error
	GetUser(ctx context.Context, username, password string) (entity.User, error)
}

type Note interface {
	GetNotes(ctx context.Context, userId int) ([]entity.Note, error)
	CreateNote(ctx context.Context, userId int, text string) error
}

type Repositories struct {
	User
	Note
}

func NewRepositories(pg *postgres.Postgres) *Repositories {
	return &Repositories{
		User: NewUserRepository(pg),
		Note: NewNoteRepository(pg),
	}
}
