package repositories

import (
	"errors"
	"hexagonal-gofr/internal/core/domain"
	"sync"
)

type UserRepositoryMemory struct {
	users map[int]*domain.User
	mu    sync.Mutex
}

func NewUserRepositoryMemory() *UserRepositoryMemory {
	return &UserRepositoryMemory{
		users: make(map[int]*domain.User),
	}
}

func (r *UserRepositoryMemory) Save(user *domain.User) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.users[user.ID]; exists {
		return errors.New("user already exists")
	}

	r.users[user.ID] = user
	return nil
}

func (r *UserRepositoryMemory) FindByID(id int) (*domain.User, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	user, exists := r.users[id]
	if !exists {
		return nil, errors.New("user not found")
	}

	return user, nil
}
