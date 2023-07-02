package auth

import (
	"go-todo/test"
	"os"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

var accessDuration = 5 * time.Minute
var refreshDuration = 15 * time.Minute
var issuer = "some-app"
var secret = test.RandomBytes(10)

var testClock = test.NewTestClock()
var component AuthTokensComponent

func TestMain(m *testing.M) {
	component = NewAuthTokensComponent(secret, accessDuration, refreshDuration, issuer, testClock)
	os.Exit(m.Run())
}

func TestShouldCreateAccessAndRefreshTokensForUser(t *testing.T) {
	userId := uuid.New()

	tokens := component.CreateTokens(userId)

	assert.Equal(t, tokens.access.expiresAt, testClock.Now().Add(accessDuration))
	assert.Equal(t, tokens.refresh.expiresAt, testClock.Now().Add(refreshDuration))

	tokenInfo := component.ValidateAccessToken(tokens.access.value)

	assert.Equal(t, TokenInfo{userId}, tokenInfo)

	testClock.AddTime(1 * time.Minute)

	refreshedTokens := component.RefreshTokens(tokens.refresh.value)

	assert.NotEqual(t, tokenInfo, refreshedTokens)

	assert.Equal(t, TokenInfo{userId}, tokenInfo)
}

func TestShouldNotValidateExpiredTokens(t *testing.T) {
	userId := uuid.New()

	tokens := component.CreateTokens(userId)

	testClock.AddTime(accessDuration)
	testClock.AddTime(1 * time.Second)

	assert.PanicsWithValue(t, ExpiredAuthTokenError, func() {
		component.ValidateAccessToken(tokens.access.value)
	})
}
