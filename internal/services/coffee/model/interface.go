package model

import "context"

//go:generate mockery --all --output=mocks

// Use cases
type IUsecase interface {
	ListCoffees(context.Context, *ListCoffeesReqUsecase) (*ListCoffeesResUsecase, error)
	GetCoffeeInfo(context.Context, *GetCoffeeInfoReqUsecase) (*GetCoffeeInfoResUsecase, error)
	ListToppings(context.Context, *ListToppingsReqUsecase) (*ListToppingsResUsecase, error)
}

// Storage
type IStorage interface {
	ListCoffees(context.Context, *ListCoffeesReqStorage) (*ListCoffeesResStorage, error)
	GetCoffeeInfo(context.Context, *GetCoffeeInfoReqStorage) (*GetCoffeeInfoResStorage, error)
	ListToppings(context.Context, *ListToppingsReqStorage) (*ListToppingsResStorage, error)
}
