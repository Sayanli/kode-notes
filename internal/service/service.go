package service

import (
	"context"
	"kode-notes/internal/entity"
	"kode-notes/internal/repository"
	"kode-notes/internal/spellchecker"
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
	Salt     string
	Speller  spellchecker.SpellChecker
}

func NewService(deps ServicesDependencies) *Service {
	return &Service{
		Auth: NewAuthService(AuthDependencies{
			userRepo: deps.Repos.User,
			signKey:  deps.SignKey,
			tokenTTL: deps.TokenTTL,
			salt:     deps.Salt,
		}),
		Note: NewNoteService(deps.Repos.Note, deps.Speller),
	}
}
