package ports

import "hexagonal-gofr/internal/core/domain"

type UserService interface {
	CreateUser(user *domain.User) error
	GetUserByID(id int) (*domain.User, error)
}
