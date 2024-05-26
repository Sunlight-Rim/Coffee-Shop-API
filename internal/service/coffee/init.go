package coffee

import (
	"database/sql"

	"coffeeshop-api/internal/service/coffee/delivery/rest"
	"coffeeshop-api/internal/service/coffee/infrastructure/storage"
	"coffeeshop-api/internal/service/coffee/usecase"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

func New(group *echo.Group, logrus *logrus.Logger, postgres *sql.DB) {
	// Init usecase
	uc := usecase.New(
		// Init secondary adapters
		logrus,
		storage.New(postgres),
	)

	// Init primary adapters
	rest.New(uc).Register(group)
}
