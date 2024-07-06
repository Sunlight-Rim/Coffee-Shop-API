package rest

import (
	"coffeeshop-api/internal/services/coffee/model"
	"coffeeshop-api/pkg/errors"

	"github.com/labstack/echo/v4"
)

func listCoffeeReq(c echo.Context) (*model.ListCoffeeReqDelivery, error) {
	var r model.ListCoffeeReqDelivery

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
