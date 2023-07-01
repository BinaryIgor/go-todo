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
	httpClient := test.NewHttpClient[shared.ApiError](t, appPort)

	type TestCase struct {
		command  CreateUserCommand
		response shared.ApiError
	}

	tests := []TestCase{
		{CreateUserCommand{"", ""}, test.ApiErrorFromAppError(NewInvalidUserNameError(""))},
		{CreateUserCommand{"A", ""}, test.ApiErrorFromAppError(NewInvalidUserNameError("A"))},
		{CreateUserCommand{"A4", ""}, test.ApiErrorFromAppError(NewInvalidUserNameError("A4"))},
		{CreateUserCommand{"99", ""}, test.ApiErrorFromAppError(NewInvalidUserNameError("99"))},
		{CreateUserCommand{"Alaa!", ""}, test.ApiErrorFromAppError(NewInvalidUserNameError("Alaa!"))},
	}

	for _, tc := range tests {
		fmt.Println("Running test...")
		r := httpClient.PostJson("users/sign-up", tc.command)
		r.ExpectStatusCode(400)
		r.ExpectJson(tc.response)
	}
}
