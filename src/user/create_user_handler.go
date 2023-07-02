package user

import (
	"github.com/google/uuid"
)

type CreateUserCommand struct {
	Name     string
	Password string
}

type CreateUserHandler struct {
	userRepository UserRepository
}

func (h *CreateUserHandler) Handle(command CreateUserCommand) uuid.UUID {
	ValidateUserNameAndPassword(command.Name, command.Password)
	newId := uuid.New()
	h.userRepository.Create(User{newId, command.Name, command.Password})
	return newId
}
