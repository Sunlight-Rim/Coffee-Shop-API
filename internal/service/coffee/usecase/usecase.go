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

// GetCoffeeInfo returns coffee description.
func (uc *usecase) GetCoffeeInfo(req *model.UsecaseGetCoffeeInfoReq) (*model.UsecaseGetCoffeeInfoRes, error) {
	resInfo, err := uc.storage.GetCoffeeInfo(&model.StorageGetCoffeeInfoReq{
		CoffeeID: req.CoffeeID,
	})
	if err != nil {
		return nil, errors.Wrap(err, "get coffee")
	}

	return &model.UsecaseGetCoffeeInfoRes{
		Coffee: resInfo.Coffee,
	}, nil
}

// ListCoffee returns list of coffee
func (uc *usecase) ListCoffee(req *model.UsecaseListCoffeeReq) (*model.UsecaseListCoffeeRes, error) {
	resList, err := uc.storage.ListCoffee(&model.StorageListCoffeeReq{
		Offset: req.Offset,
	})
	if err != nil {
		return nil, errors.Wrap(err, "list coffee")
	}

	return &model.UsecaseListCoffeeRes{
		CoffeeList: resList.CoffeeList,
	}, nil
}
