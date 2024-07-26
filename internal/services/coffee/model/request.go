package model

// ListCoffees

// easyjson:json
type ListCoffeesReqDelivery struct {
	Offset uint64 `json:"offset" query:"offset"`
}

type ListCoffeesReqUsecase struct {
	Offset uint64
}

type ListCoffeesReqStorage struct {
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
