package rest

import (
	"coffeeshop-api/internal/server"

	"github.com/labstack/echo/v4"
)

func (h *handler) Register(group *echo.Group) {
	auth := group.Group("/auth")
	user := group.Group("/user", server.AuthMW)

	/*
		swagger:route POST /api/auth/signup Auth SignupRequest

		Register a new user.
		Password must be longer than 6 characters and contain at least one special character.
		Username must be string with 6-40 characters.

			schemes: http
			responses:
				200: SignupResponse
				default: ErrorResponse
	*/
	auth.POST("/signup", h.signup)
	/*
		swagger:route POST /api/auth/signin Auth SigninRequest

		Sign in to user account.

			schemes: http
			responses:
				200: SigninResponse
				default: ErrorResponse
	*/
	auth.POST("/signin", h.signin)
	/*
		swagger:route POST /api/auth/refresh Auth RefreshRequestNull

		Create new tokens pair by refresh token.

			schemes: http
			responses:
				200: RefreshResponse
				default: ErrorResponse
	*/
	auth.POST("/refresh", h.refresh)
	/*
		swagger:route POST /api/auth/signout Auth SignoutRequestNull

		Remove tokens from cookies.

			schemes: http
			security:
				accessToken: []
			responses:
				200: SignoutResponse
				default: ErrorResponse
	*/
	auth.POST("/signout", h.signout, server.AuthMW)
	/*
		swagger:route POST /api/auth/signout-all Auth SignoutAllRequestNull

		Revoke all user refresh tokens (sessions) and removes tokens from cookies.

			schemes: http
			security:
				accessToken: []
			responses:
				200: SignoutAllResponse
				default: ErrorResponse
	*/
	auth.POST("/signout-all", h.signoutAll, server.AuthMW)

	/*
		swagger:route GET /api/user User GetMeNull

		Get user own account information.

			schemes: http
			security:
				accessToken: []
			responses:
				200: GetMeResponse
				default: ErrorResponse
	*/
	user.GET("", h.getMe)
	/*
		swagger:route PUT /api/user/password User ChangePasswordRequest

		Change user password.

			schemes: http
			security:
				accessToken: []
			responses:
				200: ChangePasswordResponse
				default: ErrorResponse
	*/
	user.PUT("/password", h.updatePassword)
	/*
		swagger:route DELETE /api/user User DeleteMeNull

		Delete user own account.

			schemes: http
			security:
				accessToken: []
			responses:
				200: DeleteMeResponse
				default: ErrorResponse
	*/
	user.DELETE("", h.deleteMe)
}
