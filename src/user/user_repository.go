package user

import (
	"go-todo/shared"
	"sync"

	"github.com/google/uuid"
)

type UserRepository interface {
	Create(user User) error
	FindByName(name string) (*User, error)
}

type inMemoryUserRepository struct {
	db map[uuid.UUID]User
	mu sync.Mutex
}

func New() inMemoryUserRepository {
	return inMemoryUserRepository{db: make(map[uuid.UUID]User)}
}

func (r *inMemoryUserRepository) Create(user User) error {
	return shared.ErrNotFound
}

func (r *inMemoryUserRepository) FindByName(name string) (*User, error) {
	return nil, shared.ErrNotFound
}
