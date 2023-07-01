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
	validateCommand(command)
	newId := uuid.New()
	h.userRepository.Create(User{newId, command.Name, command.Password})
	return newId
}

func validateCommand(command CreateUserCommand) {
	NewInvalidUserNameError(command.Name).Throw()
}
