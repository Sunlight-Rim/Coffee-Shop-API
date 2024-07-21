package model

import (
	modelCoffee "coffeeshop-api/internal/services/coffee/model"
	modelOrders "coffeeshop-api/internal/services/orders/model"
	modelUsers "coffeeshop-api/internal/services/users/model"
	"coffeeshop-api/pkg/tools"
)

// General errors model. Response field is null.
// swagger:response ErrorResponse
type ErrorResponse struct {
	// in: body
	Body struct {
		// example: null
		Response any                 `json:"response"`
		Error    tools.ErrorResponse `json:"error"`
	} `json:"body"`
}

// A list of errors containing error codes and text descriptions.
// swagger:response ErrorsListResponse
type ErrorsListResponse struct {
	// in: body
	Body struct {
		Response struct {
			Language string `json:"language"`
		} `json:"error_code"`
	} `json:"body"`
}

// Informs whether the service is alive or not.
// swagger:response HealthResponse
// in: body
type _ string

// Returns user ID.
// swagger:response SignupResponse
type SignupResponse struct {
	// in: body
	Body struct {
		Response modelUsers.SignupResDelivery `json:"response"`
		// example: null
		Error any `json:"error"`
	} `json:"body"`
}

// Empty response.
// swagger:response SigninResponse
type SigninResponse struct {
	// in: body
	Body struct {
		// example: null
		Response any `json:"response"`
		// example: null
		Error any `json:"error"`
	} `json:"body"`
}

// Empty response.
// swagger:response RefreshResponse
type RefreshResponse struct {
	// in: body
	Body struct {
		// example: null
		Response any `json:"response"`
		// example: null
		Error any `json:"error"`
	} `json:"body"`
}

// Empty response.
// swagger:response SignoutResponse
type SignoutResponse struct {
	// in: body
	Body struct {
		// example: null
		Response any `json:"response"`
		// example: null
		Error any `json:"error"`
	} `json:"body"`
}

// Returns revoked tokens.
// swagger:response SignoutAllResponse
type SignoutAllResponse struct {
	// in: body
	Body struct {
		Response modelUsers.SignoutAllResDelivery `json:"response"`
		// example: null
		Error any `json:"error"`
	} `json:"body"`
}

// Returns user account data.
// swagger:response GetMeResponse
type GetMeResponse struct {
	// in: body
	Body struct {
		Response modelUsers.GetMeResDelivery `json:"response"`
		// example: null
		Error any `json:"error"`
	} `json:"body"`
}

// Empty response.
// swagger:response ChangePasswordResponse
type ChangePasswordResponse struct {
	// in: body
	Body struct {
		// example: null
		Response any `json:"response"`
		// example: null
		Error any `json:"error"`
	} `json:"body"`
}

// Returns user deleted account data.
// swagger:response DeleteMeResponse
type DeleteMeResponse struct {
	// in: body
	Body struct {
		Response modelUsers.DeleteMeResDelivery `json:"response"`
		// example: null
		Error any `json:"error"`
	} `json:"body"`
}

// Returns coffee info.
// swagger:response GetCoffeeInfoResponse
type GetCoffeeInfoResponse struct {
	// in: body
	Body struct {
		Response modelCoffee.GetCoffeeInfoResDelivery `json:"response"`
		// example: null
		Error any `json:"error"`
	} `json:"body"`
}

// Returns coffee list.
// swagger:response ListCoffeeResponse
type ListCoffeeResponse struct {
	// in: body
	Body struct {
		Response modelCoffee.ListCoffeeResDelivery `json:"response"`
		// example: null
		Error any `json:"error"`
	} `json:"body"`
}

// Returns toppings list.
// swagger:response ListToppingsResponse
type ListToppingsResponse struct {
	// in: body
	Body struct {
		Response modelCoffee.ListToppingsResDelivery `json:"response"`
		// example: null
		Error any `json:"error"`
	} `json:"body"`
}

// Returns orders statuses updates.
// swagger:response OrdersStatusesResponse
type OrdersStatusesResponse struct {
	// in: body
	Body struct {
		Response modelOrders.OrdersStatusesResDelivery `json:"response"`
		// example: null
		Error any `json:"error"`
	} `json:"body"`
}

// Returns orders list.
// swagger:response ListOrdersResponse
type ListOrdersResponse struct {
	// in: body
	Body struct {
		Response modelOrders.ListOrdersResDelivery `json:"response"`
		// example: null
		Error any `json:"error"`
	} `json:"body"`
}

// Returns order information.
// swagger:response GetOrderInfoResponse
type GetOrderInfoResponse struct {
	// in: body
	Body struct {
		Response modelOrders.GetOrderInfoResDelivery `json:"response"`
		// example: null
		Error any `json:"error"`
	} `json:"body"`
}

// Returns toppings list.
// swagger:response CreateOrderResponse
type CreateOrderResponse struct {
	// in: body
	Body struct {
		Response modelOrders.CreateOrderResDelivery `json:"response"`
		// example: null
		Error any `json:"error"`
	} `json:"body"`
}

// Order was cancelled.
// swagger:response CancelOrderResponse
type CancelOrderResponse struct {
	// in: body
	Body struct {
		Response modelOrders.CancelOrderResDelivery `json:"response"`
		// example: null
		Error any `json:"error"`
	} `json:"body"`
}

// Order was completed.
// swagger:response EmployeeCompleteOrderResponse
type EmployeeCompleteOrderResponse struct {
	// in: body
	Body struct {
		Response modelOrders.EmployeeCompleteOrderResDelivery `json:"response"`
		// example: null
		Error any `json:"error"`
	} `json:"body"`
}
