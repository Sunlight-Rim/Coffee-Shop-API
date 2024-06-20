package middleware

import (
	"github.com/labstack/echo/v4"
	echomw "github.com/labstack/echo/v4/middleware"
)

// CORS middleware.
var CORS echo.MiddlewareFunc

func initCORS() {
	CORS = echomw.CORSWithConfig(echomw.CORSConfig{
		AllowOrigins: []string{
			"http://localhost:8080",
			"http://localhost:1337",
		},
		AllowHeaders:     []string{echo.HeaderContentType, "*"},
		AllowCredentials: true,
	})
}
