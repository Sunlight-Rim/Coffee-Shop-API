package model

import (
	"coffeeshop-api/pkg/claims"
)

// Use cases
type IUsecase interface {
	Signup(*SignupReqUsecase) (*SignupResUsecase, error)
	Signin(*SigninReqUsecase) (*SigninResUsecase, error)
	Refresh(*RefreshReqUsecase) (*RefreshResUsecase, error)
	SignoutAll(*SignoutAllReqUsecase) (*SignoutAllResUsecase, error)
	GetMe(*GetMeReqUsecase) (*GetMeResUsecase, error)
	ChangePassword(*ChangePasswordReqUsecase) error
	DeleteMe(*DeleteMeReqUsecase) (*DeleteMeResUsecase, error)
}

// Storage
type IStorage interface {
	CreateUser(*CreateUserReqStorage) (*CreateUserResStorage, error)
	CheckCredentials(*CheckCredentialsReqStorage) (*CheckCredentialsResStorage, error)
	GetMe(*GetMeReqStorage) (*GetMeResStorage, error)
	ChangePassword(*ChangePasswordReqStorage) error
	DeleteMe(*DeleteMeReqStorage) (*DeleteMeResStorage, error)
}

// Cache service
type ICache interface {
	SaveUserRefreshToken(userID uint64, token Token) error
	RevokeUserRefreshToken(userID uint64, token string) error
	RevokeAllUserRefreshTokens(userID uint64) (tokens []string, err error)
}

// Token service
type IToken interface {
	Parse(token string) (claims *claims.Claims, err error)
	CreatePair(claims *claims.Claims) (tokensPair *TokensPair, err error)
}
