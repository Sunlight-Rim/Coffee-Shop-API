package delivery

import (
	"coffeeshop-api/internal/server/middleware"

	"github.com/labstack/echo/v4"
)

func (h *handler) Register(group *echo.Group) {
	order := group.Group("/order", middleware.Auth)

	/*
	   swagger:route POST /api/order Orders CreateOrderRequest

	   Make a new coffee order.

	   	schemes: http
	   	responses:
	   		200: CreateOrderResponse
	   		default: ErrorResponse
	*/
	order.POST("", h.createOrder)
}
