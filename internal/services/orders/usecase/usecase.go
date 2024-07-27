package usecase

import (
	"coffeeshop-api/internal/services/orders/model"
	"coffeeshop-api/pkg/errors"
	"coffeeshop-api/pkg/tools"
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

func (uc *usecase) ListOrders(req *model.ListOrdersReqUsecase) (*model.ListOrdersResUsecase, error) {
	orders, err := uc.storage.ListOrders(&model.ListOrdersReqStorage{
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
func (uc *usecase) GetOrderInfo(req *model.GetOrderInfoReqUsecase) (*model.GetOrderInfoResUsecase, error) {
	if err := req.Validate(); err != nil {
		return nil, errors.Wrap(err, "request validation")
	}

	order, err := uc.storage.GetOrderInfo(&model.GetOrderInfoReqStorage{
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
func (uc *usecase) CreateOrder(req *model.CreateOrderReqUsecase) (*model.CreateOrderResUsecase, error) {
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

	if err := uc.storage.CheckAllCoffeeIDsExists(&model.CheckAllCoffeeIDsExistsReqStorage{
		CoffeeIDs: tools.SliceOfMapKeys(coffeeIDs),
	}); err != nil {
		return nil, errors.Wrap(err, "check coffeeIDs exists")
	}

	if err := uc.storage.CheckAllToppingsExists(&model.CheckAllToppingsExistsReqStorage{
		Toppings: tools.SliceOfMapKeys(toppings),
	}); err != nil {
		return nil, errors.Wrap(err, "check toppings exists")
	}

	order, err := uc.storage.CreateOrder(&model.CreateOrderReqStorage{
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

func (uc *usecase) CancelOrder(*model.CancelOrderReqUsecase) (*model.CancelOrderResUsecase, error) {
	return nil, nil
}

// EmployeeCompleteOrder marks order as completed.
func (uc *usecase) EmployeeCompleteOrder(req *model.EmployeeCompleteOrderReqUsecase) (*model.EmployeeCompleteOrderResUsecase, error) {
	order, err := uc.storage.EmployeeCompleteOrder(&model.EmployeeCompleteOrderReqStorage{
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
