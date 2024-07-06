package usecase

import (
	"coffeeshop-api/internal/services/users/model"
	"coffeeshop-api/pkg/errors"
	"coffeeshop-api/pkg/tools"
)

type usecase struct {
	logger  model.ILogger
	storage model.IStorage
	cache   model.ICache
	token   model.IToken
}

// New usecase.
func New(logger model.ILogger, storage model.IStorage, cache model.ICache, token model.IToken) *usecase {
	return &usecase{
		logger:  logger,
		storage: storage,
		cache:   cache,
		token:   token,
	}
}

// Signup creates user in database.
func (uc *usecase) Signup(req *model.SignupReqUsecase) (*model.SignupResUsecase, error) {
	if err := req.Validate(); err != nil {
		return nil, errors.Wrap(err, "request validation")
	}

	userInfo, err := uc.storage.CreateUser(&model.CreateUserReqStorage{
		Username:     req.Username,
		Email:        req.Email,
		Phone:        req.Phone,
		PasswordHash: tools.SHA256(req.Password),
	})
	if err != nil {
		return nil, errors.Wrap(err, "create user")
	}

	return &model.SignupResUsecase{
		UserID: userInfo.UserID,
	}, nil
}

// Signin checks credentials and generates tokens.
func (uc *usecase) Signin(req *model.SigninReqUsecase) (*model.SigninResUsecase, error) {
	if err := req.Validate(); err != nil {
		return nil, errors.Wrap(err, "request validation")
	}

	// Check credentials
	userInfo, err := uc.storage.CheckCredentials(&model.CheckCredentialsReqStorage{
		Email:        req.Email,
		PasswordHash: tools.SHA256(req.Password),
	})
	if err != nil {
		return nil, errors.Wrap(err, "check credentials")
	}

	// Create new tokens pair
	accessToken, refreshToken, err := uc.token.CreatePair(&model.Claims{
		UserID: userInfo.UserID,
	})
	if err != nil {
		return nil, errors.Wrap(err, "create access token")
	}

	// Save refresh token
	if err := uc.cache.SaveUserRefreshToken(userInfo.UserID, refreshToken); err != nil {
		return nil, errors.Wrap(err, "save refresh token")
	}

	return &model.SigninResUsecase{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

// Refresh revokes refresh token and updates tokens.
func (uc *usecase) Refresh(req *model.RefreshReqUsecase) (*model.RefreshResUsecase, error) {
	// Parse refresh token
	claims, err := uc.token.Parse(req.RefreshToken)
	if err != nil {
		return nil, errors.Wrap(err, "parse token")
	}

	// Revoke refresh token
	if err := uc.cache.RevokeUserRefreshToken(claims.UserID, req.RefreshToken); err != nil {
		return nil, errors.Wrap(err, "revoke token")
	}

	// Check if user account is not deleted
	deleted, err := uc.storage.IsDeleted(&model.IsDeletedReqStorage{
		UserID: claims.UserID,
	})
	if err != nil {
		return nil, errors.Wrap(err, "check if account deleted")
	}
	if deleted.Deleted {
		return nil, errors.Wrap(errors.DeletedAccount, "account was deleted")
	}

	// Create new tokens pair
	accessToken, refreshToken, err := uc.token.CreatePair(claims)
	if err != nil {
		return nil, errors.Wrap(err, "create token")
	}

	// Save refresh token
	if err := uc.cache.SaveUserRefreshToken(claims.UserID, refreshToken); err != nil {
		return nil, errors.Wrap(err, "save refresh token")
	}

	return &model.RefreshResUsecase{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

// SignoutAll revokes all refresh tokens.
func (uc *usecase) SignoutAll(req *model.SignoutAllReqUsecase) (*model.SignoutAllResUsecase, error) {
	// Revoke all refresh tokens
	refreshTokens, err := uc.cache.RevokeAllUserRefreshTokens(req.UserID)
	if err != nil {
		return nil, errors.Wrap(err, "revoke all tokens")
	}

	return &model.SignoutAllResUsecase{
		RefreshTokens: refreshTokens,
	}, nil
}

// GetMe returns user account inforamtion.
func (uc *usecase) GetMe(req *model.GetMeReqUsecase) (*model.GetMeResUsecase, error) {
	// Get user
	userInfo, err := uc.storage.GetMe(&model.GetMeReqStorage{
		UserID: req.UserID,
	})
	if err != nil {
		return nil, errors.Wrap(err, "storage")
	}

	return &model.GetMeResUsecase{
		User: userInfo.User,
	}, nil
}

// ChangePassword updates user password.
func (uc *usecase) ChangePassword(req *model.ChangePasswordReqUsecase) error {
	if err := req.Validate(); err != nil {
		return errors.Wrap(err, "request validation")
	}

	// Update password
	if err := uc.storage.ChangePassword(&model.ChangePasswordReqStorage{
		UserID:          req.UserID,
		OldPasswordHash: tools.SHA256(req.OldPassword),
		NewPasswordHash: tools.SHA256(req.NewPassword),
	}); err != nil {
		return errors.Wrap(err, "change password")
	}

	return nil
}

// DeleteMe deletes user account.
func (uc *usecase) DeleteMe(req *model.DeleteMeReqUsecase) (*model.DeleteMeResUsecase, error) {
	// Delete user
	userInfo, err := uc.storage.DeleteMe(&model.DeleteMeReqStorage{
		UserID: req.UserID,
	})
	if err != nil {
		return nil, errors.Wrap(err, "storage")
	}

	return &model.DeleteMeResUsecase{
		User: userInfo.User,
	}, nil
}
