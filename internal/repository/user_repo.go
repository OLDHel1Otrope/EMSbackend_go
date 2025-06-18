package repository

import (
	"internal/model"

	"github.com/jmoiron/sqlx"
)

type UserRepository interface {
	GetUserByID(id string) (*model.User, error)
	CreateUser(user *model.User) (*model.User, error)
	ArchiveUser(id string) error
}

type userRepo struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) UserRepository {
	return &userRepo{db: db}
}

func (r *userRepo) GetUserByID(id string) (*model.User, error) {
	var user model.User
	err := r.db.Get(&user, "SELECT * FROM users WHERE id = $1", id)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepo) CreateUser(user *model.User) (*model.User, error) {
	query := `
		INSERT INTO users (name, email, password)
		VALUES ($1, $2, $3)
		RETURNING id, created_at
	`
	err := r.db.QueryRowx(query, user.Name, user.Email, user.Password).Scan(&user.ID, &user.CreatedAt)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *userRepo) ArchiveUser(id string) error {
	query := `UPDATE users SET archived_at = NOW() WHERE id = $1`
	_, err := r.db.Exec(query, id)
	return err
}

func (r *userRepo) GetPasswordByEmail(email string) (string, error) {
	SQL := `SELECT password from users where email = $1 AND archived_at IS NULL`
	var pwd string
	err := r.db.Get(&pwd, SQL, email)
	return pwd, err
}
