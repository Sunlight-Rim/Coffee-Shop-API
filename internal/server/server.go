package server

import (
	"context"
	"fmt"

	"coffeeshop-api/pkg/errors"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type server struct {
	RoutesGroup *echo.Group
	echo        *echo.Echo
}

// New application server.
func New(logger logrus.FieldLogger) *server {
	echo := echo.New()

	s := &server{
		RoutesGroup: echo.Group("/api"),
		echo:        echo,
	}

	s.initMiddlewares(logger, viper.GetString("token.secret"))

	s.echo.Use(middleware.Recover(), middleware.CORS(), middleware.RequestID(), LoggerMW)
	s.echo.HTTPErrorHandler = errorHandler

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
