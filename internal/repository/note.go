package repository

import (
	"context"
	"kode-notes/internal/entity"

	"github.com/jackc/pgx/v5/pgxpool"
)

type NoteRepository struct {
	*pgxpool.Pool
}

func NewNoteRepository(pg *pgxpool.Pool) *NoteRepository {
	return &NoteRepository{
		pg,
	}
}

func (r *NoteRepository) CreateNote(ctx context.Context, userId int, text string, mistakes []byte) error {
	_, err := r.Pool.Exec(ctx, "INSERT INTO notes (user_id, text, mistakes) VALUES ($1, $2, $3)", userId, text, mistakes)
	if err != nil {
		return ErrCannotCreateNote
	}
	return nil
}

func (r *NoteRepository) GetNotes(ctx context.Context, userId int) ([]entity.Note, error) {
	notes := make([]entity.Note, 0)

	rows, err := r.Pool.Query(ctx, "SELECT id, user_id, text, mistakes FROM notes WHERE user_id = $1", userId)
	if err != nil {
		return nil, ErrCannotGetNotex
	}
	defer rows.Close()

	for rows.Next() {
		var note entity.Note
		if err := rows.Scan(&note.Id, &note.UserId, &note.Text, &note.Mistakes); err != nil {
			return nil, ErrCannotGetNotex
		}
		notes = append(notes, note)
	}

	return notes, nil
}
