package app

import (
	"context"
	"os/signal"
	"syscall"

	"coffeeshop-api/internal/server"
	"coffeeshop-api/internal/services/coffee"
	"coffeeshop-api/internal/services/orders"
	"coffeeshop-api/internal/services/users"

	logger "github.com/sirupsen/logrus"
)

func init() {
	readFlags()
	readConfig()
	initLogger()
	initTools()
}

func Start() {
	// Init infrastructure
	var (
		storage = connectStorage()
		cache   = connectCache()
	)

	// Run migrations
	runMigrations(storage)

	// Init tracer
	tracerProvider := initTracer()

	// Init server
	s := server.New()

	// Init services
	users.New(s.ApiGroup, storage, cache)
	coffee.New(s.ApiGroup, storage)
	orders.New(s.ApiGroup, storage)

	// Start server
	go func() {
		logger.Fatalf("Server: %v", s.Start())
	}()

	go func() {
		logger.Fatalf("Profiler: %v", s.StartProfiler())
	}()

	// Graceful shutdown
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	<-ctx.Done()
	ctx = context.Background()

	if err := storage.Close(); err != nil {
		logger.Errorf("Close DB: %s", err)
	}
	if err := cache.Close(); err != nil {
		logger.Errorf("Close cache: %s", err)
	}
	if err := tracerProvider.Shutdown(ctx); err != nil {
		logger.Errorf("Tracing shutdown: %s", err)
	}

	logger.Fatalf("Shutdown: %v", s.Shutdown(ctx))
}
