package rest

import (
	"coffeeshop-api/internal/service/coffee/model"
	"coffeeshop-api/pkg/errors"
	"coffeeshop-api/pkg/tools"

	"github.com/labstack/echo/v4"
)

type handler struct {
	uc model.IUsecase
}

func New(uc model.IUsecase) *handler {
	return &handler{uc: uc}
}

func (h *handler) getCoffee(c echo.Context) (err error) {
	var (
		req GetCoffeeReq
		res GetCoffeeRes
	)

	// Send response
	defer func() { tools.SendResponse(c, res, err) }()

	// Parse request
	if req, err = newGetCoffeeReq(c); err != nil {
		return errors.Wrap(err, "request")
	}

	// Call usecase
	resUC, err := h.uc.GetCoffee(&model.GetCoffeeReq{
		CoffeeID: req.CoffeeID,
	})
	if err != nil {
		return errors.Wrap(err, "get coffee")
	}

	res.Coffee = Coffee(*resUC.Coffee)

	return
}

func (h *handler) listCoffee(c echo.Context) (err error) {
	var (
		req ListCoffeeReq
		res ListCoffeeRes
	)

	// Send response
	defer func() { tools.SendResponse(c, res, err) }()

	// Parse request
	if req, err = newListCoffeeReq(c); err != nil {
		return errors.Wrap(err, "request")
	}

	// Call usecase
	resUC, err := h.uc.ListCoffee(&model.ListCoffeeReq{
		Offset: req.Offset,
	})
	if err != nil {
		return errors.Wrap(err, "list coffee")
	}

	res.CoffeeList = make([]Coffee, 0, len(resUC.CoffeeList))
	for i := range resUC.CoffeeList {
		res.CoffeeList = append(res.CoffeeList, Coffee(resUC.CoffeeList[i]))
	}

	return
}
