package service

import (
	"context"
	"fmt"
	"kode-notes/internal/entity"
	"kode-notes/internal/repository"
	"kode-notes/internal/spellchecker"
)

type NoteService struct {
	repo    repository.Note
	speller spellchecker.SpellChecker
}

func NewNoteService(noteRepo repository.Note, speller spellchecker.SpellChecker) *NoteService {
	return &NoteService{
		repo:    noteRepo,
		speller: speller,
	}
}

func (s *NoteService) CreateNote(ctx context.Context, userId int, text string) error {
	mistakes, err := s.speller.Check(text)
	if err != nil {
		return fmt.Errorf("service - CreateNote - s.spellchecker.Check: %w", err)
	}
	return s.repo.CreateNote(ctx, userId, text, mistakes)
}

func (s *NoteService) GetNotes(ctx context.Context, userId int) ([]entity.Note, error) {
	return s.repo.GetNotes(ctx, userId)
}
