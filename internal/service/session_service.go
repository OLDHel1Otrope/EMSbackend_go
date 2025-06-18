package service

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"server.go/internal/model"
	"server.go/internal/repository"
)

type SessionService struct {
	sessionRepo repository.SessionRepository
}

func NewSessionService(sessionRepo repository.SessionRepository) *SessionService {
	return &SessionService{
		sessionRepo: sessionRepo,
	}
}

func (s *SessionService) Login(user *model.User) (*model.Session, error) {
	claims := jwt.MapClaims{
		"user_id": user.ID,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(s.jwtSecret))
	if err != nil {
		return nil, err
	}

	// Store session in DB
	return s.sessionRepo.CreateSession(user, signedToken)
}

func (s *SessionService) Logout(token string) error {
	return s.sessionRepo.DelSession(token)
}

func (s *SessionService) GetSession(token string) (*model.Session, error) {
	return s.sessionRepo.GetSession(token)
}
