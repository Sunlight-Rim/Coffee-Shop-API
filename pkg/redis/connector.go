package redis

import (
	"context"
	"fmt"

	"github.com/pkg/errors"
	"github.com/redis/go-redis/v9"
)

// Connect and ping a redis.
func Connect(host, port, password string, database int) (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", host, port),
		Password: password,
		DB:       database,
	})

	if err := client.Ping(context.Background()).Err(); err != nil {
		return nil, errors.Wrap(err, "ping")
	}

	return client, nil
}
