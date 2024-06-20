package model

import "github.com/sirupsen/logrus"

// Use cases
type IUsecase interface {
	GetCoffeeInfo(*GetCoffeeInfoReqUsecase) (*GetCoffeeInfoResUsecase, error)
	ListCoffee(*ListCoffeeReqUsecase) (*ListCoffeeResUsecase, error)
}

// Logger
type ILogger interface {
	logrus.FieldLogger
}

// Storage
type IStorage interface {
	GetCoffeeInfo(*GetCoffeeInfoReqStorage) (*GetCoffeeInfoResStorage, error)
	ListCoffee(*ListCoffeeReqStorage) (*ListCoffeeResStorage, error)
}
