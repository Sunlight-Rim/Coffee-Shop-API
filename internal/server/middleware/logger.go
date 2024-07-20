package middleware

import (
	"coffeeshop-api/pkg/errors"

	"github.com/labstack/echo/v4"
	echomw "github.com/labstack/echo/v4/middleware"
	"github.com/sirupsen/logrus"
)

// Logger middleware provides response logger.
var Logger echo.MiddlewareFunc

func initLogger(logger logrus.FieldLogger) {
	Logger = echomw.RequestLoggerWithConfig(echomw.RequestLoggerConfig{
		LogURI:       true,
		LogMethod:    true,
		LogStatus:    true,
		LogLatency:   true,
		LogRemoteIP:  true,
		LogRequestID: true,
		LogUserAgent: true,
		LogError:     true,
		LogHeaders:   []string{"*"},

		LogValuesFunc: func(c echo.Context, values echomw.RequestLoggerValues) error {
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
