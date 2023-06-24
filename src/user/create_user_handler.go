package user

import (
	"fmt"

	"github.com/google/uuid"
)

type CreateUserCommand struct {
	Name     string
	Password string
}

type CreateUserHandler struct {
}

func (h *CreateUserHandler) handle(command CreateUserCommand) uuid.UUID {
	fmt.Println("Creating new user!", command)
	return uuid.New()
}
