package model

import "github.com/sirupsen/logrus"

// Use cases
type IUsecase interface {
	GetCoffee(*GetCoffeeReq) (*GetCoffeeRes, error)
	ListCoffee(*ListCoffeeReq) (*ListCoffeeRes, error)
}

// Logger
type ILogger interface {
	logrus.FieldLogger
}

// Storage
type IStorage interface {
	GetCoffee(*StorageGetCoffeeReq) (*StorageGetCoffeeRes, error)
	ListCoffee(*StorageListCoffeeReq) (*StorageListCoffeeRes, error)
}
