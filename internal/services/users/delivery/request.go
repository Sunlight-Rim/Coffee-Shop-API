package delivery

import (
	"coffeeshop-api/internal/services/users/model"
	"coffeeshop-api/pkg/errors"

	"github.com/labstack/echo/v4"
	easyjson "github.com/mailru/easyjson"
)

func signupReq(c echo.Context) (*model.SignupReqDelivery, error) {
	var r model.SignupReqDelivery

	if err := easyjson.UnmarshalFromReader(c.Request().Body, &r); err != nil {
		return nil, errors.Wrapf(errors.InvalidRequestFormat, "parsing, %v", err)
	}

	return &r, nil
}

func signinReq(c echo.Context) (*model.SigninReqDelivery, error) {
	var r model.SigninReqDelivery

	if err := easyjson.UnmarshalFromReader(c.Request().Body, &r); err != nil {
		return nil, errors.Wrapf(errors.InvalidRequestFormat, "parsing, %v", err)
	}

	return &r, nil
}

func refreshReq(c echo.Context) (*model.RefreshReqDelivery, error) {
	cookie, err := c.Cookie("refresh")
	if err != nil {
		return nil, errors.Wrapf(errors.MissingToken, "missing refresh token in cookie, %v", err)
	}

	return &model.RefreshReqDelivery{RefreshToken: cookie.Value}, nil
}

func signoutAllReq(c echo.Context) *model.SignoutAllReqDelivery {
	return &model.SignoutAllReqDelivery{
		UserID: c.Get("claims").(*model.Claims).UserID,
	}
}

func getMeReq(c echo.Context) *model.GetMeReqDelivery {
	return &model.GetMeReqDelivery{
		UserID: c.Get("claims").(*model.Claims).UserID,
	}
}

func changePasswordReq(c echo.Context) (*model.ChangePasswordReqDelivery, error) {
	var r model.ChangePasswordReqDelivery

	if err := easyjson.UnmarshalFromReader(c.Request().Body, &r); err != nil {
		return nil, errors.Wrapf(errors.InvalidRequestFormat, "parsing, %v", err)
	}

	r.UserID = c.Get("claims").(*model.Claims).UserID

	return &r, nil
}

func deleteMeReq(c echo.Context) *model.DeleteMeReqDelivery {
	return &model.DeleteMeReqDelivery{
		UserID: c.Get("claims").(*model.Claims).UserID,
	}
}
