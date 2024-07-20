package token

import (
	"time"

	"coffeeshop-api/internal/services/users/model"
	"coffeeshop-api/pkg/claims"
	"coffeeshop-api/pkg/errors"

	"github.com/golang-jwt/jwt/v5"
)

type provider struct {
	secret     []byte
	accessExp  time.Duration
	refreshExp time.Duration
}

func New(secret string, accessExp, refreshExp time.Duration) *provider {
	return &provider{
		secret:     []byte(secret),
		accessExp:  accessExp * time.Minute,
		refreshExp: refreshExp * time.Minute,
	}
}

// Parse token.
func (p *provider) Parse(token string) (*claims.Claims, error) {
	var claims claims.Claims

	if _, err := jwt.ParseWithClaims(token, &claims, func(*jwt.Token) (any, error) {
		return p.secret, nil
	}); err != nil {
		if errors.Is(err, jwt.ErrTokenExpired) {
			return nil, errors.Wrapf(errors.ExpiredToken, "expired token, %v", err)
		}

		return nil, errors.Wrapf(errors.InvalidToken, "invalid token, %v", err)
	}

	return &claims, nil
}

// CreatePair creates new tokens pair.
func (p *provider) CreatePair(claims *claims.Claims) (*model.TokensPair, error) {
	var (
		accessToken  model.Token
		refreshToken model.Token
		err          error
	)

	// Access
	claims.ExpiresAt = &jwt.NumericDate{
		Time: time.Now().Add(p.accessExp),
	}
	accessToken.Exp = claims.ExpiresAt.Time
	accessToken.String, err = jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(p.secret)
	if err != nil {
		return nil, errors.Wrap(err, "signing access token")
	}

	// Refresh
	claims.ExpiresAt = &jwt.NumericDate{
		Time: time.Now().Add(p.refreshExp),
	}
	refreshToken.Exp = claims.ExpiresAt.Time
	refreshToken.String, err = jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(p.secret)
	if err != nil {
		return nil, errors.Wrap(err, "signing refresh token")
	}

	return &model.TokensPair{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}
