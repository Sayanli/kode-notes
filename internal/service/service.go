package service

import (
	"context"
	"kode-notes/internal/entity"
	"kode-notes/internal/repository"
	"time"
)

type Auth interface {
	Login(ctx context.Context, username, password string) (string, error)
	Register(ctx context.Context, username, password string) error
	ParseToken(token string) (int, error)
}

type Note interface {
	GetNotes(ctx context.Context, userId int) ([]entity.Note, error)
	CreateNote(ctx context.Context, userId int, text string) error
}

type Service struct {
	Auth
	Note
}

type ServicesDependencies struct {
	Repos    *repository.Repositories
	SignKey  string
	TokenTTL time.Duration
}

func NewService(deps ServicesDependencies) *Service {
	return &Service{
		Auth: NewAuthService(deps.Repos.User),
		Note: NewNoteService(deps.Repos.Note),
	}
}
