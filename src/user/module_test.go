package user

import (
	"fmt"
	"go-todo/shared"
	"go-todo/test"
	"os"
	"testing"
)

var appPort int
var tokensSecret = "9123"

func TestMain(m *testing.M) {
	appPort = test.StartApp(Module(tokensSecret))
	os.Exit(m.Run())
}

func TestShouldReturnBadRequestWithInvalidCreateUserCommands(t *testing.T) {
	httpClient := test.NewHttpClient(t, appPort)

	type test struct {
		command  CreateUserCommand
		response shared.ApiError
	}

	tests := []test{
		{CreateUserCommand{"", ""}, shared.NewApiError("INVALID_NAME", "Name is not valid")},
	}

	for _, tc := range tests {
		fmt.Println("Running test...")
		r := httpClient.PostJson("users/sign-up", tc.command)
		httpClient.ExpectResponseCode(r, 400)
	}
}
