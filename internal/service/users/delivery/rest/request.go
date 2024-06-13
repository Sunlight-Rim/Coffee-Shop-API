package rest

import (
	"coffeeshop-api/internal/service/users/model"
	"coffeeshop-api/pkg/errors"

	"github.com/labstack/echo/v4"
	easyjson "github.com/mailru/easyjson"
)

func signupReq(c echo.Context) (*model.DeliverySignupReq, error) {
	var r model.DeliverySignupReq

	if err := easyjson.UnmarshalFromReader(c.Request().Body, &r); err != nil {
		return nil, errors.Wrapf(errors.InvalidRequestFormat, "parsing, %v", err)
	}

	return &r, nil
}

func signinReq(c echo.Context) (*model.DeliverySigninReq, error) {
	var r model.DeliverySigninReq

	if err := easyjson.UnmarshalFromReader(c.Request().Body, &r); err != nil {
		return nil, errors.Wrapf(errors.InvalidRequestFormat, "parsing, %v", err)
	}

	return &r, nil
}

func refreshReq(c echo.Context) (*model.DeliveryRefreshReq, error) {
	cookie, err := c.Cookie("refresh")
	if err != nil {
		return nil, errors.Wrapf(errors.MissingToken, "missing refresh token in cookie, %v", err)
	}

	return &model.DeliveryRefreshReq{RefreshToken: cookie.Value}, nil
}

func signoutAllReq(c echo.Context) *model.DeliverySignoutAllReq {
	return &model.DeliverySignoutAllReq{
		UserID: c.Get("claims").(*model.JWTClaims).UserID,
	}
}

func getMeReq(c echo.Context) *model.DeliveryGetMeReq {
	return &model.DeliveryGetMeReq{
		UserID: c.Get("claims").(*model.JWTClaims).UserID,
	}
}

func changePasswordReq(c echo.Context) (*model.DeliveryChangePasswordReq, error) {
	var r model.DeliveryChangePasswordReq

	if err := easyjson.UnmarshalFromReader(c.Request().Body, &r); err != nil {
		return nil, errors.Wrapf(errors.InvalidRequestFormat, "parsing, %v", err)
	}

	r.UserID = c.Get("claims").(*model.JWTClaims).UserID

	return &r, nil
}

func deleteMeReq(c echo.Context) *model.DeliveryDeleteMeReq {
	return &model.DeliveryDeleteMeReq{
		UserID: c.Get("claims").(*model.JWTClaims).UserID,
	}
}
