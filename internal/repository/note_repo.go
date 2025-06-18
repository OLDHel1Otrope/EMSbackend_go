package repository

import (
	"github.com/jmoiron/sqlx"
	"server.go/internal/model"
)

type NoteRepository interface {
	GetNotesByUserID(user_id int) (*model.Note, error)
	GetNoteByNoteID(note_id int) (*model.Note, error)
}

type noteRepo struct {
	db *sqlx.DB
}

func NewNoteRepository(db *sqlx.DB) UserRepository {
	return &userRepo{db: db}
}

func (r *noteRepo) GetNotesByUserID(user_id int) (*model.Note, error) {
	var note model.Note
	err := r.db.Get(&note, "SELECT * FROM notes WHERE user_id = $1", user_id)
	return &note, err
}

func (r *noteRepo) GetNoteByNoteID(note_id int) (*model.Note, error) {
	var note model.Note
	err := r.db.Get(&note, "SELECT * FROM notes WHERE id = $1", user_id)
	return &note, err
}
