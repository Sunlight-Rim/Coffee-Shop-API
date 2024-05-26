package server

import (
	"context"
	"fmt"
	"net/http"

	"coffeeshop-api/pkg/errors"
	"coffeeshop-api/pkg/tools"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/spf13/viper"
)

type server struct {
	APIGroup *echo.Group
	echo     *echo.Echo
}

// New application server.
func New() *server {
	s := server{
		echo: echo.New(),
	}

	initAuthMiddleware(viper.GetString("token.secret"))

	s.echo.Use( /*middleware.Recover(),*/ middleware.CORS(), middleware.RequestID(), LoggerMW)
	s.echo.HTTPErrorHandler = errorHandler

	s.APIGroup = s.echo.Group("/api")
	registerRoutes(s.APIGroup)

	return &s
}

// Start server.
func (s *server) Start() error {
	return s.echo.Start(fmt.Sprintf(
		"%s:%s",
		viper.GetString("server.application.host"),
		viper.GetString("server.application.port"),
	))
}

// Shutdown server.
func (s *server) Shutdown(ctx context.Context) error {
	if err := s.echo.Shutdown(ctx); err != nil {
		return errors.Wrap(err, "echo shutdown")
	}

	return nil
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

		case httpError.Message == echojwt.ErrJWTInvalid.Message:
			tools.SendResponse(c, nil, errors.InvalidToken)
		}
	}
}
