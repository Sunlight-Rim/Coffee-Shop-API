package orders

import (
	"database/sql"

	// "coffeeshop-api/internal/service/orders/delivery/rest"
	// "coffeeshop-api/internal/service/orders/delivery/websocket"
	// "coffeeshop-api/internal/service/orders/infrastructure/storage"
	// "coffeeshop-api/internal/service/orders/usecase"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

func New(group *echo.Group, logrus *logrus.Logger, postgres *sql.DB) {
	// // Init usecase
	// uc := usecase.New(
	// 	// Init secondary adapters
	// 	logrus,
	// 	storage.New(postgres),
	// )

	// // Init primary adapters
	// rest.New(uc).Register(group)
	// websocket.New(uc).Register(group)
}
