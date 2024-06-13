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
	        name: Authorization
	        in: cookie
	        description: JWT authorization token stored in a cookie.

swagger:meta
*/
package server

import (
	"net/http"

	"coffeeshop-api/pkg/errors"

	"github.com/labstack/echo/v4"
)

// Register application utility routes.
func (s *server) register() {
	/*
		swagger:route GET /api/errors Web null

		List of API errors.

			schemes: https
			responses:
				200: ErrorsListResponse
				default: ErrorResponse
	*/
	s.RoutesGroup.GET("/errors", func(c echo.Context) error {
		return c.JSONBlob(http.StatusOK, errors.ResponseList)
	})
	/*
		swagger:route GET /api/health Web null

		Health check.

			schemes: https
			responses:
				200: HealthResponse
				default: ErrorResponse
	*/
	s.RoutesGroup.GET("/health", func(c echo.Context) error {
		return c.JSONBlob(http.StatusOK, []byte("Success!"))
	})
}
