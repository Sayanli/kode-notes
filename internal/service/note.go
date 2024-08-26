package service

import (
	"context"
	"kode-notes/internal/entity"
	"kode-notes/internal/repository"
	"kode-notes/internal/spellchecker"
	"log/slog"
	"strconv"
)

type NoteService struct {
	repo    repository.Note
	speller spellchecker.SpellChecker
	logger  *slog.Logger
}

func NewNoteService(noteRepo repository.Note, speller spellchecker.SpellChecker, logger *slog.Logger) *NoteService {
	return &NoteService{
		repo:    noteRepo,
		speller: speller,
		logger:  logger,
	}
}

func (s *NoteService) CreateNote(ctx context.Context, userId int, text string) error {
	const op = "service.Note.CreateNote"
	s.logger = s.logger.With("op", op)

	if text == "" {
		return ErrTextRequired
	}
	mistakes, err := s.speller.Check(text)
	if err != nil {
		s.logger.Error("cannot check mistakes", slog.String("userId", strconv.Itoa(userId)))
		return ErrCannotCheckMistakes
	}
	return s.repo.CreateNote(ctx, userId, text, mistakes)
}

func (s *NoteService) GetNotes(ctx context.Context, userId int) ([]entity.Note, error) {
	return s.repo.GetNotes(ctx, userId)
}
