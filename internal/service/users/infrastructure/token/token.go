package token

import (
	"time"

	"coffeeshop-api/internal/service/users/model"
	"coffeeshop-api/pkg/errors"

	"github.com/golang-jwt/jwt/v5"
)

type service struct {
	secret     []byte
	accessExp  time.Duration
	refreshExp time.Duration
}

func New(secret string, accessExp, refreshExp time.Duration) *service {
	return &service{
		secret:     []byte(secret),
		accessExp:  accessExp * time.Minute,
		refreshExp: refreshExp * time.Minute,
	}
}

// Parse token.
func (s *service) Parse(token string) (*model.JWTClaims, error) {
	var claims model.JWTClaims

	if _, err := jwt.ParseWithClaims(token, &claims, func(*jwt.Token) (any, error) {
		return s.secret, nil
	}); err != nil {
		if errors.Is(err, jwt.ErrTokenExpired) {
			return nil, errors.Wrapf(errors.ExpiredToken, "expired token, %v", err)
		}

		return nil, errors.Wrapf(errors.InvalidToken, "invalid token, %v", err)
	}

	return &claims, nil
}

// CreatePair creates new tokens pair with Access and Refresh tokens.
func (s *service) CreatePair(claims *model.JWTClaims) (*model.Token, *model.Token, error) {
	var (
		accessToken  model.Token
		refreshToken model.Token
		err          error
	)

	// Access
	claims.ExpiresAt.Time = time.Now().Add(s.accessExp)

	accessToken.Exp = claims.ExpiresAt.Time
	accessToken.String, err = jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(s.secret)
	if err != nil {
		return nil, nil, errors.Wrap(err, "signing access token")
	}

	// Refresh
	claims.ExpiresAt.Time = time.Now().Add(s.refreshExp)

	refreshToken.Exp = claims.ExpiresAt.Time
	accessToken.String, err = jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(s.secret)
	if err != nil {
		return nil, nil, errors.Wrap(err, "signing refresh token")
	}

	return &accessToken, &refreshToken, nil
}
