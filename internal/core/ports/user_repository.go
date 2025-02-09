package ports

import "hexagonal-gofr/internal/core/domain"

type UserRepository interface {
	Save(user *domain.User) error
	FindByID(id int) (*domain.User, error)
}
