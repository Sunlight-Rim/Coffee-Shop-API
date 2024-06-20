package model

import "github.com/sirupsen/logrus"

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

// Logger
type ILogger interface {
	logrus.FieldLogger
}

// Storage
type IStorage interface {
	Create(*CreateReqStorage) (*CreateResStorage, error)
	CheckCredentials(*CheckCredentialsReqStorage) (*CheckCredentialsResStorage, error)
	IsDeleted(*IsDeletedReqStorage) (*IsDeletedResStorage, error)
	GetMe(*GetMeReqStorage) (*GetMeResStorage, error)
	ChangePassword(*ChangePasswordReqStorage) error
	DeleteMe(*DeleteMeReqStorage) (*DeleteMeResStorage, error)
}

// Cache service
type ICache interface {
	SaveUserRefreshToken(userID uint64, token *Token) error
	RevokeUserRefreshToken(userID uint64, token string) error
	RevokeAllUserRefreshTokens(userID uint64) (tokens []string, err error)
}

// Token service
type IToken interface {
	Parse(token string) (claims *Claims, err error)
	CreatePair(claims *Claims) (accessToken, refreshToken *Token, err error)
}
