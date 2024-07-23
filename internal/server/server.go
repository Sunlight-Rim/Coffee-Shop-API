package server

import (
	"context"
	"fmt"
	"net/http"

	"coffeeshop-api/internal/server/middleware"
	"coffeeshop-api/pkg/errors"
	"coffeeshop-api/pkg/tools"

	"github.com/labstack/echo/v4"
	echomw "github.com/labstack/echo/v4/middleware"
	"github.com/spf13/viper"
)

type server struct {
	ApiGroup *echo.Group
	echo     *echo.Echo
}

// New application server, init middlewares and basic application routes.
func New() *server {
	// Init middlewares
	middleware.Init(viper.GetString("token.secret"))

	// Init echo
	echo := echo.New()
	echo.Use(echomw.Recover(), echomw.RequestID(), middleware.Logger, middleware.CORS)
	echo.HTTPErrorHandler = errorHandler

	// Init server
	s := &server{
		ApiGroup: echo.Group("/api"),
		echo:     echo,
	}

	s.register()

	return s
}

// Start server.
func (s *server) Start() error {
	if err := s.echo.Start(fmt.Sprintf(
		"%s:%s",
		viper.GetString("server.application.host"),
		viper.GetString("server.application.port"),
	)); err != nil {
		return errors.Wrap(err, "start")
	}

	return nil
}

// Shutdown server.
func (s *server) Shutdown(ctx context.Context) error {
	if err := s.echo.Shutdown(ctx); err != nil {
		return errors.Wrap(err, "shutdown")
	}

	return nil
}

// Custom Echo error handler.
func errorHandler(err error, c echo.Context) {
	if httpError := new(echo.HTTPError); errors.As(err, &httpError) {
		switch httpError.Code {
		case http.StatusNotFound:
			tools.SendResponse(c, nil, errors.NotFound)

		case http.StatusMethodNotAllowed:
			tools.SendResponse(c, nil, errors.MethodNotAllowed)
		}
	}
}
