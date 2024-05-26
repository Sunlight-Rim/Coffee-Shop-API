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

type SignupReq struct {
	Username string
	Email    string
	Phone    uint64
	Password string
}

func (r *SignupReq) Validate() error {
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

type StorageCreateReq struct {
	Username     string
	Email        string
	Phone        uint64
	PasswordHash string
}

// Signin

type SigninReq struct {
	Email    string
	Password string
}

func (r *SigninReq) Validate() error {
	if !emailRegex.MatchString(r.Email) {
		return errors.Wrap(errors.InvalidRequestContent, "email")
	}

	if !passwordRegex.MatchString(r.Password) || utf8.RuneCountInString(r.Password) < 6 {
		return errors.Wrap(errors.InvalidRequestContent, "password")
	}

	return nil
}

type StorageCheckCredentialsReq struct {
	Email        string
	PasswordHash string
}

// IsDeleted

type StorageIsDeletedReq struct {
	UserID uint64
}

// Refresh

type RefreshReq struct {
	RefreshToken string
}

// SignoutAll

type SignoutAllReq struct {
	UserID uint64
}

// GetMe

type GetMeReq struct {
	UserID uint64
}

type StorageGetMeReq struct {
	UserID uint64
}

// ChangePassword

type ChangePasswordReq struct {
	UserID      uint64
	NewPassword string
}

func (r *ChangePasswordReq) Validate() error {
	if !passwordRegex.MatchString(r.NewPassword) || utf8.RuneCountInString(r.NewPassword) < 6 {
		return errors.Wrap(errors.InvalidRequestContent, "password")
	}

	return nil
}

type StorageChangePasswordReq struct {
	UserID          uint64
	NewPasswordHash string
}

// DeleteMe

type DeleteMeReq struct {
	UserID uint64
}

type StorageDeleteMeReq struct {
	UserID uint64
}
