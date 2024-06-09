package usecase

import (
	"coffeeshop-api/internal/service/users/model"
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
func (uc *usecase) Signup(req *model.UsecaseSignupReq) (*model.UsecaseSignupRes, error) {
	if err := req.Validate(); err != nil {
		return nil, errors.Wrap(err, "request validation")
	}

	resCreate, err := uc.storage.Create(&model.StorageCreateReq{
		Username:     req.Username,
		Email:        req.Email,
		Phone:        req.Phone,
		PasswordHash: tools.SHA256(req.Password),
	})
	if err != nil {
		return nil, errors.Wrap(err, "create user")
	}

	return &model.UsecaseSignupRes{
		UserID: resCreate.UserID,
	}, nil
}

// Signin checks credentials and generates tokens.
func (uc *usecase) Signin(req *model.UsecaseSigninReq) (*model.UsecaseSigninRes, error) {
	if err := req.Validate(); err != nil {
		return nil, errors.Wrap(err, "request validation")
	}

	// Check credentials
	resCheck, err := uc.storage.CheckCredentials(&model.StorageCheckCredentialsReq{
		Email:        req.Email,
		PasswordHash: tools.SHA256(req.Password),
	})
	if err != nil {
		return nil, errors.Wrap(err, "check credentials")
	}

	// Create new tokens pair
	accessToken, refreshToken, err := uc.token.CreatePair(&model.JWTClaims{
		UserID: resCheck.UserID,
	})
	if err != nil {
		return nil, errors.Wrap(err, "create access token")
	}

	// Save refresh token
	if err := uc.cache.SaveUserRefreshToken(resCheck.UserID, refreshToken); err != nil {
		return nil, errors.Wrap(err, "save refresh token")
	}

	return &model.UsecaseSigninRes{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

// Refresh revokes refresh token and updates tokens.
func (uc *usecase) Refresh(req *model.UsecaseRefreshReq) (*model.UsecaseRefreshRes, error) {
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
	resDeleted, err := uc.storage.IsDeleted(&model.StorageIsDeletedReq{
		UserID: claims.UserID,
	})
	if err != nil {
		return nil, errors.Wrap(err, "check if account deleted")
	}
	if resDeleted.Deleted {
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

	return &model.UsecaseRefreshRes{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

// SignoutAll revokes all refresh tokens.
func (uc *usecase) SignoutAll(req *model.UsecaseSignoutAllReq) (*model.UsecaseSignoutAllRes, error) {
	// Revoke all refresh tokens
	refreshTokens, err := uc.cache.RevokeAllUserRefreshTokens(req.UserID)
	if err != nil {
		return nil, errors.Wrap(err, "revoke all tokens")
	}

	return &model.UsecaseSignoutAllRes{
		RefreshTokens: refreshTokens,
	}, nil
}

// GetMe returns user account inforamtion.
func (uc *usecase) GetMe(req *model.UsecaseGetMeReq) (*model.UsecaseGetMeRes, error) {
	// Get user
	resMe, err := uc.storage.GetMe(&model.StorageGetMeReq{
		UserID: req.UserID,
	})
	if err != nil {
		return nil, errors.Wrap(err, "storage")
	}

	return &model.UsecaseGetMeRes{
		User: resMe.User,
	}, nil
}

// ChangePassword changes user password.
func (uc *usecase) ChangePassword(req *model.UsecaseChangePasswordReq) error {
	if err := req.Validate(); err != nil {
		return errors.Wrap(err, "request validation")
	}

	// Update password
	if err := uc.storage.ChangePassword(&model.StorageChangePasswordReq{
		UserID:          req.UserID,
		NewPasswordHash: tools.SHA256(req.NewPassword),
	}); err != nil {
		return errors.Wrap(err, "change password")
	}

	return nil
}

// DeleteMe deletes user account.
func (uc *usecase) DeleteMe(req *model.UsecaseDeleteMeReq) (*model.UsecaseDeleteMeRes, error) {
	// Delete user
	resDelete, err := uc.storage.DeleteMe(&model.StorageDeleteMeReq{
		UserID: req.UserID,
	})
	if err != nil {
		return nil, errors.Wrap(err, "storage")
	}

	return &model.UsecaseDeleteMeRes{
		User: resDelete.User,
	}, nil
}
