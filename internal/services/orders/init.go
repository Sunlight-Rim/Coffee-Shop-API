package orders

import (
	"database/sql"

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
	// delivery.New(uc).Register(group)
}
