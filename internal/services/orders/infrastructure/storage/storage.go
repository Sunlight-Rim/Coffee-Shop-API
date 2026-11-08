package storage

import (
	"context"
	"database/sql"

	"coffeeshop-api/internal/services/orders/model"
	"coffeeshop-api/pkg/errors"
	"coffeeshop-api/pkg/tracer"

	sq "github.com/Masterminds/squirrel"
	"github.com/lib/pq"
	logger "github.com/sirupsen/logrus"
)

var builder = sq.StatementBuilder.PlaceholderFormat(sq.Dollar)

const ordersLimit = 10

type storage struct {
	db *sql.DB
}

func New(db *sql.DB) *storage {
	return &storage{db: db}
}

func (s *storage) ListOrders(ctx context.Context, req *model.ListOrdersReqStorage) (*model.ListOrdersResStorage, error) {
	ctx, span := tracer.Start(ctx, "ListOrders")
	defer span.End()

	query, args := builder.
		Select(
			"id",
			"created_at",
		).
		From("api.orders").
		Where(sq.Eq{"user_id": req.UserID}).
		OrderBy("created_at DESC").
		Limit(ordersLimit).
		Offset(req.Offset).
		MustSql()

	rows, err := s.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, errors.Wrap(err, "get orders list")
	}

	defer func() {
		if errClose := rows.Close(); errClose != nil {
			logger.Errorf("Close rows: %s", errClose)
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
			return nil, errors.Wrap(err, "scan orders list")
		}

		orders = append(orders, order)
	}

	if err := rows.Err(); err != nil {
		return nil, errors.Wrap(err, "rows error")
	}

	return &model.ListOrdersResStorage{Orders: orders}, nil
}

