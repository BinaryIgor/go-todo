package user

import (
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

type AuthTokensComponent struct {
	secret []byte
}

func (c *AuthTokensComponent) generateTokens(userId uuid.UUID) AuthTokens {
	return AuthTokens{
		access:  AuthToken{userId.String() + string(c.secret), time.Now().Add(1 * time.Hour)},
		refresh: AuthToken{userId.String() + string(c.secret), time.Now().Add(10 * time.Hour)},
	}
}
