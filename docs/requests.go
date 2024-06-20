package model

import (
	modelCoffee "coffeeshop-api/internal/services/coffee/model"
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
