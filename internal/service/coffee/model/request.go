package model

// GetCoffee

type GetCoffeeReq struct {
	CoffeeID uint64
}

type StorageGetCoffeeReq struct {
	CoffeeID uint64
}

// ListCoffee

type ListCoffeeReq struct {
	Offset uint64
}

type StorageListCoffeeReq struct {
	Offset uint64
}
