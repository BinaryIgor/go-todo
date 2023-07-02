package user

import (
	"go-todo/shared"
	"net/http"

	"encoding/hex"

	"time"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

type ModuleConfig struct {
	JwtSecret          string
	JwtAccessDuration  time.Duration
	JwtRefreshDuration time.Duration
	JwtIssuer          string
	Clock              shared.Clock
}

func ModuleConfigFromEnvVariables(clock shared.Clock) ModuleConfig {
	return ModuleConfig{
		JwtSecret:          shared.MustGetenv("JWT_SECRET"),
		JwtAccessDuration:  time.Duration(shared.MustGetenvAsInt("JWT_ACCESS_DURATION")),
		JwtRefreshDuration: time.Duration(shared.MustGetenvAsInt("JWT_REFRESH_DURATION")),
		JwtIssuer:          shared.MustGetenv("JWT_ISSUER"),
		Clock:              clock,
	}
}

type SignUpResponse struct {
	Id uuid.UUID `json:"id"`
}

func Module(config ModuleConfig) shared.AppModule {
	tokensSecretBytes, err := hex.DecodeString(config.JwtSecret)
	if err != nil {
		panic(err)
	}

	userRepository := NewUserRepository()

	createUserHandler := CreateUserHandler{userRepository}

	authTokensComponent := NewAuthTokensComponent(tokensSecretBytes,
		config.JwtAccessDuration, config.JwtRefreshDuration, config.JwtIssuer,
		config.Clock)

	signInHandler := SignInHandler{userRepository, authTokensComponent}

	router := chi.NewRouter()

	router.Post("/sign-up", func(w http.ResponseWriter, r *http.Request) {
		var command CreateUserCommand
		shared.MustReadJsonBody(r, &command)
		userId := createUserHandler.Handle(command)
		shared.WriteJsonResponse(w, 201, SignUpResponse{userId})
	})

	router.Post("/sign-in", func(w http.ResponseWriter, r *http.Request) {
		var command SignInCommand
		shared.MustReadJsonBody(r, &command)
		response := signInHandler.Handle(command)
		shared.WriteJsonOkResponse(w, response)
	})

	return shared.AppModule{Router: router, Path: "/users"}
}
