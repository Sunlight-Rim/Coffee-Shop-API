package redis

import (
	"context"
	"fmt"
	"time"

	"github.com/pkg/errors"
	"github.com/redis/go-redis/v9"
)

type Options struct {
	Host        string
	Port        string
	Password    string
	Database    int
	PingTimeout time.Duration
}

// Connect and ping a redis.
func Connect(opts *Options) (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", opts.Host, opts.Port),
		Password: opts.Password,
		DB:       opts.Database,
	})

	ctx, cancel := context.WithTimeout(context.Background(), opts.PingTimeout*time.Second)
	defer cancel()

	if err := client.Ping(ctx).Err(); err != nil {
		return nil, errors.Wrap(err, "ping")
	}

	return client, nil
}
