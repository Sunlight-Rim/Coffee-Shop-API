package storage

import (
	"context"
	"database/sql"

	"coffeeshop-api/internal/services/coffee/model"
	"coffeeshop-api/pkg/errors"
	"coffeeshop-api/pkg/tracer"

	sq "github.com/Masterminds/squirrel"
	logger "github.com/sirupsen/logrus"
)

var builder = sq.StatementBuilder.PlaceholderFormat(sq.Dollar)

const (
	coffeeLimit   = 10
	toppingsLimit = 10
)

type storage struct {
	db *sql.DB
}

func New(db *sql.DB) *storage {
	return &storage{db: db}
}

func (s *storage) ListCoffees(ctx context.Context, req *model.ListCoffeesReqStorage) (*model.ListCoffeesResStorage, error) {
	ctx, span := tracer.Start(ctx, "ListCoffees")
	defer span.End()

	query, args := builder.
		Select(
			"id",
			"title",
			"description",
			"image",
			"weight",
			"price",
		).
		From("api.coffee").
		OrderBy("id ASC").
		Limit(coffeeLimit).
		Offset(req.Offset).
		MustSql()

	rows, err := s.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, errors.Wrap(err, "get coffee list")
	}

	defer func() {
		if errClose := rows.Close(); errClose != nil {
			logger.Errorf("Close rows: %s", errClose)
		}
	}()

	var (
		coffeeList []model.Coffee
		coffee     model.Coffee
	)

	for rows.Next() {
		if err := rows.Scan(
			&coffee.ID,
			&coffee.Title,
			&coffee.Description,
			&coffee.Image,
			&coffee.Weight,
			&coffee.Price,
		); err != nil {
			return nil, errors.Wrap(err, "scan coffee list")
		}

		coffeeList = append(coffeeList, coffee)
	}

	if err := rows.Err(); err != nil {
		return nil, errors.Wrap(err, "rows error")
	}

	return &model.ListCoffeesResStorage{CoffeeList: coffeeList}, nil
}

func (s *storage) GetCoffeeInfo(ctx context.Context, req *model.GetCoffeeInfoReqStorage) (*model.GetCoffeeInfoResStorage, error) {
	ctx, span := tracer.Start(ctx, "GetCoffeeInfo")
	defer span.End()

	var coffee model.Coffee

	query, args := builder.
		Select(
			"id",
			"title",
			"description",
			"image",
			"weight",
			"price",
		).
		From("api.coffee").
		Where(sq.Eq{"id": req.CoffeeID}).
		MustSql()

	if err := s.db.QueryRowContext(ctx, query, args...).Scan(
		&coffee.ID,
		&coffee.Title,
		&coffee.Description,
		&coffee.Image,
		&coffee.Weight,
		&coffee.Price,
	); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.Wrap(errors.CoffeeNotExists, "coffee not found")
		}

		return nil, errors.Wrap(err, "get coffee")
	}

	return &model.GetCoffeeInfoResStorage{Coffee: &coffee}, nil
}

func (s *storage) ListToppings(ctx context.Context, req *model.ListToppingsReqStorage) (*model.ListToppingsResStorage, error) {
	ctx, span := tracer.Start(ctx, "ListToppings")
	defer span.End()

	query, args := builder.
		Select("unnest(enum_range(NULL::api.topping))").
		Limit(toppingsLimit).
		Offset(req.Offset).
		MustSql()

	rows, err := s.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, errors.Wrap(err, "get toppings list")
	}

	defer func() {
		if errClose := rows.Close(); errClose != nil {
			logger.Errorf("Close rows: %s", errClose)
		}
	}()

	var (
		toppingsList []string
		topping      string
	)

	for rows.Next() {
		if err := rows.Scan(&topping); err != nil {
			return nil, errors.Wrap(err, "scan toppings list")
		}

		toppingsList = append(toppingsList, topping)
	}

	if err := rows.Err(); err != nil {
		return nil, errors.Wrap(err, "rows error")
	}

	return &model.ListToppingsResStorage{ToppingsList: toppingsList}, nil
}
