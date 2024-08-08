package storage

import (
	"database/sql"

	"coffeeshop-api/internal/services/orders/model"
	"coffeeshop-api/pkg/errors"

	"github.com/lib/pq"
	logger "github.com/sirupsen/logrus"
)

const ordersLimit = 10

type storage struct {
	db *sql.DB
}

func New(db *sql.DB) *storage {
	return &storage{db: db}
}

func (s *storage) ListOrders(req *model.ListOrdersReqStorage) (*model.ListOrdersResStorage, error) {
	rows, err := s.db.Query(`
		SELECT
			id,
			created_at
		FROM api.orders
		WHERE user_id = $1
		ORDER BY created_at DESC
		LIMIT $2 OFFSET $3
	`,
		req.UserID,
		ordersLimit,
		req.Offset,
	)
	if err != nil {
		return nil, errors.Wrap(err, "get orders list")
	}

	defer func() {
		if err := rows.Close(); err != nil {
			logger.WithField("error", err).Error("Close rows")
		}
	}()

	var (
		orders []model.ListOrdersOrder
		order  model.ListOrdersOrder
	)

	for rows.Next() {
		if err := rows.Scan(
			&order.OrderID,
			&order.OrderCreatedAt,
		); err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return &model.ListOrdersResStorage{}, nil
			}

			return nil, errors.Wrap(err, "scan orders list")
		}

		orders = append(orders, order)
	}

	return &model.ListOrdersResStorage{Orders: orders}, nil
}

func (s *storage) GetOrderInfo(req *model.GetOrderInfoReqStorage) (*model.GetOrderInfoResStorage, error) {
	var order model.GetOrderInfoOrder

	// Get order
	if err := s.db.QueryRow(`
		SELECT
			id,
			"status",
			"address",
			created_at
		FROM api.orders
		WHERE
			user_id = $1 AND
			id = $2
	`,
		req.UserID,
		req.OrderID,
	).Scan(
		&order.OrderID,
		&order.Status,
		&order.Address,
		&order.CreatedAt,
	); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.Wrap(errors.OrderNotExists, "order not found")
		}

		return nil, errors.Wrap(err, "get order info")
	}

	// Get order items
	rows, err := s.db.Query(`
		SELECT
			c.id,
			c.title,
			c."image",
			oi.topping
		FROM api.order_items AS oi
		JOIN api.coffee AS c ON oi.coffee_id = c.id
		WHERE oi.order_id = $1
	`,
		req.OrderID,
	)
	if err != nil {
		return nil, errors.Wrap(err, "get order items")
	}

	defer func() {
		if err := rows.Close(); err != nil {
			logger.WithField("error", err).Error("Close rows")
		}
	}()

	var item model.GetOrderInfoOrderItem

	for rows.Next() {
		if err := rows.Scan(
			&item.CoffeeID,
			&item.CoffeeTitle,
			&item.CoffeeImage,
			&item.Topping,
		); err != nil {
			return nil, errors.Wrap(err, "scan order item")
		}

		order.Items = append(order.Items, item)
	}

	return &model.GetOrderInfoResStorage{Order: order}, nil
}

func (s *storage) CheckCoffeeIDsExists(req *model.CheckCoffeeIDsExistsReqStorage) error {
	var count int

	if err := s.db.QueryRow(`
		SELECT count(*)
		FROM api.coffee
		WHERE id = any($1::INT[])
	`,
		pq.Array(req.CoffeeIDs),
	).Scan(&count); err != nil {
		return errors.Wrap(err, "scan coffee IDs count")
	}

	if count != len(req.CoffeeIDs) {
		return errors.Wrap(errors.CoffeeNotExists, "requested coffee ID not found")
	}

	return nil
}

func (s *storage) CheckToppingsExists(req *model.CheckToppingsExistsReqStorage) error {
	var count int

	if err := s.db.QueryRow(`
		SELECT count(*)
		FROM unnest(enum_range(NULL::api.topping)) AS t(name)
		WHERE t.name::TEXT = any($1::TEXT[])
	`,
		pq.Array(req.Toppings),
	).Scan(&count); err != nil {
		return errors.Wrap(err, "scan toppings count")
	}

	if count != len(req.Toppings) {
		return errors.Wrap(errors.ToppingNotExists, "requested topping not found")
	}

	return nil
}

func (s *storage) CreateOrder(req *model.CreateOrderReqStorage) (_ *model.CreateOrderResStorage, err error) {
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
		if _, err = stmt.Exec(
			orderID,
			req.Items[i].CoffeeID,
			&req.Items[i].Topping,
		); err != nil {
			return nil, errors.Wrap(err, "insert item")
		}
	}

	return &model.CreateOrderResStorage{OrderID: orderID}, nil
}

func (s *storage) CancelOrder(req *model.CancelOrderReqStorage) (*model.CancelOrderResStorage, error) {
	var order model.CancelOrderResStorage

	if err := s.db.QueryRow(`
		SELECT
			id,
			user_id,
			created_at,
			"status"
		FROM api.orders
		WHERE id = $1
	`,
		req.OrderID,
	).Scan(
		&order.OrderID,
		&order.OrderCustomerID,
		&order.OrderCreatedAt,
		&order.OrderStatus,
	); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.Wrap(errors.OrderNotExists, "order not found")
		}

		return nil, errors.Wrap(err, "get order")
	}

	if order.OrderStatus != "cooking" {
		return nil, errors.Wrap(errors.CannotCancelThisOrder, "order cannot be cancelled")
	}

	if err := s.db.QueryRow(`
		UPDATE api.orders
		SET "status" = 'cancelled'
		WHERE id = $1
		RETURNING "status"
	`,
		req.OrderID,
	).Scan(&order.OrderStatus); err != nil {
		return nil, errors.Wrap(err, "change order status")
	}

	return &order, nil
}

func (s *storage) EmployeeCompleteOrder(req *model.EmployeeCompleteOrderReqStorage) (*model.EmployeeCompleteOrderResStorage, error) {
	var order model.EmployeeCompleteOrderResStorage

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
		&order.OrderID,
		&order.OrderCustomerID,
		&order.OrderCreatedAt,
		&order.OrderStatus,
	); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.Wrap(errors.OrderNotExists, "order not found")
		}

		return nil, errors.Wrap(err, "change order status")
	}

	return &order, nil
}
