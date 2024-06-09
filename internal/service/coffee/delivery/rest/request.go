package rest

import (
	"coffeeshop-api/internal/service/coffee/model"
	"coffeeshop-api/pkg/errors"

	"github.com/labstack/echo/v4"
)

func getCoffeeReq(c echo.Context) (*model.DeliveryGetCoffeeInfoReq, error) {
	var r model.DeliveryGetCoffeeInfoReq

	if err := c.Bind(&r); err != nil {
		return nil, errors.Wrapf(errors.InvalidRequestFormat, "parsing, %v", err)
	}

	return &r, nil
}

func listCoffeeReq(c echo.Context) (*model.DeliveryListCoffeeReq, error) {
	var r model.DeliveryListCoffeeReq

	if err := c.Bind(&r); err != nil {
		return nil, errors.Wrapf(errors.InvalidRequestFormat, "parsing, %v", err)
	}

	return &r, nil
}
