package delivery

import (
	"coffeeshop-api/internal/services/orders/model"
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

func (h *handler) createOrder(c echo.Context) (err error) {
	var (
		req *model.CreateOrderReqDelivery
		res *model.CreateOrderResDelivery
	)

	// Send response
	defer func() { tools.SendResponse(c, res, err) }()

	// Parse request
	if req, err = createOrderReq(c); err != nil {
		return errors.Wrap(err, "request")
	}

	// Call usecase
	orderInfo, err := h.uc.CreateOrder(&model.CreateOrderReqUsecase{
		Items: req.Items,
	})
	if err != nil {
		return errors.Wrap(err, "create order")
	}

	res = &model.CreateOrderResDelivery{
		OrderID: orderInfo.OrderID,
	}

	return
}
