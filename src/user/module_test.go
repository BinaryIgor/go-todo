package user

import (
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

	invalidNameTests := shared.MapSlice[string, TestCase](
		[]string{"", "Z", "A!", "A4", "99", "Ala!", test.RandomString(MAX_NAME_LENGTH + 1)},
		func(name string) TestCase {
			return TestCase{CreateUserCommand{name, ""},
				test.ApiErrorFromAppError(NewInvalidUserNameError(name))}
		},
	)

	invalidPasswordTests := shared.MapSlice[string, TestCase](
		[]string{"", "A", "34", "aBZ871a", "LongPasswordWithoutDigit", "lowercasedigits99",
			"UPPERCASEDIGITS1", test.RandomString(MAX_PASSWORD_LENGTH) + "aB6"},
		func(password string) TestCase {
			return TestCase{CreateUserCommand{"Valid Name", password},
				test.ApiErrorFromAppError(InvalidUserPasswordError)}
		},
	)

	tests := append(invalidNameTests, invalidPasswordTests...)

	for _, tc := range tests {
		r := httpClient.PostJson("users/sign-up", tc.command)
		r.ExpectStatusCode(400)
		r.ExpectJson(tc.response)
	}
}
