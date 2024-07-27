package model

import "time"

// OrdersStatuses

// easyjson:json
type OrdersStatusesResDelivery struct {
	OrderID        uint64    `json:"id"`
	OrderCreatedAt time.Time `json:"created_at"`
	OrderStatus    string    `json:"status"`
}

type OrdersStatusesResUsecase struct {
	OrderID        uint64
	OrderCreatedAt time.Time
	OrderStatus    string
}

type OrdersStatusesResStorage struct {
	OrderID        uint64
	OrderCreatedAt time.Time
	OrderStatus    string
}

// ListOrders

type ListOrdersOrder struct {
	OrderID        uint64    `json:"id"`
	OrderCreatedAt time.Time `json:"created_at"`
}

// easyjson:json
type ListOrdersResDelivery struct {
	Orders []ListOrdersOrder `json:"orders"`
}

type ListOrdersResUsecase struct {
	Orders []ListOrdersOrder
}

type ListOrdersResStorage struct {
	Orders []ListOrdersOrder
}

// GetOrderInfo

type GetOrderInfoOrderItem struct {
	CoffeeID    uint64  `json:"coffee_id"`
	CoffeeTitle string  `json:"coffee_title"`
	CoffeeImage string  `json:"coffee_image"`
	Topping     *string `json:"topping"`
}

type GetOrderInfoOrder struct {
	OrderID   uint64                  `json:"id"`
	Status    string                  `json:"status"`
	Address   string                  `json:"address"`
	CreatedAt time.Time               `json:"created_at"`
	Items     []GetOrderInfoOrderItem `json:"items"`
}

// easyjson:json
type GetOrderInfoResDelivery struct {
	Order GetOrderInfoOrder `json:"order"`
}

type GetOrderInfoResUsecase struct {
	Order GetOrderInfoOrder
}

type GetOrderInfoResStorage struct {
	Order GetOrderInfoOrder
}

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

// CancelOrder

// easyjson:json
type CancelOrderResDelivery struct {
	OrderID uint64 `json:"id"`
}

type CancelOrderResUsecase struct {
	OrderID         uint64
	OrderCustomerID uint64
	OrderCreatedAt  time.Time
	OrderStatus     string
}

type CancelOrderResStorage struct {
	OrderID         uint64
	OrderCustomerID uint64
	OrderCreatedAt  time.Time
	OrderStatus     string
}

// EmployeeCompleteOrder

// easyjson:json
type EmployeeCompleteOrderResDelivery struct {
	OrderID uint64 `json:"id"`
}

type EmployeeCompleteOrderResUsecase struct {
	OrderID         uint64
	OrderCustomerID uint64
	OrderCreatedAt  time.Time
	OrderStatus     string
}

type EmployeeCompleteOrderResStorage struct {
	OrderID         uint64
	OrderCustomerID uint64
	OrderCreatedAt  time.Time
	OrderStatus     string
}
