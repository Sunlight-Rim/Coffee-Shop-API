package server

import (
	"fmt"

	"github.com/labstack/echo-contrib/pprof"
	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
)

// Start profiler.
func (*server) StartProfiler() error {
	e := echo.New()

	e.HideBanner = true

	pprof.Register(e)

	return e.Start(fmt.Sprintf(
		"%s:%s",
		viper.GetString("server.pprof.host"),
		viper.GetString("server.pprof.port"),
	))
}
