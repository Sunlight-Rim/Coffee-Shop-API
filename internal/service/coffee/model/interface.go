package model

import "github.com/sirupsen/logrus"

// Use cases
type IUsecase interface {
	GetCoffeeInfo(*UsecaseGetCoffeeInfoReq) (*UsecaseGetCoffeeInfoRes, error)
	ListCoffee(*UsecaseListCoffeeReq) (*UsecaseListCoffeeRes, error)
}

// Logger
type ILogger interface {
	logrus.FieldLogger
}

// Storage
type IStorage interface {
	GetCoffeeInfo(*StorageGetCoffeeInfoReq) (*StorageGetCoffeeInfoRes, error)
	ListCoffee(*StorageListCoffeeReq) (*StorageListCoffeeRes, error)
}
