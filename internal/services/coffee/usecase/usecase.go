package usecase

import (
	"coffeeshop-api/internal/services/coffee/model"
	"coffeeshop-api/pkg/errors"
	"context"
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
func (uc *usecase) ListCoffees(ctx context.Context, req *model.ListCoffeesReqUsecase) (*model.ListCoffeesResUsecase, error) {
	coffees, err := uc.storage.ListCoffees(ctx, &model.ListCoffeesReqStorage{
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
func (uc *usecase) GetCoffeeInfo(ctx context.Context, req *model.GetCoffeeInfoReqUsecase) (*model.GetCoffeeInfoResUsecase, error) {
	coffee, err := uc.storage.GetCoffeeInfo(ctx, &model.GetCoffeeInfoReqStorage{
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
func (uc *usecase) ListToppings(ctx context.Context, req *model.ListToppingsReqUsecase) (*model.ListToppingsResUsecase, error) {
	toppings, err := uc.storage.ListToppings(ctx, &model.ListToppingsReqStorage{
		Offset: req.Offset,
	})
	if err != nil {
		return nil, errors.Wrap(err, "list toppings")
	}

	return &model.ListToppingsResUsecase{
		ToppingsList: toppings.ToppingsList,
	}, nil
}
