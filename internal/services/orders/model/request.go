package model

import "coffeeshop-api/pkg/errors"

// easyjson:json
type OrderItem struct {
	CoffeeID uint64 `json:"coffee_id"`
	Topping  string `json:"topping"`
}

// CreateOrder

// easyjson:json
type CreateOrderReqDelivery struct {
	UserID  uint64
	Address string      `json:"address"`
	Items   []OrderItem `json:"items"`
}

type CreateOrderReqUsecase struct {
	UserID  uint64
	Address string
	Items   []OrderItem
}

type CreateOrderReqStorage struct {
	UserID  uint64
	Address string
	Items   []OrderItem
}

func (req *CreateOrderReqUsecase) Validate() error {
	if lenItems := len(req.Items); lenItems < 0 || lenItems > 10 {
		return errors.Wrapf(errors.InvalidRequestContent, "incorrect items count: %d", len(req.Items))
	}

	if req.Address == "" {
		return errors.Wrap(errors.InvalidRequestContent, "empty address")
	}

	return nil
}
