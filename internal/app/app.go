package app

import (
	"context"
	"os/signal"
	"syscall"

	"coffeeshop-api/internal/server"
	"coffeeshop-api/internal/service/coffee"
	"coffeeshop-api/internal/service/orders"
	"coffeeshop-api/internal/service/users"
)

func init() {
	readFlags()
	readConfig()
}

func Start() {
	// Init server
	s := server.New()

	// Init services
	var (
		logger  = newLogger()
		storage = connectStorage()
		cache   = connectCache()
	)

	users.New(s.RoutesGroup, logger, storage, cache)
	orders.New(s.RoutesGroup, logger, storage)
	coffee.New(s.RoutesGroup, logger, storage)

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
