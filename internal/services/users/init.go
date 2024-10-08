package users

import (
	"database/sql"

	"coffeeshop-api/internal/services/users/delivery"
	"coffeeshop-api/internal/services/users/infrastructure/cache"
	"coffeeshop-api/internal/services/users/infrastructure/storage"
	"coffeeshop-api/internal/services/users/infrastructure/token"
	"coffeeshop-api/internal/services/users/usecase"

	"github.com/labstack/echo/v4"
	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
)

// New service.
func New(group *echo.Group, postgres *sql.DB, redis *redis.Client) {
	// Init usecase
	uc := usecase.New(
		// Init secondary adapters
		storage.New(postgres),
		cache.New(redis),
		token.New(
			viper.GetString("token.secret"),
			viper.GetDuration("token.access_exp"),
			viper.GetDuration("token.refresh_exp"),
		),
	)

	// Init primary adapters
	delivery.New(uc).Register(group)
}
