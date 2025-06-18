package repository

import (
	"server.go/internal/model"

	"github.com/jmoiron/sqlx"
)

type SessionRepository interface {
	GetSession(sessionToken string) (*model.Session, error)
	CreateSession(user *model.User) (*model.Session, error)
	DelSession(sessionToken string) error
}

func NewSessionRepository(db *sqlx.DB) SessionRepository {
	return &sessionRepo{db: db}
}

type sessionRepo struct {
	db *sqlx.DB
}

func (r *sessionRepo) GetSession(sessionToken string) (*model.Session, error) {
	const query = `
		SELECT user_id,
			CASE 
				WHEN created_at > NOW() - INTERVAL '100 Days' THEN TRUE 
				ELSE FALSE 
			END AS is_session_valid
		FROM sessions 
		WHERE id=$1 AND archived_at IS NULL
	`

	var sess model.Session
	err := r.db.Get(&sess, query, sessionToken)
	if err != nil {
		return nil, err
	}
	return &sess, nil
}

func (r *sessionRepo) CreateSession(user *model.User) (*model.Session, error) {
	const query = `
		INSERT INTO sessions (id, user_id, token, created_at) 
		VALUES (gen_random_uuid(), $1, $2, NOW()) 
		RETURNING user_id, TRUE AS is_session_valid
	`
	var sess model.Session
	err := r.db.Get(&sess, query, user.ID, user.Token)
	if err != nil {
		return nil, err
	}
	return &sess, nil
}

func (r *sessionRepo) DelSession(sessionToken string) error {
	const query = `UPDATE sessions SET archived_at = NOW() WHERE id = $1`
	_, err := r.db.Exec(query, sessionToken)
	return err
}
