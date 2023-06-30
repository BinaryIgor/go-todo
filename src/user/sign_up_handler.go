package user

import (
	"fmt"

	"github.com/google/uuid"
)

type SignUpCommand struct {
	Name     string
	Password string
}

type SignUpHandler struct {
	tokensComponent AuthTokensComponent
}

type SignUpResponse struct {
	Id     uuid.UUID
	Name   string
	Tokens AuthTokens
}

func NewSignUpHandler(tokensComponent AuthTokensComponent) SignUpHandler {
	return SignUpHandler{tokensComponent}
}

func (h *SignUpHandler) Handle(command SignUpCommand) SignUpResponse {
	fmt.Println("Sign up user")
	id := uuid.New()
	return SignUpResponse{id, "Some user", h.tokensComponent.CreateTokens(id)}
}
