package model

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
