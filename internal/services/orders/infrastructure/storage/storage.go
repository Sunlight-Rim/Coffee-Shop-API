package storage

import (
	"database/sql"

	"coffeeshop-api/internal/services/orders/model"
	"coffeeshop-api/pkg/errors"

	"github.com/sirupsen/logrus"
)

type storage struct {
	db *sql.DB
}

func New(db *sql.DB) *storage {
	return &storage{db: db}
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
				logrus.Errorf("Rollback error: %v. Query error: %v", errRollback, err)
			}
		} else {
			if errCommit := tx.Commit(); errCommit != nil {
				logrus.Errorf("Commit error: %v", errCommit)
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
		return nil, errors.Wrap(err, "prepare")
	}

	for i := range req.Items {
		if _, err := stmt.Exec(
			orderID,
			req.Items[i].CoffeeID,
			req.Items[i].Topping,
		); err != nil {
			return nil, errors.Wrap(err, "insert item")
		}
	}

	return &model.CreateOrderResStorage{OrderID: orderID}, nil
}
