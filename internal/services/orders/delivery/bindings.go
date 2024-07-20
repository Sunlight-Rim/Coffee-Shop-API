package delivery

import (
	"coffeeshop-api/internal/services/orders/model"
	"coffeeshop-api/pkg/claims"
	"coffeeshop-api/pkg/errors"

	"github.com/labstack/echo/v4"
	"github.com/mailru/easyjson"
)

func createOrderReq(c echo.Context) (*model.CreateOrderReqDelivery, error) {
	var r model.CreateOrderReqDelivery

	if err := easyjson.UnmarshalFromReader(c.Request().Body, &r); err != nil {
		return nil, errors.Wrapf(errors.InvalidRequestFormat, "parsing, %v", err)
	}

	r.UserID = c.Get("claims").(*claims.Claims).UserID

	return &r, nil
}
