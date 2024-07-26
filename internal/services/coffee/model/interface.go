package model

// Use cases
type IUsecase interface {
	ListCoffees(*ListCoffeesReqUsecase) (*ListCoffeesResUsecase, error)
	GetCoffeeInfo(*GetCoffeeInfoReqUsecase) (*GetCoffeeInfoResUsecase, error)
	ListToppings(*ListToppingsReqUsecase) (*ListToppingsResUsecase, error)
}

// Storage
type IStorage interface {
	ListCoffees(*ListCoffeesReqStorage) (*ListCoffeesResStorage, error)
	GetCoffeeInfo(*GetCoffeeInfoReqStorage) (*GetCoffeeInfoResStorage, error)
	ListToppings(*ListToppingsReqStorage) (*ListToppingsResStorage, error)
}
