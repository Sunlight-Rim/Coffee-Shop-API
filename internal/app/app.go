package app

import (
	"context"
	"os/signal"
	"syscall"

	"coffeeshop-api/internal/server"
	"coffeeshop-api/internal/services/coffee"
	"coffeeshop-api/internal/services/orders"
	"coffeeshop-api/internal/services/users"
)

func init() {
	readFlags()
	readConfig()
	initTools()
}

func Start() {
	// Init infrastructure
	var (
		logger  = newLogger()
		storage = connectStorage()
		cache   = connectCache()
	)

	// Init server
	s := server.New(logger)

	// Init services
	users.New(s.ApiGroup, logger, storage, cache)
	coffee.New(s.ApiGroup, logger, storage)
	orders.New(s.ApiGroup, logger, storage)

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
	logger.Fatalf("Shutdown: %v", s.Shutdown(ctx))
}
