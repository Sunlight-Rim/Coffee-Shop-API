package model

// ListCoffee

// easyjson:json
type ListCoffeeReqDelivery struct {
	Offset uint64 `json:"offset" query:"offset"`
}

type ListCoffeeReqUsecase struct {
	Offset uint64
}

type ListCoffeeReqStorage struct {
	Offset uint64
}

// GetCoffee

// easyjson:json
type GetCoffeeInfoReqDelivery struct {
	// in: path
	CoffeeID uint64 `json:"id" param:"id"`
}

type GetCoffeeInfoReqUsecase struct {
	CoffeeID uint64
}

type GetCoffeeInfoReqStorage struct {
	CoffeeID uint64
}

// ListToppings

// easyjson:json
type ListToppingsReqDelivery struct {
	Offset uint64 `json:"offset" query:"offset"`
}

type ListToppingsReqUsecase struct {
	Offset uint64
}

type ListToppingsReqStorage struct {
	Offset uint64
}
