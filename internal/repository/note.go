package repository

import (
	"context"
	"fmt"
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

func (r *NoteRepository) CreateNote(ctx context.Context, userId int, text string) error {
	_, err := r.Pool.Exec(ctx, "INSERT INTO notes (user_id, text) VALUES ($1, $2)", userId, text)
	if err != nil {
		return fmt.Errorf("repository - CreateNote - r.Pool.Exec: %w", err)
	}
	return nil
}

func (r *NoteRepository) GetNotes(ctx context.Context, userId int) ([]entity.Note, error) {
	notes := make([]entity.Note, 0)

	rows, err := r.Pool.Query(ctx, "SELECT id, user_id, text FROM notes WHERE user_id = $1", userId)
	if err != nil {
		return nil, fmt.Errorf("repository - GetNotes - r.Pool.Query: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var note entity.Note
		if err := rows.Scan(&note.Id, &note.UserId, &note.Text); err != nil {
			return nil, fmt.Errorf("repository - GetNotes - rows.Scan: %w", err)
		}
		notes = append(notes, note)
	}

	return notes, nil
}
