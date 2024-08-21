package model

import "context"

// Use cases
type IUsecase interface {
	ListOrders(context.Context, *ListOrdersReqUsecase) (*ListOrdersResUsecase, error)
	GetOrderInfo(context.Context, *GetOrderInfoReqUsecase) (*GetOrderInfoResUsecase, error)
	CreateOrder(context.Context, *CreateOrderReqUsecase) (*CreateOrderResUsecase, error)
	CancelOrder(context.Context, *CancelOrderReqUsecase) (*CancelOrderResUsecase, error)
	EmployeeCompleteOrder(context.Context, *EmployeeCompleteOrderReqUsecase) (*EmployeeCompleteOrderResUsecase, error)
}

// Storage
type IStorage interface {
	ListOrders(context.Context, *ListOrdersReqStorage) (*ListOrdersResStorage, error)
	GetOrderInfo(context.Context, *GetOrderInfoReqStorage) (*GetOrderInfoResStorage, error)
	CheckCoffeeIDsExists(context.Context, *CheckCoffeeIDsExistsReqStorage) error
	CheckToppingsExists(context.Context, *CheckToppingsExistsReqStorage) error
	CreateOrder(context.Context, *CreateOrderReqStorage) (*CreateOrderResStorage, error)
	CancelOrder(context.Context, *CancelOrderReqStorage) (*CancelOrderResStorage, error)
	EmployeeCompleteOrder(context.Context, *EmployeeCompleteOrderReqStorage) (*EmployeeCompleteOrderResStorage, error)
}
