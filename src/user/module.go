package user

import (
	"go-todo/shared"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

type UserModule struct {
	Router http.Handler
}

type SignInResponse struct {
	Id uuid.UUID `json:"id"`
}

func Module(tokensSecret []byte) UserModule {
	createUserHandler := CreateUserHandler{}
	//TODO: use it

	authTokensComponent := AuthTokensComponent{tokensSecret}

	signUpHandler := NewSignUpHandler(authTokensComponent)

	router := chi.NewRouter()

	router.Post("/sign-up", func(w http.ResponseWriter, r *http.Request) {
		var command CreateUserCommand
		shared.MustReadJsonBody(r, &command)
		result := createUserHandler.handle(command)
		if result.IsSuccess() {
			shared.WriteJsonResponse(w, 201, SignInResponse{result.Value()})
		} else {
			shared.WriteJsonErrorResponse(w, 400, result.Error())
		}
	})

	router.Post("/sign-in", func(w http.ResponseWriter, r *http.Request) {
		var command SignUpCommand
		shared.MustReadJsonBody(r, &command)
		response := signUpHandler.handle(command)
		shared.WriteJsonOkResponse(w, response)
	})

	return UserModule{router}
}
