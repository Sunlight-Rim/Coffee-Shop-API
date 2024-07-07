package delivery

import (
	"coffeeshop-api/internal/services/coffee/model"
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

func (h *handler) listCoffee(c echo.Context) (err error) {
	var (
		req *model.ListCoffeeReqDelivery
		res *model.ListCoffeeResDelivery
	)

	// Send response
	defer func() { tools.SendResponse(c, res, err) }()

	// Parse request
	if req, err = listCoffeeReq(c); err != nil {
		return errors.Wrap(err, "request")
	}

	// Call usecase
	list, err := h.uc.ListCoffee(&model.ListCoffeeReqUsecase{
		Offset: req.Offset,
	})
	if err != nil {
		return errors.Wrap(err, "list coffee")
	}

	res = &model.ListCoffeeResDelivery{
		CoffeeList: list.CoffeeList,
	}

	return
}

func (h *handler) getCoffeeInfo(c echo.Context) (err error) {
	var (
		req *model.GetCoffeeInfoReqDelivery
		res *model.GetCoffeeInfoResDelivery
	)

	// Send response
	defer func() { tools.SendResponse(c, res, err) }()

	// Parse request
	if req, err = getCoffeeInfoReq(c); err != nil {
		return errors.Wrap(err, "request")
	}

	// Call usecase
	info, err := h.uc.GetCoffeeInfo(&model.GetCoffeeInfoReqUsecase{
		CoffeeID: req.CoffeeID,
	})
	if err != nil {
		return errors.Wrap(err, "get coffee")
	}

	res = &model.GetCoffeeInfoResDelivery{
		Coffee: info.Coffee,
	}

	return
}

func (h *handler) listToppings(c echo.Context) (err error) {
	var (
		req *model.ListToppingsReqDelivery
		res *model.ListToppingsResDelivery
	)

	// Send response
	defer func() { tools.SendResponse(c, res, err) }()

	// Parse request
	if req, err = listToppingsReq(c); err != nil {
		return errors.Wrap(err, "request")
	}

	// Call usecase
	list, err := h.uc.ListToppings(&model.ListToppingsReqUsecase{
		Offset: req.Offset,
	})
	if err != nil {
		return errors.Wrap(err, "list toppings")
	}

	res = &model.ListToppingsResDelivery{
		ToppingsList: list.ToppingsList,
	}

	return
}
