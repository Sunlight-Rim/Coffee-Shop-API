package usecase

import (
	"coffeeshop-api/internal/services/users/model"
	"coffeeshop-api/pkg/claims"
	"coffeeshop-api/pkg/errors"
	"coffeeshop-api/pkg/tools"
	"context"
	"strings"
)

type usecase struct {
	storage model.IStorage
	cache   model.ICache
	token   model.IToken
}

// New usecase.
func New(
	storage model.IStorage,
	cache model.ICache,
	token model.IToken,
) *usecase {
	return &usecase{
		storage: storage,
		cache:   cache,
		token:   token,
	}
}

// Signup creates user in database.
func (uc *usecase) Signup(ctx context.Context, req *model.SignupReqUsecase) (*model.SignupResUsecase, error) {
	if err := req.Validate(); err != nil {
		return nil, errors.Wrap(err, "request validation")
	}

	user, err := uc.storage.CreateUser(ctx, &model.CreateUserReqStorage{
		Username:     req.Username,
		Phone:        req.Phone,
		Email:        strings.ToLower(req.Email),
		PasswordHash: tools.SHA256(req.Password),
	})
	if err != nil {
		return nil, errors.Wrap(err, "create user")
	}

	return &model.SignupResUsecase{
		UserID: user.UserID,
	}, nil
}

// Signin checks credentials and generates tokens.
func (uc *usecase) Signin(ctx context.Context, req *model.SigninReqUsecase) (*model.SigninResUsecase, error) {
	if err := req.Validate(); err != nil {
		return nil, errors.Wrap(err, "request validation")
	}

	// Check credentials
	user, err := uc.storage.CheckCredentials(ctx, &model.CheckCredentialsReqStorage{
		Email:        strings.ToLower(req.Email),
		PasswordHash: tools.SHA256(req.Password),
	})
	if err != nil {
		return nil, errors.Wrap(err, "check credentials")
	}

	// Create new tokens pair
	tokensPair, err := uc.token.CreatePair(&claims.Claims{
		UserID: user.UserID,
	})
	if err != nil {
		return nil, errors.Wrap(err, "create tokens pair")
	}

	// Save refresh token
	if err := uc.cache.SaveUserRefreshToken(user.UserID, tokensPair.RefreshToken); err != nil {
		return nil, errors.Wrap(err, "save refresh token")
	}

	return &model.SigninResUsecase{
		TokensPair: tokensPair,
	}, nil
}

// Refresh revokes refresh token and updates tokens.
func (uc *usecase) Refresh(ctx context.Context, req *model.RefreshReqUsecase) (*model.RefreshResUsecase, error) {
	// Parse refresh token
	claims, err := uc.token.Parse(req.RefreshToken)
	if err != nil {
		return nil, errors.Wrap(err, "parse refresh token")
	}

	// Revoke refresh token
	if err := uc.cache.RevokeUserRefreshToken(claims.UserID, req.RefreshToken); err != nil {
		return nil, errors.Wrap(err, "revoke refresh token")
	}

	// Create new tokens pair
	tokensPair, err := uc.token.CreatePair(claims)
	if err != nil {
		return nil, errors.Wrap(err, "create tokens pair")
	}

	// Save refresh token
	if err := uc.cache.SaveUserRefreshToken(claims.UserID, tokensPair.RefreshToken); err != nil {
		return nil, errors.Wrap(err, "save refresh token")
	}

	return &model.RefreshResUsecase{
		TokensPair: tokensPair,
	}, nil
}

// SignoutAll revokes all refresh tokens.
func (uc *usecase) SignoutAll(ctx context.Context, req *model.SignoutAllReqUsecase) (*model.SignoutAllResUsecase, error) {
	// Revoke all refresh tokens
	refreshTokens, err := uc.cache.RevokeAllUserRefreshTokens(req.UserID)
	if err != nil {
		return nil, errors.Wrap(err, "revoke all refresh tokens")
	}

	return &model.SignoutAllResUsecase{
		RefreshTokens: refreshTokens,
	}, nil
}

// GetMe returns user account inforamtion.
func (uc *usecase) GetMe(ctx context.Context, req *model.GetMeReqUsecase) (*model.GetMeResUsecase, error) {
	// Get user
	user, err := uc.storage.GetMe(ctx, &model.GetMeReqStorage{
		UserID: req.UserID,
	})
	if err != nil {
		return nil, errors.Wrap(err, "storage")
	}

	return &model.GetMeResUsecase{
		User: user.User,
	}, nil
}

// ChangePassword updates user password.
func (uc *usecase) ChangePassword(ctx context.Context, req *model.ChangePasswordReqUsecase) error {
	if err := req.Validate(); err != nil {
		return errors.Wrap(err, "request validation")
	}

	// Update password
	if err := uc.storage.ChangePassword(ctx, &model.ChangePasswordReqStorage{
		UserID:          req.UserID,
		OldPasswordHash: tools.SHA256(req.OldPassword),
		NewPasswordHash: tools.SHA256(req.NewPassword),
	}); err != nil {
		return errors.Wrap(err, "change password")
	}

	return nil
}

// DeleteMe deletes user account.
func (uc *usecase) DeleteMe(ctx context.Context, req *model.DeleteMeReqUsecase) (*model.DeleteMeResUsecase, error) {
	// Revoke user all refresh tokens
	if _, err := uc.cache.RevokeAllUserRefreshTokens(req.UserID); err != nil {
		return nil, errors.Wrap(err, "revoke all refresh tokens")
	}

	// Delete user
	user, err := uc.storage.DeleteMe(ctx, &model.DeleteMeReqStorage{
		UserID: req.UserID,
	})
	if err != nil {
		return nil, errors.Wrap(err, "storage")
	}

	return &model.DeleteMeResUsecase{
		User: user.User,
	}, nil
}
