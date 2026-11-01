package usecase

import (
	"strings"
	"testing"

	"coffeeshop-api/internal/services/users/mocks"
	"coffeeshop-api/internal/services/users/model"
	"coffeeshop-api/pkg/tools"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestSignup(t *testing.T) {
	mStorage := mocks.NewIStorage(t)
	mCache := mocks.NewICache(t)
	mToken := mocks.NewIToken(t)

	uc := New(mStorage, mCache, mToken)

	t.Run("signup", func(t *testing.T) {
		// Data
		var (
			username = "username"
			email    = "email@test.com"
			phone    = uint64(1234567890)
			password = "password.123"
			userID   = uint64(123)

			expectedRes = model.SignupResUsecase{
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
		require.Equal(t, expectedRes, *res)
	})
}
