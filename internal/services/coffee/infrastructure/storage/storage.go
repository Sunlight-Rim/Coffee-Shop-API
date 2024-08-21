package storage

import (
	"context"
	"database/sql"

	"coffeeshop-api/internal/services/coffee/model"
	"coffeeshop-api/pkg/errors"

	logger "github.com/sirupsen/logrus"
)

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
	rows, err := s.db.QueryContext(ctx, `
		SELECT
			id,
			title,
			"description",
			"image",
			"weight",
			price
		FROM api.coffee
		ORDER BY id ASC
		LIMIT $1 OFFSET $2
	`,
		coffeeLimit,
		req.Offset,
	)
	if err != nil {
		return nil, errors.Wrap(err, "get coffee list")
	}

	defer func() {
		if errClose := rows.Close(); errClose != nil {
			logger.WithField("error", errClose).Error("Close rows")
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
	var coffee model.Coffee

	if err := s.db.QueryRowContext(ctx, `
		SELECT
			id,
			title,
			"description",
			"image",
			"weight",
			price
		FROM api.coffee
		WHERE id = $1
	`,
		req.CoffeeID,
	).Scan(
		&coffee.ID,
		&coffee.Title,
		&coffee.Description,
		&coffee.Image,
		&coffee.Weight,
		&coffee.Price,
	); err != nil {
		return nil, errors.Wrap(err, "get coffee")
	}

	return &model.GetCoffeeInfoResStorage{Coffee: &coffee}, nil
}

func (s *storage) ListToppings(ctx context.Context, req *model.ListToppingsReqStorage) (*model.ListToppingsResStorage, error) {
	rows, err := s.db.QueryContext(ctx, `
		SELECT unnest(enum_range(NULL::api.topping))
		LIMIT $1 OFFSET $2
	`,
		toppingsLimit,
		req.Offset,
	)
	if err != nil {
		return nil, errors.Wrap(err, "get toppings list")
	}

	defer func() {
		if errClose := rows.Close(); errClose != nil {
			logger.WithField("error", errClose).Error("Close rows")
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
