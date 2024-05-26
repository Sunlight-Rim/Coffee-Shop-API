package rest

import (
	"coffeeshop-api/pkg/errors"

	"github.com/labstack/echo/v4"
)

// easyjson:json
type GetCoffeeReq struct {
	// in: path
	CoffeeID uint64 `json:"id" param:"id"`
}

func newGetCoffeeReq(c echo.Context) (GetCoffeeReq, error) {
	var r GetCoffeeReq

	if err := c.Bind(&r); err != nil {
		return GetCoffeeReq{}, errors.Wrapf(errors.InvalidRequestFormat, "parsing, %v", err)
	}

	return r, nil
}

// easyjson:json
type ListCoffeeReq struct {
	// in: query
	Offset uint64 `json:"offset" query:"offset"`
}

func newListCoffeeReq(c echo.Context) (ListCoffeeReq, error) {
	var r ListCoffeeReq

	if err := c.Bind(&r); err != nil {
		return ListCoffeeReq{}, errors.Wrapf(errors.InvalidRequestFormat, "parsing, %v", err)
	}

	return r, nil
}
