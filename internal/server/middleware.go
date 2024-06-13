package server

import (
	"net/http"

	"coffeeshop-api/internal/service/users/model"
	"coffeeshop-api/pkg/errors"
	"coffeeshop-api/pkg/tools"

	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sirupsen/logrus"
)

var (
	// LoggerMW provides response logger.
	LoggerMW echo.MiddlewareFunc

	// AuthMW validates and parses access token and puts the claims to the context.
	AuthMW echo.MiddlewareFunc
)

// initMiddlewares initializes the middlewares of application server.
func (s *server) initMiddlewares(logger logrus.FieldLogger, tokenSecret string) {
	LoggerMW = newLoggerMW(logger)
	AuthMW = newAuthMW([]byte(tokenSecret))
}

func newLoggerMW(logger logrus.FieldLogger) echo.MiddlewareFunc {
	return middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
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

			logger := logger.WithFields(logrus.Fields{
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
				logger.Info("Successful")

			case ok:
				logger.Warn("Registered error")

			case errors.Is(values.Error, echo.ErrNotFound),
				errors.Is(values.Error, echo.ErrMethodNotAllowed):
				logger.Warn("Not found")

			default:
				logger.Error("Unregistered internal error")
			}

			return nil
		},
	})
}

func newAuthMW(tokenSecret []byte) echo.MiddlewareFunc {
	return echojwt.WithConfig(echojwt.Config{
		TokenLookup: "cookie:access",
		ContextKey:  "claims",
		ParseTokenFunc: func(c echo.Context, token string) (any, error) {
			var claims model.JWTClaims

			if _, err := jwt.ParseWithClaims(token, &claims, func(*jwt.Token) (any, error) {
				return tokenSecret, nil
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

// Custom Echo error handler.
func errorHandler(err error, c echo.Context) {
	if httpError := new(echo.HTTPError); errors.As(err, &httpError) {
		switch {
		case httpError.Code == http.StatusNotFound:
			tools.SendResponse(c, nil, errors.NotFound)

		case httpError.Code == http.StatusMethodNotAllowed:
			tools.SendResponse(c, nil, errors.MethodNotAllowed)

		case httpError.Message == echojwt.ErrJWTMissing.Message:
			tools.SendResponse(c, nil, errors.MissingToken)
		}
	}
}
