package storage

import (
	"database/sql"

	"coffeeshop-api/internal/services/orders/model"
	"coffeeshop-api/pkg/errors"

	logger "github.com/sirupsen/logrus"
)

type storage struct {
	db *sql.DB
}

func New(db *sql.DB) *storage {
	return &storage{db: db}
}

func (s *storage) ListOrders(*model.ListOrdersReqStorage) (*model.ListOrdersResStorage, error) {
	return nil, nil
}
func (s *storage) GetOrderInfo(*model.GetOrderInfoReqStorage) (*model.GetOrderInfoResStorage, error) {
	return nil, nil
}

func (s *storage) CreateOrder(req *model.CreateOrderReqStorage) (*model.CreateOrderResStorage, error) {
	tx, err := s.db.Begin()
	if err != nil {
		return nil, errors.Wrap(err, "start transaction")
	}

	// Manage transaction
	defer func() {
		if err != nil {
			if errRollback := tx.Rollback(); errRollback != nil {
				logger.Errorf("Rollback error: %v", errRollback)
			}
		} else {
			if errCommit := tx.Commit(); errCommit != nil {
				logger.Errorf("Commit error: %v", errCommit)
			}
		}
	}()

	// Insert order

	var orderID uint64

	if err := tx.QueryRow(`
		INSERT INTO api.orders(
			user_id,
			"address"
		)
		VALUES ($1, $2)
		RETURNING id
	`,
		req.UserID,
		req.Address,
	).Scan(&orderID); err != nil {
		return nil, errors.Wrap(err, "insert order")
	}

	// Insert items

	stmt, err := tx.Prepare(`
		INSERT INTO api.order_items(
			order_id,
			coffee_id,
			topping
		)
		VALUES ($1, $2, $3)
	`)
	if err != nil {
		return nil, errors.Wrap(err, "prepare items")
	}

	for i := range req.Items {
		if _, err := stmt.Exec(
			orderID,
			req.Items[i].CoffeeID,
			&req.Items[i].Topping,
		); err != nil {
			return nil, errors.Wrap(err, "insert item")
		}
	}

	return &model.CreateOrderResStorage{OrderID: orderID}, nil
}

func (s *storage) CancelOrder(*model.CancelOrderReqStorage) (*model.CancelOrderResStorage, error) {
	return nil, nil
}

func (s *storage) EmployeeCompleteOrder(req *model.EmployeeCompleteOrderReqStorage) (*model.EmployeeCompleteOrderResStorage, error) {
	var orderInfo model.EmployeeCompleteOrderResStorage

	if err := s.db.QueryRow(`
		UPDATE api.orders
		SET "status" = 'ready to receive'
		WHERE id = $1
		RETURNING
			id,
			user_id,
			created_at,
			"status"
	`,
		req.OrderID,
	).Scan(
		&orderInfo.OrderID,
		&orderInfo.OrderCustomerID,
		&orderInfo.OrderCreatedAt,
		&orderInfo.OrderStatus,
	); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.Wrap(errors.OrderNotExists, "order not found")
		}

		return nil, errors.Wrap(err, "change order status")
	}

	return &orderInfo, nil
}