func (s *storage) GetOrderInfo(ctx context.Context, req *model.GetOrderInfoReqStorage) (*model.GetOrderInfoResStorage, error) {
	ctx, span := tracer.Start(ctx, "GetOrderInfo")
	defer span.End()

	var order model.GetOrderInfoOrder

	// Get order
	query, args := builder.
		Select(
			"id",
			"status",
			"address",
			"created_at",
		).
		From("api.orders").
		Where(sq.Eq{
			"user_id": req.UserID,
			"id":      req.OrderID,
		}).
		MustSql()

	if err := s.db.QueryRowContext(ctx, query, args...).Scan(
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
	query, args = builder.
		Select(
			"c.id",
			"c.title",
			"c.image",
			"oi.topping",
		).
		From("api.order_items AS oi").
		Join("api.coffee AS c ON oi.coffee_id = c.id").
		Where(sq.Eq{"order_id": req.OrderID}).
		MustSql()

	rows, err := s.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, errors.Wrap(err, "get order items")
	}

	defer func() {
		if errClose := rows.Close(); errClose != nil {
			logger.Errorf("Close rows: %s", errClose)
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

	if err := rows.Err(); err != nil {
		return nil, errors.Wrap(err, "rows error")
	}

	return &model.GetOrderInfoResStorage{Order: order}, nil
}

func (s *storage) CheckCoffeeIDsExists(ctx context.Context, req *model.CheckCoffeeIDsExistsReqStorage) error {
	ctx, span := tracer.Start(ctx, "CheckCoffeeIDsExists")
	defer span.End()

	var count int

	query, args := builder.
		Select("count(*)").
		From("api.coffee").
		Where(sq.Expr("id = ANY(?)", pq.Array(req.CoffeeIDs))).
		MustSql()

	if err := s.db.QueryRowContext(ctx, query, args...).Scan(&count); err != nil {
		return errors.Wrap(err, "scan coffee IDs count")
	}

	if count != len(req.CoffeeIDs) {
		return errors.Wrap(errors.CoffeeNotExists, "requested coffee ID not found")
	}

	return nil
}

func (s *storage) CheckToppingsExists(ctx context.Context, req *model.CheckToppingsExistsReqStorage) error {
	ctx, span := tracer.Start(ctx, "CheckToppingsExists")
	defer span.End()

	var count int

	query, args := builder.
		Select("count(*)").
		From("unnest(enum_range(NULL::api.topping)) AS t(name)").
		Where(sq.Expr("t.name::TEXT = ANY(?)", pq.Array(req.Toppings))).
		MustSql()

	if err := s.db.QueryRowContext(ctx, query, args...).Scan(&count); err != nil {
		return errors.Wrap(err, "scan toppings count")
	}

	if count != len(req.Toppings) {
		return errors.Wrap(errors.ToppingNotExists, "requested topping not found")
	}

	return nil
}

func (s *storage) CreateOrder(ctx context.Context, req *model.CreateOrderReqStorage) (_ *model.CreateOrderResStorage, err error) {
	ctx, span := tracer.Start(ctx, "CreateOrder")
	defer span.End()

	tx, err := s.db.BeginTx(ctx, nil)
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

	query, args := builder.
		Insert("api.orders").
		Columns(
			"user_id",
			"address",
		).
		Values(
			req.UserID,
			req.Address,
		).
		Suffix("RETURNING id").
		MustSql()

	if err := tx.QueryRowContext(ctx, query, args...).Scan(&orderID); err != nil {
		return nil, errors.Wrap(err, "insert order")
	}

	// Insert items
	query, _ = builder.
		Insert("api.order_items").
		Columns(
			"order_id",
			"coffee_id",
			"topping",
		).
		Values("?", "?", "?").
		MustSql()

	stmt, err := tx.PrepareContext(ctx, query)
	if err != nil {
		return nil, errors.Wrap(err, "prepare items")
	}

	defer func() {
		if errClose := stmt.Close(); errClose != nil {
			logger.Errorf("Close stmt: %s", errClose)
		}
	}()

	for i := range req.Items {
		if _, err = stmt.ExecContext(
			ctx,
			orderID,
			req.Items[i].CoffeeID,
			&req.Items[i].Topping,
		); err != nil {
			return nil, errors.Wrap(err, "insert item")
		}
	}

	return &model.CreateOrderResStorage{OrderID: orderID}, nil
}

func (s *storage) CancelOrder(ctx context.Context, req *model.CancelOrderReqStorage) (*model.CancelOrderResStorage, error) {
	ctx, span := tracer.Start(ctx, "CancelOrder")
	defer span.End()

	var order model.CancelOrderResStorage

	query, args := builder.
		Select(
			"id",
			"user_id",
			"created_at",
			"status",
		).
		From("api.orders").
		Where(sq.Eq{"id": req.OrderID}).
		MustSql()

	if err := s.db.QueryRowContext(ctx, query, args...).Scan(
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

	// Here may be status race condition (if status setted
	// as cooking now), but its acceptable

	query, args = builder.
		Update("api.orders").
		Set("status", "cancelled").
		Where(sq.Eq{"id": req.OrderID}).
		Suffix("RETURNING status").
		MustSql()

	if err := s.db.QueryRowContext(ctx, query, args...).Scan(&order.OrderStatus); err != nil {
		return nil, errors.Wrap(err, "change order status")
	}

	return &order, nil
}

func (s *storage) EmployeeCompleteOrder(ctx context.Context, req *model.EmployeeCompleteOrderReqStorage) (*model.EmployeeCompleteOrderResStorage, error) {
	ctx, span := tracer.Start(ctx, "EmployeeCompleteOrder")
	defer span.End()

	var order model.EmployeeCompleteOrderResStorage

	query, args := builder.
		Update("api.orders").
		Set("status", "ready to receive").
		Where(sq.Eq{"id": req.OrderID}).
		Suffix(`RETURNING 
			id,
			user_id,
			created_at,
			status
		`).
		MustSql()

	if err := s.db.QueryRowContext(ctx, query, args...).Scan(
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
