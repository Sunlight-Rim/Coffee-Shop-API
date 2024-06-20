package rest

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
	userInfo, err := h.uc.Signup(&model.SignupReqUsecase{
		Username: req.Username,
		Email:    req.Email,
		Phone:    req.Phone,
		Password: req.Password,
	})
	if err != nil {
		return errors.Wrap(err, "signup")
	}

	res = &model.SignupResDelivery{
		UserID: userInfo.UserID,
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
	tokens, err := h.uc.Signin(&model.SigninReqUsecase{
		Email:    req.Email,
		Password: req.Password,
	})
	if err != nil {
		return errors.Wrap(err, "signin")
	}

	setCookieTokens(c, tokens.AccessToken, tokens.RefreshToken)

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
	tokens, err := h.uc.Refresh(&model.RefreshReqUsecase{
		RefreshToken: req.RefreshToken,
	})
	if err != nil {
		return errors.Wrap(err, "refresh")
	}

	setCookieTokens(c, tokens.AccessToken, tokens.RefreshToken)

	return
}

func (h *handler) signout(c echo.Context) (err error) {
	// Send response
	defer func() { tools.SendResponse(c, nil, err) }()

	setCookieTokens(c, &model.Token{}, &model.Token{})

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
	tokens, err := h.uc.SignoutAll(&model.SignoutAllReqUsecase{
		UserID: req.UserID,
	})
	if err != nil {
		return errors.Wrap(err, "signout all")
	}

	res = &model.SignoutAllResDelivery{
		RefreshTokens: tokens.RefreshTokens,
	}

	setCookieTokens(c, &model.Token{}, &model.Token{})

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
	userInfo, err := h.uc.GetMe(&model.GetMeReqUsecase{
		UserID: req.UserID,
	})
	if err != nil {
		return errors.Wrap(err, "get")
	}

	res = &model.GetMeResDelivery{
		User: userInfo.User,
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
	if err := h.uc.ChangePassword(&model.ChangePasswordReqUsecase{
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
	userInfo, err := h.uc.DeleteMe(&model.DeleteMeReqUsecase{
		UserID: req.UserID,
	})
	if err != nil {
		return errors.Wrap(err, "delete user")
	}

	res.User = userInfo.User

	return
}

func setCookieTokens(c echo.Context, accessToken, refreshToken *model.Token) {
	tools.SetCookie(c, "access", accessToken.String, "/", accessToken.Exp)
	tools.SetCookie(c, "refresh", refreshToken.String, "/api/auth/refresh", refreshToken.Exp)
}
