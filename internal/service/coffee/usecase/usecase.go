package usecase

import (
	"coffeeshop-api/internal/service/coffee/model"
	"coffeeshop-api/pkg/errors"
)

type usecase struct {
	logger  model.ILogger
	storage model.IStorage
}

func New(logger model.ILogger, storage model.IStorage) *usecase {
	return &usecase{
		logger:  logger,
		storage: storage,
	}
}

func (uc *usecase) GetCoffee(req *model.GetCoffeeReq) (*model.GetCoffeeRes, error) {
	var res model.GetCoffeeRes

	// Get coffee
	resStore, err := uc.storage.GetCoffee(&model.StorageGetCoffeeReq{
		CoffeeID: req.CoffeeID,
	})
	if err != nil {
		return nil, errors.Wrap(err, "storage")
	}

	res.Coffee = resStore.Coffee

	return &res, nil
}

func (uc *usecase) ListCoffee(req *model.ListCoffeeReq) (*model.ListCoffeeRes, error) {
	var res model.ListCoffeeRes

	// Get coffee list
	resStore, err := uc.storage.ListCoffee(&model.StorageListCoffeeReq{
		Offset: req.Offset,
	})
	if err != nil {
		return nil, errors.Wrap(err, "storage")
	}

	res.CoffeeList = resStore.CoffeeList

	return &res, nil
}
