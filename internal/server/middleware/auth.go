package middleware

import (
	"coffeeshop-api/pkg/claims"
	"coffeeshop-api/pkg/errors"
	"coffeeshop-api/pkg/tools"

	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

// Auth middleware validates and parses access token
// and puts the claims to the context.
var Auth echo.MiddlewareFunc

func initAuth(secret string) {
	Auth = echojwt.WithConfig(echojwt.Config{
		TokenLookup:    "cookie:access",
		ContextKey:     "claims",
		ParseTokenFunc: parseToken([]byte(secret)),
		ErrorHandler:   errorHandler,
	})
}

func parseToken(secret []byte) func(c echo.Context, token string) (any, error) {
	return func(c echo.Context, token string) (any, error) {
		var claims claims.Claims

		if _, err := jwt.ParseWithClaims(token, &claims, func(*jwt.Token) (any, error) {
			return secret, nil
		}); err != nil {
			if errors.Is(err, jwt.ErrTokenExpired) {
				return nil, errors.Wrapf(errors.ExpiredToken, "expired token, %v", err)
			}
			return nil, errors.Wrapf(errors.InvalidToken, "invalid token, %v", err)
		}

		return &claims, nil
	}
}

func errorHandler(c echo.Context, jwtErr error) (err error) {
	switch {
	case errors.Is(jwtErr, echojwt.ErrJWTMissing):
		err = errors.Wrapf(errors.MissingToken, "echojwt middleware: %v", jwtErr)

	case errors.Is(jwtErr, echojwt.ErrJWTInvalid):
		err = errors.Wrapf(errors.InvalidToken, "echojwt middleware: %v", jwtErr)
	}

	tools.SendResponse(c, nil, err)

	return err
}
