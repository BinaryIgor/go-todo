package user

import (
	"fmt"

	"go-todo/shared"

	"github.com/google/uuid"
)

type CreateUserCommand struct {
	Name     string
	Password string
}

type CreateUserHandler struct {
	userRepository UserRepository
}

func (h *CreateUserHandler) handle(command CreateUserCommand) shared.Result[uuid.UUID] {
	fmt.Println("Creating new user!", command)
	result := validateCommand(command)
	if result.IsFailure() {
		return shared.FailureResult[uuid.UUID](result.Error())
	}

	newId := uuid.New()
	err := h.userRepository.Create(User{newId, command.Name, command.Password})
	if err != nil {
		return shared.FailureResult[uuid.UUID](shared.NewAppErrorWithMessage(err, "Failure to create user"))
	}

	return shared.SuccessResult(newId)
}

func validateCommand(command CreateUserCommand) shared.EmptyResult {
	//TODO: validate
	return shared.SuccessEmptyResult()
}
