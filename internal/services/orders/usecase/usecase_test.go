package usecase

import (
	"testing"

	"coffeeshop-api/internal/services/orders/model"
	"coffeeshop-api/internal/services/orders/model/mocks"
	"coffeeshop-api/pkg/errors"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func Test_ListOrders(t *testing.T) {
	t.Run("listOrders success", func(t *testing.T) {
		t.Parallel()

		mStorage := mocks.NewIStorage(t)

		uc := New(mStorage)

		// Data
		var (
			userID = uint64(123)
			offset = uint64(123)

			orders = []model.ListOrdersOrder{
				{OrderID: 1},
				{OrderID: 2},
				{OrderID: 3},
			}

			expectedRes = &model.ListOrdersResUsecase{Orders: orders}
		)

		// Mocks calls
		mStorage.On("ListOrders", mock.Anything, &model.ListOrdersReqStorage{
			UserID: userID,
			Offset: offset,
		}).Return(&model.ListOrdersResStorage{Orders: orders}, nil)

		// Testing method
		res, err := uc.ListOrders(t.Context(), &model.ListOrdersReqUsecase{
			UserID: userID,
			Offset: offset,
		})

		// Validation
		require.NoError(t, err)
		require.Equal(t, expectedRes, res)
	})

	t.Run("listOrders fail", func(t *testing.T) {
		t.Parallel()

		mStorage := mocks.NewIStorage(t)

		uc := New(mStorage)

		// Data
		var (
			userID = uint64(123)
			offset = uint64(123)

			expectedErr = errors.New("test error")
		)

		// Mocks calls
		mStorage.On("ListOrders", mock.Anything, &model.ListOrdersReqStorage{
			UserID: userID,
			Offset: offset,
		}).Return(nil, expectedErr)

		// Testing method
		res, err := uc.ListOrders(t.Context(), &model.ListOrdersReqUsecase{
			UserID: userID,
			Offset: offset,
		})

		// Validation
		require.Nil(t, res)
		require.ErrorIs(t, err, expectedErr)
	})
}
