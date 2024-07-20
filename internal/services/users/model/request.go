package model

import (
	"regexp"
	"unicode/utf8"

	"coffeeshop-api/pkg/errors"
)

var (
	usernameRegex = regexp.MustCompile(`^[A-z0-9]{4,40}$`)
	emailRegex    = regexp.MustCompile(`^[\w-\.]+@([\w-]+\.)+[\w-]{2,4}$`)
	passwordRegex = regexp.MustCompile(`[.,\(\);:\\\/\[\]\{\}@$!%*#?&=]`)
)

// Signup

// easyjson:json
type SignupReqDelivery struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Phone    uint64 `json:"phone"`
	Password string `json:"password"`
}

type SignupReqUsecase struct {
	Username string
	Email    string
	Phone    uint64
	Password string
}

type CreateUserReqStorage struct {
	Username     string
	Email        string
	Phone        uint64
	PasswordHash string
}

func (req *SignupReqUsecase) Validate() error {
	if !usernameRegex.MatchString(req.Username) {
		return errors.Wrap(errors.InvalidRequestContent, "username")
	}

	if !emailRegex.MatchString(req.Email) {
		return errors.Wrap(errors.InvalidRequestContent, "email")
	}

	if !passwordRegex.MatchString(req.Password) || utf8.RuneCountInString(req.Password) < 6 {
		return errors.Wrap(errors.InvalidRequestContent, "password")
	}

	return nil
}

// Signin

// easyjson:json
type SigninReqDelivery struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SigninReqUsecase struct {
	Email    string
	Password string
}

type CheckCredentialsReqStorage struct {
	Email        string
	PasswordHash string
}

func (req *SigninReqUsecase) Validate() error {
	if !emailRegex.MatchString(req.Email) {
		return errors.Wrap(errors.InvalidRequestContent, "email")
	}

	if !passwordRegex.MatchString(req.Password) || utf8.RuneCountInString(req.Password) < 6 {
		return errors.Wrap(errors.InvalidRequestContent, "password")
	}

	return nil
}

// Refresh

// easyjson:json
type RefreshReqDelivery struct {
	RefreshToken string
}

type RefreshReqUsecase struct {
	RefreshToken string
}

// SignoutAll

// easyjson:json
type SignoutAllReqDelivery struct {
	UserID uint64 `json:"-"`
}

type SignoutAllReqUsecase struct {
	UserID uint64
}

// GetMe

// easyjson:json
type GetMeReqDelivery struct {
	UserID uint64 `json:"-"`
}

type GetMeReqUsecase struct {
	UserID uint64
}

type GetMeReqStorage struct {
	UserID uint64
}

// ChangePassword

// easyjson:json
type ChangePasswordReqDelivery struct {
	UserID      uint64 `json:"-"`
	OldPassword string `json:"old_password"`
	NewPassword string `json:"new_password"`
}

type ChangePasswordReqUsecase struct {
	UserID      uint64
	OldPassword string
	NewPassword string
}

type ChangePasswordReqStorage struct {
	UserID          uint64
	OldPasswordHash string
	NewPasswordHash string
}

func (req *ChangePasswordReqUsecase) Validate() error {
	if !passwordRegex.MatchString(req.OldPassword) || utf8.RuneCountInString(req.OldPassword) < 6 {
		return errors.Wrap(errors.InvalidRequestContent, "old password")
	}

	if !passwordRegex.MatchString(req.NewPassword) || utf8.RuneCountInString(req.NewPassword) < 6 {
		return errors.Wrap(errors.InvalidRequestContent, "new password")
	}

	return nil
}

// DeleteMe

// easyjson:json
type DeleteMeReqDelivery struct {
	UserID uint64 `json:"-"`
}

type DeleteMeReqUsecase struct {
	UserID uint64
}

type DeleteMeReqStorage struct {
	UserID uint64
}
