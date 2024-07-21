package delivery

import (
	"coffeeshop-api/internal/services/orders/model"
	"coffeeshop-api/pkg/claims"
	"coffeeshop-api/pkg/errors"

	"github.com/labstack/echo/v4"
	"github.com/mailru/easyjson"
)

func ordersStatusesReq(c echo.Context) (*model.OrdersStatusesReqDelivery, error) {
	var r model.OrdersStatusesReqDelivery

	r.UserID = c.Get("claims").(*claims.Claims).UserID

	return &r, nil
}

func listOrdersReq(c echo.Context) (*model.ListOrdersReqDelivery, error) {
	var r model.ListOrdersReqDelivery

	if err := c.Bind(&r); err != nil {
		return nil, errors.Wrapf(errors.InvalidRequestFormat, "parsing, %v", err)
	}

	r.UserID = c.Get("claims").(*claims.Claims).UserID

	return &r, nil
}

func getOrderInfoReq(c echo.Context) (*model.GetOrderInfoReqDelivery, error) {
	var r model.GetOrderInfoReqDelivery

	if err := c.Bind(&r); err != nil {
		return nil, errors.Wrapf(errors.InvalidRequestFormat, "parsing, %v", err)
	}

	r.UserID = c.Get("claims").(*claims.Claims).UserID

	return &r, nil
}

func createOrderReq(c echo.Context) (*model.CreateOrderReqDelivery, error) {
	var r model.CreateOrderReqDelivery

	if err := easyjson.UnmarshalFromReader(c.Request().Body, &r); err != nil {
		return nil, errors.Wrapf(errors.InvalidRequestFormat, "parsing, %v", err)
	}

	r.UserID = c.Get("claims").(*claims.Claims).UserID

	return &r, nil
}

func cancelOrderReq(c echo.Context) (*model.CancelOrderReqDelivery, error) {
	var r model.CancelOrderReqDelivery

	if err := c.Bind(&r); err != nil {
		return nil, errors.Wrapf(errors.InvalidRequestFormat, "parsing, %v", err)
	}

	r.UserID = c.Get("claims").(*claims.Claims).UserID

	return &r, nil
}

func employeeCompleteOrderReq(c echo.Context) (*model.EmployeeCompleteOrderReqDelivery, error) {
	var r model.EmployeeCompleteOrderReqDelivery

	if err := c.Bind(&r); err != nil {
		return nil, errors.Wrapf(errors.InvalidRequestFormat, "parsing, %v", err)
	}

	return &r, nil
}
