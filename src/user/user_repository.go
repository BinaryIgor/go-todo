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

func NewUserRepository() UserRepository {
	return &inMemoryUserRepository{db: make(map[uuid.UUID]User)}
}

func (r *inMemoryUserRepository) Create(user User) {
	r.db[user.Id] = user
}

func (r *inMemoryUserRepository) FindByName(name string) *User {
	for _, v := range r.db {
		if v.Name == name {
			return &v
		}
	}
	return nil
}
