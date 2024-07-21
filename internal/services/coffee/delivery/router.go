package delivery

import "github.com/labstack/echo/v4"

func (h *handler) Register(group *echo.Group) {
	coffee := group.Group("/coffee")
	toppings := group.Group("/toppings")

	/*
		swagger:route GET /api/coffee Coffee ListCoffeeRequest

		List coffee assortment.

			schemes: http
			responses:
				200: ListCoffeeResponse
				default: ErrorResponse
	*/
	coffee.GET("/list", h.listCoffee)
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
		swagger:route GET /api/toppings/list Coffee ListToppingsRequest

		List toppings assortment.

			schemes: http
			responses:
				200: ListToppingsResponse
				default: ErrorResponse
	*/
	toppings.GET("/list", h.listToppings)
}
