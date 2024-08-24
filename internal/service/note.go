package service

import (
	"context"
	"kode-notes/internal/entity"
	"kode-notes/internal/repository"
)

type NoteService struct {
	repo repository.Note
}

func NewNoteService(noteRepo repository.Note) *NoteService {
	return &NoteService{
		repo: noteRepo,
	}
}

func (s *NoteService) CreateNote(ctx context.Context, userId int, text string) error {
	return s.repo.CreateNote(ctx, userId, text)
}

func (s *NoteService) GetNotes(ctx context.Context, userId int) ([]entity.Note, error) {
	return s.repo.GetNotes(ctx, userId)
}
