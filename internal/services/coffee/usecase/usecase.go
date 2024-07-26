package usecase

import (
	"coffeeshop-api/internal/services/coffee/model"
	"coffeeshop-api/pkg/errors"
)

type usecase struct {
	storage model.IStorage
}

func New(storage model.IStorage) *usecase {
	return &usecase{
		storage: storage,
	}
}

// ListCoffees returns list of coffee.
func (uc *usecase) ListCoffees(req *model.ListCoffeesReqUsecase) (*model.ListCoffeesResUsecase, error) {
	coffees, err := uc.storage.ListCoffees(&model.ListCoffeesReqStorage{
		Offset: req.Offset,
	})
	if err != nil {
		return nil, errors.Wrap(err, "list coffees")
	}

	return &model.ListCoffeesResUsecase{
		CoffeeList: coffees.CoffeeList,
	}, nil
}

// GetCoffeeInfo returns coffee description.
func (uc *usecase) GetCoffeeInfo(req *model.GetCoffeeInfoReqUsecase) (*model.GetCoffeeInfoResUsecase, error) {
	coffee, err := uc.storage.GetCoffeeInfo(&model.GetCoffeeInfoReqStorage{
		CoffeeID: req.CoffeeID,
	})
	if err != nil {
		return nil, errors.Wrap(err, "get coffee")
	}

	return &model.GetCoffeeInfoResUsecase{
		Coffee: coffee.Coffee,
	}, nil
}

// ListToppings returns list of toppings.
func (uc *usecase) ListToppings(req *model.ListToppingsReqUsecase) (*model.ListToppingsResUsecase, error) {
	toppings, err := uc.storage.ListToppings(&model.ListToppingsReqStorage{
		Offset: req.Offset,
	})
	if err != nil {
		return nil, errors.Wrap(err, "list toppings")
	}

	return &model.ListToppingsResUsecase{
		ToppingsList: toppings.ToppingsList,
	}, nil
}
