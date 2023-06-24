package user

import (
	"go-todo/shared"
	"net/http"

	"github.com/google/uuid"
)

type UserModule struct {
	Handlers map[string]http.HandlerFunc
}

type SignInResponse struct {
	Id uuid.UUID `json:"id"`
}

func Module(tokensSecret []byte) UserModule {
	createUserHandler := CreateUserHandler{}
	//TODO: use it

	authTokensComponent := AuthTokensComponent{tokensSecret}

	signUpHandler := NewSignUpHandler(authTokensComponent)

	handlers := make(map[string]http.HandlerFunc)

	handlers["/users/sign-up"] = func(w http.ResponseWriter, r *http.Request) {
		var command CreateUserCommand
		shared.MustReadJsonBody(r, &command)
		id := createUserHandler.handle(command)
		shared.WriteJsonResponse(w, 201, SignInResponse{id})
	}

	handlers["/users/sign-in"] = func(w http.ResponseWriter, r *http.Request) {
		var command SignUpCommand
		shared.MustReadJsonBody(r, &command)
		response := signUpHandler.handle(command)
		shared.WriteJsonOkResponse(w, response)
	}

	return UserModule{handlers}
}
