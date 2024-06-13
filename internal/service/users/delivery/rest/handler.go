package rest

import (
	"net/http"
	"time"

	"coffeeshop-api/internal/service/users/model"
	"coffeeshop-api/pkg/errors"
	"coffeeshop-api/pkg/tools"

	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
)

type handler struct {
	uc model.IUsecase
}

func New(uc model.IUsecase) *handler {
	return &handler{uc: uc}
}

func (h *handler) signup(c echo.Context) (err error) {
	var (
		req *model.DeliverySignupReq
		res *model.DeliverySignupRes
	)

	// Send response
	defer func() { tools.SendResponse(c, res, err) }()

	// Parse request
	if req, err = signupReq(c); err != nil {
		return errors.Wrap(err, "request")
	}

	// Call usecase
	ucRes, err := h.uc.Signup(&model.UsecaseSignupReq{
		Username: req.Username,
		Email:    req.Email,
		Phone:    req.Phone,
		Password: req.Password,
	})
	if err != nil {
		return errors.Wrap(err, "signup")
	}

	res = &model.DeliverySignupRes{
		UserID: ucRes.UserID,
	}

	return
}

func (h *handler) signin(c echo.Context) (err error) {
	var req *model.DeliverySigninReq

	// Send response
	defer func() { tools.SendResponse(c, nil, err) }()

	// Parse request
	if req, err = signinReq(c); err != nil {
		return errors.Wrap(err, "request")
	}

	// Call usecase
	ucRes, err := h.uc.Signin(&model.UsecaseSigninReq{
		Email:    req.Email,
		Password: req.Password,
	})
	if err != nil {
		return errors.Wrap(err, "signin")
	}

	setCookieTokens(c, ucRes.AccessToken, ucRes.RefreshToken)

	return
}

func (h *handler) refresh(c echo.Context) (err error) {
	var req *model.DeliveryRefreshReq

	// Send response
	defer func() { tools.SendResponse(c, nil, err) }()

	// Parse request
	if req, err = refreshReq(c); err != nil {
		return errors.Wrap(err, "request")
	}

	// Call usecase
	ucRes, err := h.uc.Refresh(&model.UsecaseRefreshReq{
		RefreshToken: req.RefreshToken,
	})
	if err != nil {
		return errors.Wrap(err, "refresh")
	}

	setCookieTokens(c, ucRes.AccessToken, ucRes.RefreshToken)

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
		req *model.DeliverySignoutAllReq
		res *model.DeliverySignoutAllRes
	)

	// Send response
	defer func() { tools.SendResponse(c, res, err) }()

	// Parse request
	req = signoutAllReq(c)

	// Call usecase
	ucRes, err := h.uc.SignoutAll(&model.UsecaseSignoutAllReq{
		UserID: req.UserID,
	})
	if err != nil {
		return errors.Wrap(err, "signout all")
	}

	res = &model.DeliverySignoutAllRes{
		RefreshTokens: ucRes.RefreshTokens,
	}

	removeCookieTokens(c)

	return
}

func (h *handler) getMe(c echo.Context) (err error) {
	var (
		req *model.DeliveryGetMeReq
		res *model.DeliveryGetMeRes
	)

	// Send response
	defer func() { tools.SendResponse(c, res, err) }()

	// Parse request
	req = getMeReq(c)

	// Call usecase
	ucRes, err := h.uc.GetMe(&model.UsecaseGetMeReq{
		UserID: req.UserID,
	})
	if err != nil {
		return errors.Wrap(err, "get")
	}

	res = &model.DeliveryGetMeRes{
		User: ucRes.User,
	}

	return
}

func (h *handler) changePassword(c echo.Context) (err error) {
	var req *model.DeliveryChangePasswordReq

	// Send response
	defer func() { tools.SendResponse(c, nil, err) }()

	// Parse request
	if req, err = changePasswordReq(c); err != nil {
		return errors.Wrap(err, "request")
	}

	// Call usecase
	if err := h.uc.ChangePassword(&model.UsecaseChangePasswordReq{
		UserID:      req.UserID,
		NewPassword: req.NewPassword,
	}); err != nil {
		return errors.Wrap(err, "update password")
	}

	return
}

func (h *handler) deleteMe(c echo.Context) (err error) {
	var (
		req *model.DeliveryDeleteMeReq
		res *model.DeliveryDeleteMeRes
	)

	// Send response
	defer func() { tools.SendResponse(c, res, err) }()

	// Parse request
	req = deleteMeReq(c)

	// Call usecase
	ucRes, err := h.uc.DeleteMe(&model.UsecaseDeleteMeReq{
		UserID: req.UserID,
	})
	if err != nil {
		return errors.Wrap(err, "delete user")
	}

	res.User = ucRes.User

	return
}

func setCookieTokens(c echo.Context, accessToken, refreshToken *model.Token) {
	c.SetCookie(&http.Cookie{
		Name:     "access",
		Value:    accessToken.String,
		Expires:  accessToken.Exp,
		Domain:   viper.GetString("cookie.domain"),
		Secure:   viper.GetBool("cookie.secure"),
		HttpOnly: viper.GetBool("cookie.http_only"),
	})

	c.SetCookie(&http.Cookie{
		Name:     "refresh",
		Path:     "/api/auth/refresh",
		Value:    refreshToken.String,
		Expires:  refreshToken.Exp,
		Domain:   viper.GetString("cookie.domain"),
		Secure:   viper.GetBool("cookie.secure"),
		HttpOnly: viper.GetBool("cookie.http_only"),
	})
}

func removeCookieTokens(c echo.Context) {
	c.SetCookie(&http.Cookie{
		Name:     "access",
		Expires:  time.Unix(0, 0),
		Domain:   viper.GetString("cookie.domain"),
		Secure:   viper.GetBool("cookie.secure"),
		HttpOnly: viper.GetBool("cookie.http_only"),
	})

	c.SetCookie(&http.Cookie{
		Name:     "refresh",
		Path:     "/api/auth/refresh",
		Expires:  time.Unix(0, 0),
		Domain:   viper.GetString("cookie.domain"),
		Secure:   viper.GetBool("cookie.secure"),
		HttpOnly: viper.GetBool("cookie.http_only"),
	})
}
