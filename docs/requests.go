package model

import (
	modelCoffee "coffeeshop-api/internal/services/coffee/model"
	modelOrders "coffeeshop-api/internal/services/orders/model"
	modelUsers "coffeeshop-api/internal/services/users/model"
)

// swagger:parameters SignupRequest
type SignupRequest struct {
	// in: body
	// required: true
	Body modelUsers.SignupReqDelivery
}

// swagger:parameters SigninRequest
type SigninRequest struct {
	// in: body
	// required: true
	Body modelUsers.SigninReqDelivery
}

// swagger:parameters ChangePasswordRequest
type ChangePasswordRequest struct {
	// in: body
	// required: true
	Body modelUsers.ChangePasswordReqDelivery
}

// swagger:parameters GetCoffeeInfoRequest
type GetCoffeeInfoRequest struct {
	// required: true
	modelCoffee.GetCoffeeInfoReqDelivery
}

// swagger:parameters ListCoffeeRequest
type ListCoffeeRequest struct {
	// in: query
	// required: true
	modelCoffee.ListCoffeeReqDelivery
}

// swagger:parameters ListToppingsRequest
type ListToppingsRequest struct {
	// in: query
	// required: true
	modelCoffee.ListToppingsReqDelivery
}

// swagger:parameters ListOrdersRequest
type ListOrdersRequest struct {
	// in: query
	// required: true
	modelOrders.ListOrdersReqDelivery
}

// swagger:parameters GetOrderInfoRequest
type GetOrderInfoRequest struct {
	// required: true
	modelOrders.GetOrderInfoReqDelivery
}

// swagger:parameters CreateOrderRequest
type CreateOrderRequest struct {
	// in: body
	// required: true
	Body modelOrders.CreateOrderReqDelivery
}

// swagger:parameters CancelOrderRequest
type CancelOrderRequest struct {
	// required: true
	modelOrders.CancelOrderReqDelivery
}

// swagger:parameters EmployeeCompleteOrderRequest
type EmployeeCompleteOrderRequest struct {
	// required: true
	modelOrders.EmployeeCompleteOrderReqDelivery
}
