package user

import (
	"go-todo/shared"
	"net/http"

	"encoding/hex"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

type ApiSignInResponse struct {
	Id uuid.UUID `json:"id"`
}

func Module(tokensSecret string) shared.AppModule {
	tokensSecretBytes, err := hex.DecodeString(tokensSecret)
	if err != nil {
		panic(err)
	}

	userRepository := NewUserRepository()

	createUserHandler := CreateUserHandler{userRepository}

	authTokensComponent := NewAuthTokensComponent(tokensSecretBytes)

	signInHandler := SignInHandler{userRepository, authTokensComponent}

	router := chi.NewRouter()

	router.Post("/sign-up", func(w http.ResponseWriter, r *http.Request) {
		var command CreateUserCommand
		shared.MustReadJsonBody(r, &command)
		userId := createUserHandler.Handle(command)
		shared.WriteJsonResponse(w, 201, ApiSignInResponse{userId})
	})

	router.Post("/sign-in", func(w http.ResponseWriter, r *http.Request) {
		var command SignInCommand
		shared.MustReadJsonBody(r, &command)
		response := signInHandler.Handle(command)
		shared.WriteJsonOkResponse(w, response)
	})

	return shared.AppModule{Router: router, Path: "/users"}
}
