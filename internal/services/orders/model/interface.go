package model

import "github.com/sirupsen/logrus"

// Use cases
type IUsecase interface {
	ListOrders(*ListOrdersReqUsecase) (*ListOrdersResUsecase, error)
	GetOrderInfo(*GetOrderInfoReqUsecase) (*GetOrderInfoResUsecase, error)
	CreateOrder(*CreateOrderReqUsecase) (*CreateOrderResUsecase, error)
	CancelOrder(*CancelOrderReqUsecase) (*CancelOrderResUsecase, error)
	EmployeeCompleteOrder(*EmployeeCompleteOrderReqUsecase) (*EmployeeCompleteOrderResUsecase, error)
}

// Logger
type ILogger interface {
	logrus.FieldLogger
}

// Storage
type IStorage interface {
	ListOrders(*ListOrdersReqStorage) (*ListOrdersResStorage, error)
	GetOrderInfo(*GetOrderInfoReqStorage) (*GetOrderInfoResStorage, error)
	CreateOrder(*CreateOrderReqStorage) (*CreateOrderResStorage, error)
	CancelOrder(*CancelOrderReqStorage) (*CancelOrderResStorage, error)
	EmployeeCompleteOrder(*EmployeeCompleteOrderReqStorage) (*EmployeeCompleteOrderResStorage, error)
}
