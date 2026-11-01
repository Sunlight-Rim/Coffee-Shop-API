package usecase

import (
	"testing"

	"coffeeshop-api/internal/services/coffee/model"
	"coffeeshop-api/internal/services/coffee/model/mocks"
	"coffeeshop-api/pkg/errors"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func Test_ListCoffees(t *testing.T) {
	t.Run("listCoffees success", func(t *testing.T) {
		t.Parallel()

		mStorage := mocks.NewIStorage(t)

		uc := New(mStorage)

		// Data
		var (
			offset = uint64(123)

			coffees = []model.Coffee{
				{ID: 1},
				{ID: 2},
				{ID: 3},
			}

			expectedRes = &model.ListCoffeesResUsecase{CoffeeList: coffees}
		)

		// Mocks calls
		mStorage.On("ListCoffees", mock.Anything, &model.ListCoffeesReqStorage{
			Offset: offset,
		}).Return(&model.ListCoffeesResStorage{CoffeeList: coffees}, nil)

		// Testing method
		res, err := uc.ListCoffees(t.Context(), &model.ListCoffeesReqUsecase{
			Offset: offset,
		})

		// Validation
		require.NoError(t, err)
		require.Equal(t, expectedRes, res)
	})

	t.Run("listCoffees fail", func(t *testing.T) {
		t.Parallel()

		mStorage := mocks.NewIStorage(t)

		uc := New(mStorage)

		// Data
		var (
			offset = uint64(123)

			expectedErr = errors.New("test error")
		)

		// Mocks calls
		mStorage.On("ListCoffees", mock.Anything, &model.ListCoffeesReqStorage{
			Offset: offset,
		}).Return(nil, expectedErr)

		// Testing method
		res, err := uc.ListCoffees(t.Context(), &model.ListCoffeesReqUsecase{
			Offset: offset,
		})

		// Validation
		require.Nil(t, res)
		require.ErrorIs(t, err, expectedErr)
	})
}
