package model

// Use cases
type IUsecase interface {
	ListCoffee(*ListCoffeeReqUsecase) (*ListCoffeeResUsecase, error)
	GetCoffeeInfo(*GetCoffeeInfoReqUsecase) (*GetCoffeeInfoResUsecase, error)
	ListToppings(*ListToppingsReqUsecase) (*ListToppingsResUsecase, error)
}

// Storage
type IStorage interface {
	ListCoffee(*ListCoffeeReqStorage) (*ListCoffeeResStorage, error)
	GetCoffeeInfo(*GetCoffeeInfoReqStorage) (*GetCoffeeInfoResStorage, error)
	ListToppings(*ListToppingsReqStorage) (*ListToppingsResStorage, error)
}
