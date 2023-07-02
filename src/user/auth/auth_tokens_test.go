package auth

import (
	"go-todo/test"
	"testing"
)

var testClock = test.TestClock{}

func NewAuthTokensComponentInstance() AuthTokensComponent {
	return NewAuthTokensComponent()
}

func TestShouldGenereateSth(t *testing.T) {
	component = NewAuthTokensComponentInstance()
}
