package services

import (
	"hexagonal-gofr/internal/core/domain"
	"hexagonal-gofr/internal/core/ports"
)

type UserServiceImpl struct {
	repo ports.UserRepository
}

func NewUserService(repo ports.UserRepository) *UserServiceImpl {
	return &UserServiceImpl{repo: repo}
}

func (s *UserServiceImpl) CreateUser(user *domain.User) error {
	return s.repo.Save(user)
}

func (s *UserServiceImpl) GetUserByID(id int) (*domain.User, error) {
	return s.repo.FindByID(id)
}
