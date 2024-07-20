package model

import "github.com/sirupsen/logrus"

// Use cases
type IUsecase interface {
	CreateOrder(*CreateOrderReqUsecase) (*CreateOrderResUsecase, error)
	CancelOrder(*CancelOrderReqUsecase) (*CancelOrderResUsecase, error)
	ListOrders(*ListOrdersReqUsecase) (*ListOrdersResUsecase, error)
	GetOrderInfo(*GetOrderInfoReqUsecase) (*GetOrderInfoResUsecase, error)
	DeleteOrder(*DeleteOrderReqUsecase) (*DeleteOrderResUsecase, error)
}

// Logger
type ILogger interface {
	logrus.FieldLogger
}

// Storage
type IStorage interface {
	CreateOrder(*CreateOrderReqStorage) (*CreateOrderResStorage, error)
	CancelOrder(*CancelOrderReqStorage) (*CancelOrderResStorage, error)
	ListOrders(*ListOrdersReqStorage) (*ListOrdersResStorage, error)
	GetOrderInfo(*GetOrderInfoReqStorage) (*GetOrderInfoResStorage, error)
	DeleteOrder(*DeleteOrderReqStorage) (*DeleteOrderResStorage, error)
}
