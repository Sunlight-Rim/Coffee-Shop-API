package orders

import (
	"database/sql"

	"coffeeshop-api/internal/services/orders/delivery"
	"coffeeshop-api/internal/services/orders/infrastructure/storage"
	"coffeeshop-api/internal/services/orders/usecase"

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

	// // Init primary adapters
	hub := delivery.NewHubSSE()
	go hub.Run()

	delivery.New(uc, hub).Register(group)
}
