package model

// easyjson:json
type Order struct {
	CoffeeID uint   `json:"coffee_id"`
	Topping  string `json:"topping"`
}

// CreateOrder

// easyjson:json
type CreateOrderReqDelivery struct {
	Orders []Order `json:"orders"`
}

type CreateOrderReqUsecase struct {
	Orders []Order
}

type CreateOrderReqStorage struct {
	Orders []Order
}
