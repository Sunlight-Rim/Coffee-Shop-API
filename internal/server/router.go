/*
Package app Coffee Shop API

Example Go application with Clean Architecture pattern.
Provides user registration, login and ordering of coffee.

		version: 0.1
		schemes: http, https
		host: localhost:1337
	    consumes:
	        - application/json

	    produces:
	        - application/json

	    securityDefinitions:
	    accessToken:
	        type: apiKey
	        in: cookie
	        name: access-token
	        description: authorization access JWT stored in a cookie.

swagger:meta
*/
package server

import (
	"net/http"

	"coffeeshop-api/pkg/errors"

	"github.com/labstack/echo/v4"
)

// Register application utility routes.
func registerRoutes(group *echo.Group) {
	/*
		swagger:route GET /api/errors Errors null

		List of API errors.

			schemes: https
			responses:
				200: ErrorsListResponse
				default: ErrorResponse
	*/
	group.GET("/errors", func(c echo.Context) error {
		return c.JSONBlob(http.StatusOK, errors.ResponseList)
	})
	/*
		swagger:route GET /api/health Health null

		Health check.

			schemes: https
			responses:
				200: HealthResponse
				default: ErrorResponse
	*/
	group.GET("/health", func(c echo.Context) error {
		return c.JSONBlob(http.StatusOK, []byte("Success!"))
	})
}
