package app

import (
	"database/sql"
	"flag"

	"coffeeshop-api/pkg/logger"
	"coffeeshop-api/pkg/postgres"
	"coffeeshop-api/pkg/redis"
	"coffeeshop-api/pkg/tools"

	goRedis "github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var flags struct {
	configPath string
}

// Read startup flags.
func readFlags() {
	flags.configPath = *flag.String("config", "configs/config.json", "path to configuration file")

	flag.Parse()
}

// Read configuration file.
func readConfig() {
	viper.SetConfigFile(flags.configPath)

	if err := viper.ReadInConfig(); err != nil {
		logrus.Fatalf("read config: %v", err)
	}
}

// Init application tools.
func initTools() {
	tools.Init(tools.Options{
		CookieDomain:   viper.GetString("cookie.domain"),
		CookieSecure:   viper.GetBool("cookie.secure"),
		CookieHttpOnly: viper.GetBool("cookie.http_only"),
	})
}

// Initialize logger.
func initLogger() {
	logger.Init(
		viper.GetString("logger.level"),
		viper.GetBool("logger.json"),
	)
}

// Connect main storage.
func connectStorage() *sql.DB {
	db, err := postgres.Connect(&postgres.Options{
		Host:            viper.GetString("database.postgres.host"),
		Port:            viper.GetString("database.postgres.port"),
		User:            viper.GetString("database.postgres.user"),
		Password:        viper.GetString("database.postgres.password"),
		DBName:          viper.GetString("database.postgres.database"),
		SSLMode:         viper.GetString("database.postgres.sslmode"),
		MaxOpenConns:    viper.GetInt("database.postgres.max_open_conns"),
		MaxIdleConns:    viper.GetInt("database.postgres.max_idle_conns"),
		ConnMaxLifetime: viper.GetDuration("database.postgres.conn_max_lifetime"),
		ConnMaxIdleTime: viper.GetDuration("database.postgres.conn_max_idle_time"),
		PingTimeout:     viper.GetDuration("database.postgres.ping_timeout_sec"),
	})
	if err != nil {
		logrus.Fatalf("connect to postgres: %v", err)
	}

	return db
}

// Connect cache.
func connectCache() *goRedis.Client {
	client, err := redis.Connect(&redis.Options{
		Host:        viper.GetString("database.redis.host"),
		Port:        viper.GetString("database.redis.port"),
		Password:    viper.GetString("database.redis.password"),
		Database:    viper.GetInt("database.redis.database"),
		PingTimeout: viper.GetDuration("database.redis.ping_timeout_sec"),
	})
	if err != nil {
		logrus.Fatalf("connect to redis: %v", err)
	}

	return client
}
