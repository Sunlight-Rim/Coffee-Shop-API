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

func (h *handler) getCoffeeInfo(c echo.Context) (err error) {
	var (
		req *model.DeliveryGetCoffeeInfoReq
		res *model.DeliveryGetCoffeeInfoRes
	)

	// Send response
	defer func() { tools.SendResponse(c, res, err) }()

	// Parse request
	if req, err = getCoffeeReq(c); err != nil {
		return errors.Wrap(err, "request")
	}

	// Call usecase
	ucRes, err := h.uc.GetCoffeeInfo(&model.UsecaseGetCoffeeInfoReq{
		CoffeeID: req.CoffeeID,
	})
	if err != nil {
		return errors.Wrap(err, "get coffee")
	}

	res = &model.DeliveryGetCoffeeInfoRes{
		Coffee: ucRes.Coffee,
	}

	return
}

func (h *handler) listCoffee(c echo.Context) (err error) {
	var (
		req *model.DeliveryListCoffeeReq
		res *model.DeliveryListCoffeeRes
	)

	// Send response
	defer func() { tools.SendResponse(c, res, err) }()

	// Parse request
	if req, err = listCoffeeReq(c); err != nil {
		return errors.Wrap(err, "request")
	}

	// Call usecase
	ucRes, err := h.uc.ListCoffee(&model.UsecaseListCoffeeReq{
		Offset: req.Offset,
	})
	if err != nil {
		return errors.Wrap(err, "list coffee")
	}

	res = &model.DeliveryListCoffeeRes{
		CoffeeList: ucRes.CoffeeList,
	}

	return
}
