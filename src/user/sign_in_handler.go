package user

import (
	"github.com/google/uuid"
)

type SignInCommand struct {
	Name     string
	Password string
}

type SignInHandler struct {
	userRespository UserRepository
	tokensComponent AuthTokensComponent
}

type SignInResponse struct {
	Id     uuid.UUID
	Name   string
	Tokens AuthTokens
}

func (h *SignInHandler) Handle(command SignInCommand) SignInResponse {
	ValidateUserNameAndPassword(command.Name, command.Password)

	user := h.userRespository.FindByName(command.Name)
	if user == nil {
		NewUserNotFoundError(command.Name).Throw()
	}

	return SignInResponse{user.Id, user.Name, h.tokensComponent.CreateTokens(user.Id)}
}
