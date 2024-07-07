package usecase

// type usecase struct {
// 	logger  model.ILogger
// 	storage model.IStorage
// }

// // New usecase.
// func New(logger model.ILogger, storage model.IStorage) *usecase {
// 	return &usecase{
// 		logger:  logger,
// 		storage: storage,
// 	}
// }

// // CreateOrder creates order in database.
// func (uc *usecase) CreateOrder(req *model.CreateOrderReqUsecase) (*model.CreateOrderResUsecase, error) {
// 	orderInfo, err := uc.storage.CreateOrder(&model.CreateOrderReqStorage{
// 		Orders: req.Orders,
// 	})
// 	if err != nil {
// 		return nil, errors.Wrap(err, "create order")
// 	}

// 	return &model.CreateOrderResUsecase{
// 		OrderID: orderInfo.OrderID,
// 	}, nil
// }
