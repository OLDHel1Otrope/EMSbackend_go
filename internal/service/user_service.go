package service

import (
	"errors"
	"time"

	"server.go/internal/model"
	"server.go/internal/repository"
	"server.go/pkg/utils"
)

type UserService interface {
	CreateUser(req model.CreateUserRequest) (*model.User, error)
	GetUserByID(id string) (*model.User, error)
	ArchiveUser(id string) error
}

type userService struct {
	userRepo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{
		userRepo: repo,
	}
}

func (s *userService) CreateUser(req model.CreateUserRequest) (*model.User, error) {
	if req.Email == "" || req.Password == "" || req.Name == "" {
		return nil, errors.New("missing required fields")
	}

	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		return nil, err
	}

	user := &model.User{
		Name:      req.Name,
		Email:     req.Email,
		Password:  hashedPassword,
		CreatedAt: time.Now(),
	}

	createdUser, err := s.userRepo.CreateUser(user)
	if err != nil {
		return nil, err
	}

	return createdUser, nil
}

func (s *userService) GetUserByID(id string) (*model.User, error) {
	return s.userRepo.GetUserByID(id)
}

func (s *userService) ArchiveUser(id string) error {
	return s.userRepo.ArchiveUser(id)
}
