package coffee

import (
	"database/sql"

	"coffeeshop-api/internal/services/coffee/delivery"
	"coffeeshop-api/internal/services/coffee/infrastructure/storage"
	"coffeeshop-api/internal/services/coffee/usecase"

	"github.com/labstack/echo/v4"
)

// New service.
func New(group *echo.Group, postgres *sql.DB) {
	// Init usecase
	uc := usecase.New(
		// Init secondary adapters
		storage.New(postgres),
	)

	// Init primary adapters
	delivery.New(uc).Register(group)
}
