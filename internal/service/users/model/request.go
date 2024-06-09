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
type DeliverySignupReq struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Phone    uint64 `json:"phone"`
	Password string `json:"password"`
}

type UsecaseSignupReq struct {
	Username string
	Email    string
	Phone    uint64
	Password string
}

type StorageCreateReq struct {
	Username     string
	Email        string
	Phone        uint64
	PasswordHash string
}

func (r *UsecaseSignupReq) Validate() error {
	if !usernameRegex.MatchString(r.Username) {
		return errors.Wrap(errors.InvalidRequestContent, "username")
	}

	if !emailRegex.MatchString(r.Email) {
		return errors.Wrap(errors.InvalidRequestContent, "email")
	}

	if !passwordRegex.MatchString(r.Password) || utf8.RuneCountInString(r.Password) < 6 {
		return errors.Wrap(errors.InvalidRequestContent, "password")
	}

	return nil
}

// Signin

// easyjson:json
type DeliverySigninReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UsecaseSigninReq struct {
	Email    string
	Password string
}

type StorageCheckCredentialsReq struct {
	Email        string
	PasswordHash string
}

func (r *UsecaseSigninReq) Validate() error {
	if !emailRegex.MatchString(r.Email) {
		return errors.Wrap(errors.InvalidRequestContent, "email")
	}

	if !passwordRegex.MatchString(r.Password) || utf8.RuneCountInString(r.Password) < 6 {
		return errors.Wrap(errors.InvalidRequestContent, "password")
	}

	return nil
}

// IsDeleted

type StorageIsDeletedReq struct {
	UserID uint64
}

// Refresh

// easyjson:json
type DeliveryRefreshReq struct {
	RefreshToken string
}

type UsecaseRefreshReq struct {
	RefreshToken string
}

// SignoutAll

// easyjson:json
type DeliverySignoutAllReq struct {
	UserID uint64 `json:"-"`
}

type UsecaseSignoutAllReq struct {
	UserID uint64
}

// GetMe

// easyjson:json
type DeliveryGetMeReq struct {
	UserID uint64 `json:"-"`
}

type UsecaseGetMeReq struct {
	UserID uint64
}

type StorageGetMeReq struct {
	UserID uint64
}

// ChangePassword

// easyjson:json
type DeliveryChangePasswordReq struct {
	UserID      uint64 `json:"-"`
	NewPassword string `json:"new_password"`
}

type UsecaseChangePasswordReq struct {
	UserID      uint64
	NewPassword string
}

type StorageChangePasswordReq struct {
	UserID          uint64
	NewPasswordHash string
}

func (r *UsecaseChangePasswordReq) Validate() error {
	if !passwordRegex.MatchString(r.NewPassword) || utf8.RuneCountInString(r.NewPassword) < 6 {
		return errors.Wrap(errors.InvalidRequestContent, "password")
	}

	return nil
}

// DeleteMe

// easyjson:json
type DeliveryDeleteMeReq struct {
	UserID uint64 `json:"-"`
}

type UsecaseDeleteMeReq struct {
	UserID uint64
}

type StorageDeleteMeReq struct {
	UserID uint64
}
