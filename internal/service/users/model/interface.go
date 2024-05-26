package model

import "github.com/sirupsen/logrus"

// Use cases
type IUsecase interface {
	Signup(*SignupReq) (*SignupRes, error)
	Signin(*SigninReq) (*SigninRes, error)
	Refresh(*RefreshReq) (*RefreshRes, error)
	SignoutAll(*SignoutAllReq) (*SignoutAllRes, error)
	GetMe(*GetMeReq) (*GetMeRes, error)
	ChangePassword(*ChangePasswordReq) error
	DeleteMe(*DeleteMeReq) (*DeleteMeRes, error)
}

// Logger
type ILogger interface {
	logrus.FieldLogger
}

// Storage
type IStorage interface {
	Create(*StorageCreateReq) (*StorageCreateRes, error)
	CheckCredentials(*StorageCheckCredentialsReq) (*StorageCheckCredentialsRes, error)
	IsDeleted(*StorageIsDeletedReq) (*StorageIsDeletedRes, error)
	GetMe(*StorageGetMeReq) (*StorageGetMeRes, error)
	ChangePassword(*StorageChangePasswordReq) error
	DeleteMe(*StorageDeleteMeReq) (*StorageDeleteMeRes, error)
}

// Cache service
type ICache interface {
	SaveUserRefreshToken(userID uint64, token *Token) error
	RevokeUserRefreshToken(userID uint64, token string) error
	RevokeAllUserRefreshTokens(userID uint64) (tokens []string, err error)
}

// Token service
type IToken interface {
	Parse(token string) (claims *JWTClaims, err error)
	CreatePair(claims *JWTClaims) (accessToken, refreshToken *Token, err error)
}
