package server

import (
	"coffeeshop-api/internal/service/users/model"
	"coffeeshop-api/pkg/errors"

	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sirupsen/logrus"
)

// Response logger.
var LoggerMW = middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
	LogURI:       true,
	LogMethod:    true,
	LogStatus:    true,
	LogLatency:   true,
	LogRemoteIP:  true,
	LogRequestID: true,
	LogUserAgent: true,
	LogError:     true,
	LogHeaders:   []string{"Cookie"},

	LogValuesFunc: func(c echo.Context, values middleware.RequestLoggerValues) error {
		errLocation, _ := errors.GetLocation(values.Error)

		logrus := logrus.WithFields(logrus.Fields{
			"uri":            values.URI,
			"method":         values.Method,
			"status":         values.Status,
			"latency":        values.Latency,
			"remote_ip":      values.RemoteIP,
			"request_id":     values.RequestID,
			"user_agent":     values.UserAgent,
			"headers":        values.Headers,
			"error":          values.Error,
			"error_location": errLocation,
		})

		switch _, ok := errors.GetCode(values.Error); {
		case values.Error == nil:
			logrus.Info("Successful")

		case ok:
			logrus.Warn("Registered error")

		case errors.Is(values.Error, echo.ErrNotFound),
			errors.Is(values.Error, echo.ErrMethodNotAllowed):
			logrus.Warn("Not found")

		default:
			logrus.Error("Unregistered internal error")
		}

		return nil
	},
})

// Validates and parses access token and puts claims to the context.
var AuthMW echo.MiddlewareFunc

func initAuthMiddleware(secret string) {
	AuthMW = echojwt.WithConfig(echojwt.Config{
		TokenLookup: "cookie:access",
		ContextKey:  "claims",
		ParseTokenFunc: func(c echo.Context, token string) (any, error) {
			var claims model.JWTClaims

			if _, err := jwt.ParseWithClaims(token, &claims, func(*jwt.Token) (any, error) {
				return []byte(secret), nil
			}); err != nil {
				if errors.Is(err, jwt.ErrTokenExpired) {
					return nil, errors.Wrapf(errors.ExpiredToken, "expired token, %v", err)
				}

				return nil, errors.Wrapf(errors.InvalidToken, "invalid token, %v", err)
			}

			return &claims, nil
		},
	})
}
