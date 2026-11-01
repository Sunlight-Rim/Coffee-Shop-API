package usecase

import (
	"strings"
	"testing"
	"time"

	"coffeeshop-api/internal/services/users/model"
	"coffeeshop-api/internal/services/users/model/mocks"
	"coffeeshop-api/pkg/claims"
	"coffeeshop-api/pkg/errors"
	"coffeeshop-api/pkg/tools"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func Test_Signup(t *testing.T) {
	t.Run("signup success", func(t *testing.T) {
		t.Parallel()

		mStorage := mocks.NewIStorage(t)
		mCache := mocks.NewICache(t)
		mToken := mocks.NewIToken(t)

		uc := New(mStorage, mCache, mToken)

		// Data
		var (
			username = "username"
			email    = "email@test.com"
			phone    = uint64(1234567890)
			password = "password.123"
			userID   = uint64(123)

			expectedRes = &model.SignupResUsecase{
				UserID: userID,
			}
		)

		// Mocks calls
		mStorage.On("CreateUser", mock.Anything, &model.CreateUserReqStorage{
			Username:     username,
			Phone:        phone,
			Email:        strings.ToLower(email),
			PasswordHash: tools.SHA256(password),
		}).Return(&model.CreateUserResStorage{UserID: userID}, nil)

		// Testing method
		res, err := uc.Signup(t.Context(), &model.SignupReqUsecase{
			Username: username,
			Email:    email,
			Phone:    phone,
			Password: password,
		})

		// Validation
		require.NoError(t, err)
		require.Equal(t, expectedRes, res)
	})

	t.Run("signup fail", func(t *testing.T) {
		t.Parallel()

		mStorage := mocks.NewIStorage(t)
		mCache := mocks.NewICache(t)
		mToken := mocks.NewIToken(t)

		uc := New(mStorage, mCache, mToken)

		// Data
		var (
			username = "username"
			email    = "email@test.com"
			phone    = uint64(1234567890)
			password = "password.123"

			expectedErr = errors.EmailIsBusy
		)

		// Mocks calls
		mStorage.On("CreateUser", mock.Anything, &model.CreateUserReqStorage{
			Username:     username,
			Phone:        phone,
			Email:        strings.ToLower(email),
			PasswordHash: tools.SHA256(password),
		}).Return(nil, errors.EmailIsBusy)

		// Testing method
		res, err := uc.Signup(t.Context(), &model.SignupReqUsecase{
			Username: username,
			Email:    email,
			Phone:    phone,
			Password: password,
		})

		// Validation
		require.ErrorIs(t, err, expectedErr)
		require.Nil(t, res)
	})
}

func Test_Signin(t *testing.T) {
	t.Run("signin successful", func(t *testing.T) {
		t.Parallel()

		mStorage := mocks.NewIStorage(t)
		mCache := mocks.NewICache(t)
		mToken := mocks.NewIToken(t)

		uc := New(mStorage, mCache, mToken)

		// Data
		var (
			email        = "email@test.com"
			password     = "password.123"
			userID       = uint64(123)
			accessToken  = "access.token.here"
			refreshToken = "refresh.token.here"

			tokensPair = model.TokensPair{
				AccessToken:  model.Token{String: accessToken, Exp: time.Now().Add(time.Minute)},
				RefreshToken: model.Token{String: refreshToken, Exp: time.Now().Add(time.Minute)},
			}
		)

		// Mocks calls
		mStorage.On("CheckCredentials", mock.Anything, &model.CheckCredentialsReqStorage{
			Email:        strings.ToLower(email),
			PasswordHash: tools.SHA256(password),
		}).Return(&model.CheckCredentialsResStorage{UserID: userID}, nil)

		mToken.On("CreatePair", &claims.Claims{
			UserID: userID,
		}).Return(&tokensPair, nil)

		mCache.On("SaveUserRefreshToken", mock.Anything, userID, tokensPair.RefreshToken).Return(nil)

		// Testing method
		res, err := uc.Signin(t.Context(), &model.SigninReqUsecase{
			Email:    email,
			Password: password,
		})

		// Validation
		require.NoError(t, err)
		require.NotNil(t, res)
		require.Equal(t, accessToken, res.TokensPair.AccessToken.String)
		require.Equal(t, refreshToken, res.TokensPair.RefreshToken.String)
	})

	t.Run("signin invalid credentials", func(t *testing.T) {
		t.Parallel()

		mStorage := mocks.NewIStorage(t)
		mCache := mocks.NewICache(t)
		mToken := mocks.NewIToken(t)

		uc := New(mStorage, mCache, mToken)

		// Data
		var (
			email       = "email@test.com"
			password    = "wrong.password"
			expectedErr = errors.InvalidCredentials
		)

		// Mocks calls
		mStorage.On("CheckCredentials", mock.Anything, &model.CheckCredentialsReqStorage{
			Email:        strings.ToLower(email),
			PasswordHash: tools.SHA256(password),
		}).Return(nil, expectedErr)

		// Testing method
		res, err := uc.Signin(t.Context(), &model.SigninReqUsecase{
			Email:    email,
			Password: password,
		})

		// Validation
		require.ErrorIs(t, err, expectedErr)
		require.Nil(t, res)
	})
}

func Test_GetMe(t *testing.T) {
	t.Run("get me successful", func(t *testing.T) {
		t.Parallel()

		mStorage := mocks.NewIStorage(t)
		mCache := mocks.NewICache(t)
		mToken := mocks.NewIToken(t)

		uc := New(mStorage, mCache, mToken)

		// Data
		var (
			userID = uint64(123)

			user = &model.User{
				ID:       userID,
				Username: "testuser",
				Email:    "test@example.com",
				Phone:    1234567890,
			}
		)

		// Mocks calls
		mStorage.On("GetMe", mock.Anything, &model.GetMeReqStorage{
			UserID: userID,
		}).Return(&model.GetMeResStorage{User: user}, nil)

		// Testing method
		res, err := uc.GetMe(t.Context(), &model.GetMeReqUsecase{
			UserID: userID,
		})

		// Validation
		require.NoError(t, err)
		require.NotNil(t, res)
		require.Equal(t, user, res.User)
	})

	t.Run("get me error", func(t *testing.T) {
		t.Parallel()

		mStorage := mocks.NewIStorage(t)
		mCache := mocks.NewICache(t)
		mToken := mocks.NewIToken(t)

		uc := New(mStorage, mCache, mToken)

		// Data
		var (
			userID = uint64(123)
		)

		// Mocks calls
		mStorage.On("GetMe", mock.Anything, &model.GetMeReqStorage{
			UserID: userID,
		}).Return(nil, errors.New("user not found"))

		// Testing method
		res, err := uc.GetMe(t.Context(), &model.GetMeReqUsecase{
			UserID: userID,
		})

		// Validation
		require.Error(t, err)
		require.Nil(t, res)
	})
}
