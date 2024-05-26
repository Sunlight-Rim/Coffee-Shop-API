package rest

import "github.com/labstack/echo/v4"

func (h *handler) Register(group *echo.Group) {
	coffee := group.Group("/coffee")

	/*
		swagger:route GET /api/coffee/{id} Coffee GetCoffeeRequest

		Get one coffee info.

			schemes: http
			responses:
				200: GetCoffeeResponse
				default: ErrorResponse
	*/
	coffee.GET("/:id", h.getCoffee)
	/*
		swagger:route GET /api/coffee/list Coffee ListCoffeeRequest

		List coffee assortment.

			schemes: http
			responses:
				200: ListCoffeeResponse
				default: ErrorResponse
	*/
	coffee.GET("/list", h.listCoffee)
}
