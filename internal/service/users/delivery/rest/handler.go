package rest

import (
	"net/http"
	"time"

	"coffeeshop-api/internal/service/users/model"
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
		req SignupReq
		res SignupRes
	)

	// Send response
	defer func() { tools.SendResponse(c, res, err) }()

	// Parse request
	if req, err = newSignupReq(c); err != nil {
		return errors.Wrap(err, "request")
	}

	// Call usecase
	resUC, err := h.uc.Signup(&model.SignupReq{
		Username: req.Username,
		Email:    req.Email,
		Phone:    req.Phone,
		Password: req.Password,
	})
	if err != nil {
		return errors.Wrap(err, "signup")
	}

	res.UserID = resUC.UserID

	return
}

func (h *handler) signin(c echo.Context) (err error) {
	var req SigninReq

	// Send response
	defer func() { tools.SendResponse(c, nil, err) }()

	// Parse request
	if req, err = newSigninReq(c); err != nil {
		return errors.Wrap(err, "request")
	}

	// Call usecase
	resUC, err := h.uc.Signin(&model.SigninReq{
		Email:    req.Email,
		Password: req.Password,
	})
	if err != nil {
		return errors.Wrap(err, "signin")
	}

	setCookieTokens(c, resUC.AccessToken, resUC.RefreshToken)

	return
}

func (h *handler) refresh(c echo.Context) (err error) {
	var req RefreshReq

	// Send response
	defer func() { tools.SendResponse(c, nil, err) }()

	// Parse request
	if req, err = newRefreshReq(c); err != nil {
		return errors.Wrap(err, "request")
	}

	// Call usecase
	resUC, err := h.uc.Refresh(&model.RefreshReq{
		RefreshToken: req.RefreshToken,
	})
	if err != nil {
		return errors.Wrap(err, "refresh")
	}

	setCookieTokens(c, resUC.AccessToken, resUC.RefreshToken)

	return
}

func (h *handler) signout(c echo.Context) (err error) {
	// Send response
	defer func() { tools.SendResponse(c, nil, err) }()

	removeCookieTokens(c)

	return
}

func (h *handler) signoutAll(c echo.Context) (err error) {
	var (
		req SignoutAllReq
		res SignoutAllRes
	)

	// Send response
	defer func() { tools.SendResponse(c, res, err) }()

	// Parse request
	req = newSignoutAllReq(c)

	// Call usecase
	resUC, err := h.uc.SignoutAll(&model.SignoutAllReq{
		UserID: req.UserID,
	})
	if err != nil {
		return errors.Wrap(err, "signout all")
	}

	res.RefreshTokens = resUC.RefreshTokens

	removeCookieTokens(c)

	return
}

func (h *handler) getMe(c echo.Context) (err error) {
	var (
		req GetMeReq
		res GetMeRes
	)

	// Send response
	defer func() { tools.SendResponse(c, res, err) }()

	// Parse request
	req = newGetMeReq(c)

	// Call usecase
	resUC, err := h.uc.GetMe(&model.GetMeReq{
		UserID: req.UserID,
	})
	if err != nil {
		return errors.Wrap(err, "get")
	}

	res.User = User(*resUC.User)

	return
}

func (h *handler) updatePassword(c echo.Context) (err error) {
	var req ChangePasswordReq

	// Send response
	defer func() { tools.SendResponse(c, nil, err) }()

	// Parse request
	if req, err = newChangePasswordReq(c); err != nil {
		return errors.Wrap(err, "request")
	}

	// Call usecase
	if err := h.uc.ChangePassword(&model.ChangePasswordReq{
		UserID:      req.UserID,
		NewPassword: req.NewPassword,
	}); err != nil {
		return errors.Wrap(err, "update password")
	}

	return
}

func (h *handler) deleteMe(c echo.Context) (err error) {
	var (
		req DeleteMeReq
		res DeleteMeRes
	)

	// Send response
	defer func() { tools.SendResponse(c, res, err) }()

	// Parse request
	req = newDeleteReq(c)

	// Call usecase
	resUC, err := h.uc.DeleteMe(&model.DeleteMeReq{
		UserID: req.UserID,
	})
	if err != nil {
		return errors.Wrap(err, "delete user")
	}

	res.User = User(*resUC.User)

	return
}

func setCookieTokens(c echo.Context, accessToken, refreshToken *model.Token) {
	c.SetCookie(&http.Cookie{
		Name:     "access",
		Value:    accessToken.String,
		Expires:  accessToken.Exp,
		HttpOnly: true,
		// Secure:   true,
	})

	c.SetCookie(&http.Cookie{
		Name:     "refresh",
		Path:     "/api/auth/refresh",
		Value:    refreshToken.String,
		Expires:  refreshToken.Exp,
		HttpOnly: true,
		// Secure:   true,
	})
}

func removeCookieTokens(c echo.Context) {
	c.SetCookie(&http.Cookie{
		Name:     "access",
		Expires:  time.Unix(0, 0),
		HttpOnly: true,
		// Secure:   true,
	})

	c.SetCookie(&http.Cookie{
		Name:     "refresh",
		Path:     "/api/auth/refresh",
		Expires:  time.Unix(0, 0),
		HttpOnly: true,
		// Secure:   true,
	})
}
