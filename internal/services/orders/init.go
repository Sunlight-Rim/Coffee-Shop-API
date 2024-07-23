package orders

import (
	"database/sql"

	"coffeeshop-api/internal/services/orders/delivery"
	"coffeeshop-api/internal/services/orders/infrastructure/storage"
	"coffeeshop-api/internal/services/orders/usecase"

	"github.com/labstack/echo/v4"
)

func New(group *echo.Group, postgres *sql.DB) {
	// Init usecase
	uc := usecase.New(
		// Init secondary adapters
		storage.New(postgres),
	)

	// // Init primary adapters
	hub := delivery.NewHubSSE()
	delivery.New(uc, hub).Register(group)
}
