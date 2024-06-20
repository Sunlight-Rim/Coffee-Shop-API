package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq"
	"github.com/pkg/errors"
)

type ConnectionOptions struct {
	Host            string
	Port            string
	User            string
	Password        string
	DBName          string
	SSLMode         string
	MaxOpenConns    int
	MaxIdleConns    int
	ConnMaxLifetime time.Duration
	ConnMaxIdleTime time.Duration
	PingTimeout     time.Duration
}

// Connect and ping a postgres.
func Connect(opts *ConnectionOptions) (*sql.DB, error) {
	db, err := sql.Open("postgres", fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		opts.Host,
		opts.Port,
		opts.User,
		opts.Password,
		opts.DBName,
		opts.SSLMode,
	))
	if err != nil {
		return nil, errors.Wrap(err, "connect")
	}

	db.SetMaxOpenConns(opts.MaxOpenConns)
	db.SetMaxIdleConns(opts.MaxIdleConns)
	db.SetConnMaxLifetime(opts.ConnMaxLifetime * time.Second)
	db.SetConnMaxIdleTime(opts.ConnMaxIdleTime * time.Second)

	ctx, cancel := context.WithTimeout(context.Background(), opts.PingTimeout*time.Second)
	defer cancel()

	if err = db.PingContext(ctx); err != nil {
		return nil, errors.Wrap(err, "ping")
	}

	return db, nil
}
