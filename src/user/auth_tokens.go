package user

import (
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
)

type AuthToken struct {
	value     string
	expiresAt time.Time
}

type AuthTokens struct {
	access  AuthToken
	refresh AuthToken
}

type TokenInfo struct {
	UserId uuid.UUID
}

type AuthTokensComponent interface {
	CreateTokens(userId uuid.UUID) AuthTokens
	RefreshTokens(refreshToken string) AuthTokens
	ValidateAccessToken(accessToken string) TokenInfo
}

type jwtAuthTokensComponent struct {
	secret []byte
}

func NewAuthTokensComponent(secret []byte) AuthTokensComponent {
	return &jwtAuthTokensComponent{secret: secret}
}

func (c *jwtAuthTokensComponent) CreateTokens(userId uuid.UUID) AuthTokens {
	return AuthTokens{
		access:  AuthToken{c.dummyToken(userId), time.Now().Add(1 * time.Hour)},
		refresh: AuthToken{c.dummyToken(userId), time.Now().Add(10 * time.Hour)},
	}
}

func (c *jwtAuthTokensComponent) dummyToken(userId uuid.UUID) string {
	return strings.Join([]string{userId.String(), string(c.secret)}, "_")
}

// TODO: improve this!
func (c *jwtAuthTokensComponent) RefreshTokens(refeshToken string) AuthTokens {
	userId := c.extractUserIdFromToken(refeshToken)
	return c.CreateTokens(userId)
}

func (c *jwtAuthTokensComponent) extractUserIdFromToken(token string) uuid.UUID {
	parts := strings.Split(token, "_")
	if len(parts) != 2 {
		ThrowInvalidAuthTokenError("Invalid token format")
	}
	userId, err := uuid.Parse(parts[0])
	if err != nil {
		ThrowInvalidAuthTokenError(fmt.Sprintln("Invalid user id, uuid required, but was:", parts[0]))
	}
	return userId
}

func (c *jwtAuthTokensComponent) ValidateAccessToken(accessToken string) TokenInfo {
	userId := c.extractUserIdFromToken(accessToken)
	return TokenInfo{UserId: userId}
}
