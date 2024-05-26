package rest

import (
	"coffeeshop-api/internal/service/users/model"
	"coffeeshop-api/pkg/errors"

	"github.com/labstack/echo/v4"
	easyjson "github.com/mailru/easyjson"
)

// easyjson:json
type SignupReq struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Phone    uint64 `json:"phone"`
	Password string `json:"password"`
}

func newSignupReq(c echo.Context) (SignupReq, error) {
	var r SignupReq

	if err := easyjson.UnmarshalFromReader(c.Request().Body, &r); err != nil {
		return SignupReq{}, errors.Wrapf(errors.InvalidRequestFormat, "parsing, %v", err)
	}

	return r, nil
}

// easyjson:json
type SigninReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func newSigninReq(c echo.Context) (SigninReq, error) {
	var r SigninReq

	if err := easyjson.UnmarshalFromReader(c.Request().Body, &r); err != nil {
		return SigninReq{}, errors.Wrapf(errors.InvalidRequestFormat, "parsing, %v", err)
	}

	return r, nil
}

// easyjson:json
type RefreshReq struct {
	RefreshToken string
}

func newRefreshReq(c echo.Context) (RefreshReq, error) {
	cookie, err := c.Cookie("refresh")
	if err != nil {
		return RefreshReq{}, errors.Wrapf(errors.MissingToken, "missing refresh token in cookie, %v", err)
	}

	return RefreshReq{RefreshToken: cookie.Value}, nil
}

// easyjson:json
type SignoutAllReq struct {
	UserID uint64 `json:"-"`
}

func newSignoutAllReq(c echo.Context) SignoutAllReq {
	return SignoutAllReq{
		UserID: c.Get("claims").(*model.JWTClaims).UserID,
	}
}

// easyjson:json
type GetMeReq struct {
	UserID uint64 `json:"-"`
}

func newGetMeReq(c echo.Context) GetMeReq {
	return GetMeReq{
		UserID: c.Get("claims").(*model.JWTClaims).UserID,
	}
}

// easyjson:json
type ChangePasswordReq struct {
	UserID      uint64 `json:"-"`
	NewPassword string `json:"new_password"`
}

func newChangePasswordReq(c echo.Context) (ChangePasswordReq, error) {
	var r ChangePasswordReq

	if err := easyjson.UnmarshalFromReader(c.Request().Body, &r); err != nil {
		return ChangePasswordReq{}, errors.Wrapf(errors.InvalidRequestFormat, "parsing, %v", err)
	}

	r.UserID = c.Get("claims").(*model.JWTClaims).UserID

	return r, nil
}

// easyjson:json
type DeleteMeReq struct {
	UserID uint64 `json:"-"`
}

func newDeleteReq(c echo.Context) DeleteMeReq {
	return DeleteMeReq{
		UserID: c.Get("claims").(*model.JWTClaims).UserID,
	}
}
