package delivery

import (
	"coffeeshop-api/internal/server/middleware"

	"github.com/labstack/echo/v4"
)

func (h *handler) Register(group *echo.Group) {
	orders := group.Group("/orders", middleware.Auth)
	events := group.Group("/events", middleware.Auth)
	employee := group.Group("/employee", middleware.Auth)

	/*
		swagger:route GET /api/events/orders/statuses Orders OrdersStatusesRequest

		Server-Sent Events for instantly receiving orders statuses.

		schemes: http
		produces:
			- text/event-stream
		responses:
			200: OrdersStatusesResponse
			default: ErrorResponse
	*/
	events.GET("/orders/statuses", h.sseOrdersStatuses)
	/*
	   swagger:route GET /api/orders Orders ListOrdersRequest

	   List user orders.

	   	schemes: http
	   	responses:
	   		200: ListOrdersResponse
	   		default: ErrorResponse
	*/
	orders.GET("", h.listOrders)
	/*
	   swagger:route GET /api/orders/{id} Orders GetOrderInfoRequest

	   Get details about coffee order.

	   	schemes: http
	   	responses:
	   		200: GetOrderInfoResponse
	   		default: ErrorResponse
	*/
	orders.GET("/:id", h.getOrderInfo)
	/*
	   swagger:route POST /api/orders Orders CreateOrderRequest

	   Create a new coffee order.

	   	schemes: http
	   	responses:
	   		200: CreateOrderResponse
	   		default: ErrorResponse
	*/
	orders.POST("", h.createOrder)
	/*
	   swagger:route POST /api/orders/{id}/cancel Orders CancelOrderRequest

	   Set coffee order status to 'cancelled'.

	   	schemes: http
	   	responses:
	   		200: CancelOrderResponse
	   		default: ErrorResponse
	*/
	orders.POST("/:id/cancel", h.cancelOrder)
	/*
	   swagger:route POST /api/employee/orders/{id}/complete Orders EmployeeCompleteOrderRequest

	   Set coffee order status to 'ready to receive' after complete the cooking.
	   In this version receives the same token as a users endpoints.

	   	schemes: http
	   	responses:
	   		200: EmployeeCompleteOrderResponse
	   		default: ErrorResponse
	*/
	employee.POST("/orders/:id/complete", h.employeeCompleteOrder)
}
