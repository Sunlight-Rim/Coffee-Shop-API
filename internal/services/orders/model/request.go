package model

import (
	"coffeeshop-api/pkg/errors"
)

// OrdersStatuses

// easyjson:json
type OrdersStatusesReqDelivery struct {
	UserID uint64 `json:"-"`
}

type OrdersStatusesReqUsecase struct {
	UserID uint64
}

// ListOrders

// easyjson:json
type ListOrdersReqDelivery struct {
	UserID uint64 `json:"-"`
	Offset uint64 `json:"offset" query:"offset"`
}

type ListOrdersReqUsecase struct {
	UserID uint64
	Offset uint64
}

type ListOrdersReqStorage struct {
	UserID uint64
	Offset uint64
}

// GetOrderInfo

// easyjson:json
type GetOrderInfoReqDelivery struct {
	UserID uint64 `json:"-"`
	// in: path
	OrderID uint64 `json:"id" param:"id"`
}

type GetOrderInfoReqUsecase struct {
	UserID  uint64
	OrderID uint64
}

type GetOrderInfoReqStorage struct {
	UserID  uint64
	OrderID uint64
}

// CreateOrder

// easyjson:json
type CreateOrderOrderItem struct {
	CoffeeID uint64  `json:"coffee_id"`
	Topping  *string `json:"topping"`
}

// easyjson:json
type CreateOrderReqDelivery struct {
	UserID  uint64                 `json:"-"`
	Address string                 `json:"address"`
	Items   []CreateOrderOrderItem `json:"items"`
}

type CreateOrderReqUsecase struct {
	UserID  uint64
	Address string
	Items   []CreateOrderOrderItem
}

type CreateOrderReqStorage struct {
	UserID  uint64
	Address string
	Items   []CreateOrderOrderItem
}

func (req *CreateOrderReqUsecase) Validate() error {
	if lenItems := len(req.Items); lenItems < 1 || lenItems > 10 {
		return errors.Wrapf(errors.InvalidRequestContent, "incorrect items count: %d", len(req.Items))
	}

	if req.Address == "" {
		return errors.Wrap(errors.InvalidRequestContent, "empty address")
	}

	return nil
}

// CancelOrder

// easyjson:json
type CancelOrderReqDelivery struct {
	UserID uint64 `json:"-"`
	// in: path
	OrderID uint64 `json:"id" param:"id"`
}

type CancelOrderReqUsecase struct {
	UserID  uint64
	OrderID uint64
}

type CancelOrderReqStorage struct {
	UserID  uint64
	OrderID uint64
}

// EmployeeCompleteOrder

// easyjson:json
type EmployeeCompleteOrderReqDelivery struct {
	// in: path
	OrderID uint64 `json:"id" param:"id"`
}

type EmployeeCompleteOrderReqUsecase struct {
	OrderID uint64
}

type EmployeeCompleteOrderReqStorage struct {
	OrderID uint64
}
