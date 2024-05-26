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

func New(logger model.ILogger, storage model.IStorage, cache model.ICache, token model.IToken) *usecase {
	return &usecase{
		logger:  logger,
		storage: storage,
		cache:   cache,
		token:   token,
	}
}

func (uc *usecase) Signup(req *model.SignupReq) (*model.SignupRes, error) {
	var res model.SignupRes

	if err := req.Validate(); err != nil {
		return nil, errors.Wrap(err, "request validation")
	}

	// Create user
	resStore, err := uc.storage.Create(&model.StorageCreateReq{
		Username:     req.Username,
		Email:        req.Email,
		Phone:        req.Phone,
		PasswordHash: tools.SHA256(req.Password),
	})
	if err != nil {
		return nil, errors.Wrap(err, "storage")
	}

	res.UserID = resStore.UserID

	return &res, nil
}

func (uc *usecase) Signin(req *model.SigninReq) (*model.SigninRes, error) {
	var res model.SigninRes

	if err := req.Validate(); err != nil {
		return nil, errors.Wrap(err, "request validation")
	}

	// Check credentials
	resStore, err := uc.storage.CheckCredentials(&model.StorageCheckCredentialsReq{
		Email:        req.Email,
		PasswordHash: tools.SHA256(req.Password),
	})
	if err != nil {
		return nil, errors.Wrap(err, "check credentials")
	}

	// Create new tokens pair
	accessToken, refreshToken, err := uc.token.CreatePair(&model.JWTClaims{
		UserID: resStore.UserID,
	})
	if err != nil {
		return nil, errors.Wrap(err, "create access token")
	}

	// Save refresh token
	if err := uc.cache.SaveUserRefreshToken(resStore.UserID, refreshToken); err != nil {
		return nil, errors.Wrap(err, "save refresh token")
	}

	res.AccessToken = accessToken
	res.RefreshToken = refreshToken

	return &res, nil
}

func (uc *usecase) Refresh(req *model.RefreshReq) (*model.RefreshRes, error) {
	var res model.RefreshRes

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
	resStore, err := uc.storage.IsDeleted(&model.StorageIsDeletedReq{
		UserID: claims.UserID,
	})
	if err != nil {
		return nil, errors.Wrap(err, "check if account deleted")
	}
	if resStore.Deleted {
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

	res.AccessToken = accessToken
	res.RefreshToken = refreshToken

	return &res, nil
}

func (uc *usecase) SignoutAll(req *model.SignoutAllReq) (*model.SignoutAllRes, error) {
	var res model.SignoutAllRes

	// Revoke all refresh tokens
	tokens, err := uc.cache.RevokeAllUserRefreshTokens(req.UserID)
	if err != nil {
		return nil, errors.Wrap(err, "revoke all tokens")
	}

	res.RefreshTokens = tokens

	return &res, nil
}

func (uc *usecase) GetMe(req *model.GetMeReq) (*model.GetMeRes, error) {
	var res model.GetMeRes

	// Get user
	resStore, err := uc.storage.GetMe(&model.StorageGetMeReq{
		UserID: req.UserID,
	})
	if err != nil {
		return nil, errors.Wrap(err, "storage")
	}

	res.User = resStore.User

	return &res, nil
}

func (uc *usecase) ChangePassword(req *model.ChangePasswordReq) error {
	if err := req.Validate(); err != nil {
		return errors.Wrap(err, "request validation")
	}

	// Update password
	if err := uc.storage.ChangePassword(&model.StorageChangePasswordReq{
		UserID:          req.UserID,
		NewPasswordHash: tools.SHA256(req.NewPassword),
	}); err != nil {
		return errors.Wrap(err, "storage")
	}

	return nil
}

func (uc *usecase) DeleteMe(req *model.DeleteMeReq) (*model.DeleteMeRes, error) {
	var res model.DeleteMeRes

	// Delete user
	resStore, err := uc.storage.DeleteMe(&model.StorageDeleteMeReq{
		UserID: req.UserID,
	})
	if err != nil {
		return nil, errors.Wrap(err, "storage")
	}

	res.User = resStore.User

	return &res, nil
}
