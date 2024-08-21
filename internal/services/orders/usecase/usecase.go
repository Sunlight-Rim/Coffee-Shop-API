package usecase

import (
	"coffeeshop-api/internal/services/orders/model"
	"coffeeshop-api/pkg/errors"
	"coffeeshop-api/pkg/tools"
	"context"
)

type usecase struct {
	storage model.IStorage
}

// New usecase.
func New(storage model.IStorage) *usecase {
	return &usecase{
		storage: storage,
	}
}

func (uc *usecase) ListOrders(ctx context.Context, req *model.ListOrdersReqUsecase) (*model.ListOrdersResUsecase, error) {
	orders, err := uc.storage.ListOrders(ctx, &model.ListOrdersReqStorage{
		UserID: req.UserID,
		Offset: req.Offset,
	})
	if err != nil {
		return nil, errors.Wrap(err, "list orders")
	}

	return &model.ListOrdersResUsecase{
		Orders: orders.Orders,
	}, nil
}
func (uc *usecase) GetOrderInfo(ctx context.Context, req *model.GetOrderInfoReqUsecase) (*model.GetOrderInfoResUsecase, error) {
	if err := req.Validate(); err != nil {
		return nil, errors.Wrap(err, "request validation")
	}

	order, err := uc.storage.GetOrderInfo(ctx, &model.GetOrderInfoReqStorage{
		UserID:  req.UserID,
		OrderID: req.OrderID,
	})
	if err != nil {
		return nil, errors.Wrap(err, "get order info")
	}

	return &model.GetOrderInfoResUsecase{
		Order: order.Order,
	}, nil
}

// CreateOrder creates order in database.
func (uc *usecase) CreateOrder(ctx context.Context, req *model.CreateOrderReqUsecase) (*model.CreateOrderResUsecase, error) {
	if err := req.Validate(); err != nil {
		return nil, errors.Wrap(err, "request validation")
	}

	// Check that all specified coffeeIDs and toppings are exists
	var (
		coffeeIDs = make(map[uint64]struct{})
		toppings  = make(map[string]struct{})
	)

	for i := range req.Items {
		coffeeIDs[req.Items[i].CoffeeID] = struct{}{}
		if req.Items[i].Topping != nil {
			toppings[*req.Items[i].Topping] = struct{}{}
		}
	}

	if err := uc.storage.CheckCoffeeIDsExists(ctx, &model.CheckCoffeeIDsExistsReqStorage{
		CoffeeIDs: tools.SliceOfMapKeys(coffeeIDs),
	}); err != nil {
		return nil, errors.Wrap(err, "check all coffeeIDs exists")
	}

	if err := uc.storage.CheckToppingsExists(ctx, &model.CheckToppingsExistsReqStorage{
		Toppings: tools.SliceOfMapKeys(toppings),
	}); err != nil {
		return nil, errors.Wrap(err, "check all toppings exists")
	}

	// Create order
	order, err := uc.storage.CreateOrder(ctx, &model.CreateOrderReqStorage{
		UserID:  req.UserID,
		Address: req.Address,
		Items:   req.Items,
	})
	if err != nil {
		return nil, errors.Wrap(err, "create order")
	}

	return &model.CreateOrderResUsecase{
		OrderID: order.OrderID,
	}, nil
}

func (uc *usecase) CancelOrder(ctx context.Context, req *model.CancelOrderReqUsecase) (*model.CancelOrderResUsecase, error) {
	order, err := uc.storage.CancelOrder(ctx, &model.CancelOrderReqStorage{
		UserID:  req.UserID,
		OrderID: req.OrderID,
	})
	if err != nil {
		return nil, errors.Wrap(err, "cancel order")
	}

	return &model.CancelOrderResUsecase{
		OrderCustomerID: order.OrderCustomerID,
		OrderID:         order.OrderID,
		OrderCreatedAt:  order.OrderCreatedAt,
		OrderStatus:     order.OrderStatus,
	}, nil
}

// EmployeeCompleteOrder marks order as completed.
func (uc *usecase) EmployeeCompleteOrder(ctx context.Context, req *model.EmployeeCompleteOrderReqUsecase) (*model.EmployeeCompleteOrderResUsecase, error) {
	order, err := uc.storage.EmployeeCompleteOrder(ctx, &model.EmployeeCompleteOrderReqStorage{
		OrderID: req.OrderID,
	})
	if err != nil {
		return nil, errors.Wrap(err, "complete order")
	}

	return &model.EmployeeCompleteOrderResUsecase{
		OrderCustomerID: order.OrderCustomerID,
		OrderID:         order.OrderID,
		OrderCreatedAt:  order.OrderCreatedAt,
		OrderStatus:     order.OrderStatus,
	}, nil
}
