package rest

import "github.com/labstack/echo/v4"

func (h *handler) Register(group *echo.Group) {
	coffee := group.Group("/coffee")

	/*
		swagger:route GET /api/coffee/{id} Coffee GetCoffeeInfoRequest

		Get one coffee information.

			schemes: http
			responses:
				200: GetCoffeeInfoResponse
				default: ErrorResponse
	*/
	coffee.GET("/:id", h.getCoffeeInfo)
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
