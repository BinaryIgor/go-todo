package user

import (
	"go-todo/test"
	"testing"
)

var appPort int
var tokensSecret = "9123"

func TestMain(m *testing.M) {
	appPort = test.StartApp(Module(tokensSecret))
}
