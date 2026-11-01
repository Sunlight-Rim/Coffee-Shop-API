package delivery

import (
	"coffeeshop-api/internal/services/coffee/model"
	"coffeeshop-api/pkg/errors"
	"coffeeshop-api/pkg/tools"
	"coffeeshop-api/pkg/tracer"

	"github.com/labstack/echo/v4"
)

type handler struct {
	uc model.IUsecase
}

func New(uc model.IUsecase) *handler {
	return &handler{uc: uc}
}

func (h *handler) listCoffees(c echo.Context) (err error) {
	ctx, span := tracer.Start(c.Request().Context(), "listCoffees")
	defer span.End()

	var (
		req *model.ListCoffeesReqDelivery
		res *model.ListCoffeesResDelivery
	)

	// Send response
	defer func() { tools.SendResponse(c, res, err) }()

	// Parse request
	if req, err = listCoffeesReq(c); err != nil {
		return errors.Wrap(err, "request")
	}

	// Call usecase
	coffees, err := h.uc.ListCoffees(ctx, &model.ListCoffeesReqUsecase{
		Offset: req.Offset,
	})
	if err != nil {
		return errors.Wrap(err, "list coffees")
	}

	res = &model.ListCoffeesResDelivery{
		CoffeeList: coffees.CoffeeList,
	}

	return
}

func (h *handler) getCoffeeInfo(c echo.Context) (err error) {
	ctx, span := tracer.Start(c.Request().Context(), "getCoffeeInfo")
	defer span.End()

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
	coffee, err := h.uc.GetCoffeeInfo(ctx, &model.GetCoffeeInfoReqUsecase{
		CoffeeID: req.CoffeeID,
	})
	if err != nil {
		return errors.Wrap(err, "get coffee")
	}

	res = &model.GetCoffeeInfoResDelivery{
		Coffee: coffee.Coffee,
	}

	return
}

func (h *handler) listToppings(c echo.Context) (err error) {
	ctx, span := tracer.Start(c.Request().Context(), "listToppings")
	defer span.End()

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
	toppings, err := h.uc.ListToppings(ctx, &model.ListToppingsReqUsecase{
		Offset: req.Offset,
	})
	if err != nil {
		return errors.Wrap(err, "list toppings")
	}

	res = &model.ListToppingsResDelivery{
		ToppingsList: toppings.ToppingsList,
	}

	return
}
