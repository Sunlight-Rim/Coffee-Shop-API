package delivery

import (
	"coffeeshop-api/internal/services/coffee/model"
	"coffeeshop-api/pkg/errors"

	"github.com/labstack/echo/v4"
)

func listCoffeesReq(c echo.Context) (*model.ListCoffeesReqDelivery, error) {
	var r model.ListCoffeesReqDelivery

	if err := c.Bind(&r); err != nil {
		return nil, errors.Wrapf(errors.InvalidRequestFormat, "parsing, %v", err)
	}

	return &r, nil
}

func getCoffeeInfoReq(c echo.Context) (*model.GetCoffeeInfoReqDelivery, error) {
	var r model.GetCoffeeInfoReqDelivery

	if err := c.Bind(&r); err != nil {
		return nil, errors.Wrapf(errors.InvalidRequestFormat, "parsing, %v", err)
	}

	return &r, nil
}

func listToppingsReq(c echo.Context) (*model.ListToppingsReqDelivery, error) {
	var r model.ListToppingsReqDelivery

	if err := c.Bind(&r); err != nil {
		return nil, errors.Wrapf(errors.InvalidRequestFormat, "parsing, %v", err)
	}

	return &r, nil
}
