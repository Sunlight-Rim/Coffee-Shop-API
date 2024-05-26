package model

import (
	restCoffee "coffeeshop-api/internal/service/coffee/delivery/rest"
	restUsers "coffeeshop-api/internal/service/users/delivery/rest"
)

// swagger:parameters SignupRequest
type SignupRequest struct {
	// in: body
	// required: true
	Body restUsers.SignupReq
}

// swagger:parameters SigninRequest
type SigninRequest struct {
	// in: body
	// required: true
	Body restUsers.SigninReq
}

// swagger:parameters ChangePasswordRequest
type ChangePasswordRequest struct {
	// in: body
	// required: true
	Body restUsers.ChangePasswordReq
}

// swagger:parameters GetCoffeeRequest
type GetCoffeeRequest struct {
	// required: true
	restCoffee.GetCoffeeReq
}

// swagger:parameters ListCoffeeRequest
type ListCoffeeRequest struct {
	restCoffee.ListCoffeeReq
}
