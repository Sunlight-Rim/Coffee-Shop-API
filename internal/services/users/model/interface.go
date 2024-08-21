package model

import (
	"coffeeshop-api/pkg/claims"
	"context"
)

// Use cases
type IUsecase interface {
	Signup(context.Context, *SignupReqUsecase) (*SignupResUsecase, error)
	Signin(context.Context, *SigninReqUsecase) (*SigninResUsecase, error)
	Refresh(context.Context, *RefreshReqUsecase) (*RefreshResUsecase, error)
	SignoutAll(context.Context, *SignoutAllReqUsecase) (*SignoutAllResUsecase, error)
	GetMe(context.Context, *GetMeReqUsecase) (*GetMeResUsecase, error)
	ChangePassword(context.Context, *ChangePasswordReqUsecase) error
	DeleteMe(context.Context, *DeleteMeReqUsecase) (*DeleteMeResUsecase, error)
}

// Storage
type IStorage interface {
	CreateUser(context.Context, *CreateUserReqStorage) (*CreateUserResStorage, error)
	CheckCredentials(context.Context, *CheckCredentialsReqStorage) (*CheckCredentialsResStorage, error)
	GetMe(context.Context, *GetMeReqStorage) (*GetMeResStorage, error)
	ChangePassword(context.Context, *ChangePasswordReqStorage) error
	DeleteMe(context.Context, *DeleteMeReqStorage) (*DeleteMeResStorage, error)
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
