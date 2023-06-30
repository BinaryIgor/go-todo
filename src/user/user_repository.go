package user

import (
	"sync"

	"github.com/google/uuid"
)

type UserRepository interface {
	Create(user User)
	FindByName(name string) *User
}

type inMemoryUserRepository struct {
	db map[uuid.UUID]User
	mu sync.Mutex
}

func NewUserRepository() inMemoryUserRepository {
	return inMemoryUserRepository{db: make(map[uuid.UUID]User)}
}

func (r *inMemoryUserRepository) Create(user User) {

}

func (r *inMemoryUserRepository) FindByName(name string) *User {
	return nil
}
