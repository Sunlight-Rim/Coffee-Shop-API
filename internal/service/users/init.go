package users

import (
	"database/sql"

	"coffeeshop-api/internal/service/users/delivery/rest"
	"coffeeshop-api/internal/service/users/infrastructure/cache"
	"coffeeshop-api/internal/service/users/infrastructure/storage"
	"coffeeshop-api/internal/service/users/infrastructure/token"
	"coffeeshop-api/internal/service/users/usecase"

	"github.com/labstack/echo/v4"
	goRedis "github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func New(group *echo.Group, logrus *logrus.Logger, postgres *sql.DB, redis *goRedis.Client) {
	// Init usecase
	uc := usecase.New(
		// Init secondary adapters
		logrus,
		storage.New(postgres),
		cache.New(redis),
		token.New(
			viper.GetString("token.secret"),
			viper.GetDuration("token.access_exp"),
			viper.GetDuration("token.refresh_exp"),
		),
	)

	// Init primary adapters
	rest.New(uc).Register(group)
}
