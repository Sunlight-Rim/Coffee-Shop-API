package model

// GetCoffee

// easyjson:json
type DeliveryGetCoffeeInfoReq struct {
	// in: path
	CoffeeID uint64 `json:"id" param:"id"`
}

type UsecaseGetCoffeeInfoReq struct {
	CoffeeID uint64
}

type StorageGetCoffeeInfoReq struct {
	CoffeeID uint64
}

// ListCoffee

// easyjson:json
type DeliveryListCoffeeReq struct {
	Offset uint64 `json:"offset" query:"offset"`
}

type UsecaseListCoffeeReq struct {
	Offset uint64
}

type StorageListCoffeeReq struct {
	Offset uint64
}
