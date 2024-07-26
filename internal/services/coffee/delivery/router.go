package delivery

import "github.com/labstack/echo/v4"

func (h *handler) Register(group *echo.Group) {
	coffees := group.Group("/coffees")
	toppings := group.Group("/toppings")

	/*
		swagger:route GET /api/coffees Coffee ListCoffeesRequest

		List coffee assortment.

			schemes: http
			responses:
				200: ListCoffeesResponse
				default: ErrorResponse
	*/
	coffees.GET("", h.listCoffees)
	/*
		swagger:route GET /api/coffees/{id} Coffee GetCoffeeInfoRequest

		Get one coffee information.

			schemes: http
			responses:
				200: GetCoffeeInfoResponse
				default: ErrorResponse
	*/
	coffees.GET("/:id", h.getCoffeeInfo)
	/*
		swagger:route GET /api/toppings Coffee ListToppingsRequest

		List toppings assortment.

			schemes: http
			responses:
				200: ListToppingsResponse
				default: ErrorResponse
	*/
	toppings.GET("", h.listToppings)
}
