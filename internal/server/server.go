package server

import (
	"context"
	"fmt"

	"coffeeshop-api/pkg/errors"

	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
)

type server struct {
	RoutesGroup *echo.Group
	echo        *echo.Echo
}

// New application server.
func New() *server {
	echo := echo.New()

	s := &server{
		RoutesGroup: echo.Group("/api"),
		echo:        echo,
	}

	s.initMiddlewares()
	s.register()

	return s
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
