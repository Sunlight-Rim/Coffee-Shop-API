package model

import (
	modelCoffee "coffeeshop-api/internal/service/coffee/model"
	modelUsers "coffeeshop-api/internal/service/users/model"
)

// swagger:parameters SignupRequest
type SignupRequest struct {
	// in: body
	// required: true
	Body modelUsers.DeliverySignupReq
}

// swagger:parameters SigninRequest
type SigninRequest struct {
	// in: body
	// required: true
	Body modelUsers.DeliverySigninReq
}

// swagger:parameters ChangePasswordRequest
type ChangePasswordRequest struct {
	// in: body
	// required: true
	Body modelUsers.DeliveryChangePasswordReq
}

// swagger:parameters GetCoffeeInfoRequest
type GetCoffeeInfoRequest struct {
	// required: true
	modelCoffee.DeliveryGetCoffeeInfoReq
}

// swagger:parameters ListCoffeeRequest
type ListCoffeeRequest struct {
	// in: query
	// required: true
	modelCoffee.DeliveryListCoffeeReq
}
