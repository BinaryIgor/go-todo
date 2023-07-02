package auth

import (
	"fmt"
	"time"

	"go-todo/shared"
	"go-todo/user"

	"github.com/golang-jwt/jwt/v5"
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
	secret          []byte
	accessDuration  time.Duration
	refreshDuration time.Duration
	issuer          string
	clock           shared.Clock
}

func NewAuthTokensComponent(secret []byte,
	accessDuration time.Duration,
	refreshDuration time.Duration,
	issuer string,
	clock shared.Clock) AuthTokensComponent {
	return &jwtAuthTokensComponent{secret: secret,
		accessDuration:  accessDuration,
		refreshDuration: refreshDuration,
		issuer:          issuer,
		clock:           clock}
}

func (c *jwtAuthTokensComponent) CreateTokens(userId uuid.UUID) AuthTokens {
	return AuthTokens{
		access:  c.createToken(userId, c.accessDuration),
		refresh: c.createToken(userId, c.refreshDuration),
	}
}

func (c *jwtAuthTokensComponent) createToken(userId uuid.UUID, expiresAfter time.Duration) AuthToken {
	now := c.clock.Now()
	expiresAt := now.Add(expiresAfter)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"issuer": c.issuer,
		"sub":    userId,
		"iss":    now,
		"exp":    expiresAt,
	})
	tokenString, err := token.SignedString(c.secret)
	if err != nil {
		panic(err)
	}
	return AuthToken{tokenString, expiresAt}
}

// TODO: improve this!
func (c *jwtAuthTokensComponent) RefreshTokens(refeshToken string) AuthTokens {
	info := c.parseToken(refeshToken)
	return c.CreateTokens(info.UserId)
}

// TODO: type??
func (c *jwtAuthTokensComponent) parseToken(tokenString string) TokenInfo {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return c.secret, nil
	})

	if err != nil {
		user.InvalidAuthTokenError.Throw()
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return c.validateClaims(claims)
	}

	panic(user.InvalidAuthTokenError)
}

// TODO check claims!
func (c *jwtAuthTokensComponent) validateClaims(claims jwt.MapClaims) TokenInfo {
	sub, _ := claims.GetSubject()

	return TokenInfo{UserId: uuid.MustParse(sub)}
}

func (c *jwtAuthTokensComponent) ValidateAccessToken(accessToken string) TokenInfo {
	return c.parseToken(accessToken)
}
