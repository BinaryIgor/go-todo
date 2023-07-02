package user

import (
	"go-todo/shared"
	"go-todo/test"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
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

func TestShouldAllowToCreateUserAndThenSignUp(t *testing.T) {
	createUserClient := test.NewHttpClient[SignUpResponse](t, appPort)

	createUserCommand := CreateUserCommand{"User", "GoodEnoughPassword12"}

	createUserResponse := createUserClient.PostJson("users/sign-up", createUserCommand)
	createUserResponse.ExpectStatusCode(201)

	userId := createUserResponse.ExpectJsonBody().Id

	signInHttpClient := test.NewHttpClient[SignInResponse](t, appPort)

	signInCommand := SignInCommand{createUserCommand.Name, createUserCommand.Password}

	signUpResponse := signInHttpClient.PostJson("users/sign-in", signInCommand)
	signUpResponse.ExpectOkStatusCode()

	signInResponseBody := signUpResponse.ExpectJsonBody()
	assert.Equal(t, userId, signInResponseBody.Id)
	assert.Equal(t, signInCommand.Name, signInResponseBody.Name)
	//TODO: test tokens!
	// assert.Equal(t, userId, signInResponseBody.Name)

}
