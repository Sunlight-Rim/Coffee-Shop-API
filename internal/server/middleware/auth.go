package middleware

import (
	"coffeeshop-api/pkg/claims"
	"coffeeshop-api/pkg/errors"
	"coffeeshop-api/pkg/tools"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

var (
	// Auth middleware validates and parses user
	// access token and puts the claims to the context.
	Auth echo.MiddlewareFunc

	// Auth middleware validates employee Bearer authorization token.
	AuthEmployee echo.MiddlewareFunc
)

func initAuth(tokenSecret, employeeToken string) {
	Auth = echojwt.WithConfig(echojwt.Config{
		TokenLookup:    "cookie:access",
		ContextKey:     "claims",
		ParseTokenFunc: parseToken([]byte(tokenSecret)),
		ErrorHandler:   errorHandler,
	})

	AuthEmployee = parseEmployeeToken(employeeToken)
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

func parseEmployeeToken(employeeToken string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			bearerToken := strings.Split(c.Request().Header.Get("Authorization"), " ")

			if len(bearerToken) < 2 {
				err := errors.Wrap(errors.MissingToken, "missing header with bearer authorization")
				tools.SendResponse(c, nil, err)
				return err
			}

			if bearerToken[0] != "Bearer" || bearerToken[1] != employeeToken {
				err := errors.Wrap(errors.InvalidToken, "invalid header with bearer authorization")
				tools.SendResponse(c, nil, err)
				return err
			}

			return next(c)
		}
	}
}
