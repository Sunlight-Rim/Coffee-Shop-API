package delivery

import (
	"coffeeshop-api/internal/services/users/model"
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

func (h *handler) signup(c echo.Context) (err error) {
	var (
		req *model.SignupReqDelivery
		res *model.SignupResDelivery
	)

	// Send response
	defer func() { tools.SendResponse(c, res, err) }()

	// Parse request
	if req, err = signupReq(c); err != nil {
		return errors.Wrap(err, "request")
	}

	// Call usecase
	user, err := h.uc.Signup(c.Request().Context(), &model.SignupReqUsecase{
		Username: req.Username,
		Email:    req.Email,
		Phone:    req.Phone,
		Password: req.Password,
	})
	if err != nil {
		return errors.Wrap(err, "signup")
	}

	res = &model.SignupResDelivery{
		UserID: user.UserID,
	}

	return
}

func (h *handler) signin(c echo.Context) (err error) {
	var req *model.SigninReqDelivery

	// Send response
	defer func() { tools.SendResponse(c, nil, err) }()

	// Parse request
	if req, err = signinReq(c); err != nil {
		return errors.Wrap(err, "request")
	}

	// Call usecase
	tokens, err := h.uc.Signin(c.Request().Context(), &model.SigninReqUsecase{
		Email:    req.Email,
		Password: req.Password,
	})
	if err != nil {
		return errors.Wrap(err, "signin")
	}

	setCookieTokens(c, tokens.TokensPair)

	return
}

func (h *handler) refresh(c echo.Context) (err error) {
	var req *model.RefreshReqDelivery

	// Send response
	defer func() { tools.SendResponse(c, nil, err) }()

	// Parse request
	if req, err = refreshReq(c); err != nil {
		return errors.Wrap(err, "request")
	}

	// Call usecase
	tokens, err := h.uc.Refresh(c.Request().Context(), &model.RefreshReqUsecase{
		RefreshToken: req.RefreshToken,
	})
	if err != nil {
		return errors.Wrap(err, "refresh")
	}

	setCookieTokens(c, tokens.TokensPair)

	return
}

func (h *handler) signout(c echo.Context) (err error) {
	// Send response
	defer func() { tools.SendResponse(c, nil, err) }()

	setCookieTokens(c, &model.TokensPair{})

	return
}

func (h *handler) signoutAll(c echo.Context) (err error) {
	var (
		req *model.SignoutAllReqDelivery
		res *model.SignoutAllResDelivery
	)

	// Send response
	defer func() { tools.SendResponse(c, res, err) }()

	// Parse request
	req = signoutAllReq(c)

	// Call usecase
	tokens, err := h.uc.SignoutAll(c.Request().Context(), &model.SignoutAllReqUsecase{
		UserID: req.UserID,
	})
	if err != nil {
		return errors.Wrap(err, "signout all")
	}

	res = &model.SignoutAllResDelivery{
		RefreshTokens: tokens.RefreshTokens,
	}

	setCookieTokens(c, &model.TokensPair{})

	return
}

func (h *handler) getMe(c echo.Context) (err error) {
	var (
		req *model.GetMeReqDelivery
		res *model.GetMeResDelivery
	)

	// Send response
	defer func() { tools.SendResponse(c, res, err) }()

	// Parse request
	req = getMeReq(c)

	// Call usecase
	user, err := h.uc.GetMe(c.Request().Context(), &model.GetMeReqUsecase{
		UserID: req.UserID,
	})
	if err != nil {
		return errors.Wrap(err, "get")
	}

	res = &model.GetMeResDelivery{
		User: user.User,
	}

	return
}

func (h *handler) changePassword(c echo.Context) (err error) {
	var req *model.ChangePasswordReqDelivery

	// Send response
	defer func() { tools.SendResponse(c, nil, err) }()

	// Parse request
	if req, err = changePasswordReq(c); err != nil {
		return errors.Wrap(err, "request")
	}

	// Call usecase
	if err := h.uc.ChangePassword(c.Request().Context(), &model.ChangePasswordReqUsecase{
		UserID:      req.UserID,
		OldPassword: req.OldPassword,
		NewPassword: req.NewPassword,
	}); err != nil {
		return errors.Wrap(err, "update password")
	}

	return
}

func (h *handler) deleteMe(c echo.Context) (err error) {
	var (
		req *model.DeleteMeReqDelivery
		res *model.DeleteMeResDelivery
	)

	// Send response
	defer func() { tools.SendResponse(c, res, err) }()

	// Parse request
	req = deleteMeReq(c)

	// Call usecase
	user, err := h.uc.DeleteMe(c.Request().Context(), &model.DeleteMeReqUsecase{
		UserID: req.UserID,
	})
	if err != nil {
		return errors.Wrap(err, "delete user")
	}

	res = &model.DeleteMeResDelivery{
		User: user.User,
	}

	return
}

func setCookieTokens(c echo.Context, tokensPair *model.TokensPair) {
	tools.SetCookie(
		c,
		"access",
		"/",
		tokensPair.AccessToken.String,
		tokensPair.AccessToken.Exp,
	)
	tools.SetCookie(
		c,
		"refresh",
		"/api/auth/refresh",
		tokensPair.RefreshToken.String,
		tokensPair.RefreshToken.Exp,
	)
}
