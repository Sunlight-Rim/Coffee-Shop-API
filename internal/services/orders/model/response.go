package model

// CreateOrder

// easyjson:json
type CreateOrderResDelivery struct {
	OrderID uint64 `json:"id"`
}

type CreateOrderResUsecase struct {
	OrderID uint64
}

type CreateOrderResStorage struct {
	OrderID uint64
}
