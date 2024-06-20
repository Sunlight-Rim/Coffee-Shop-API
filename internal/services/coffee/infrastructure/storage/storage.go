package storage

import (
	"database/sql"

	"coffeeshop-api/internal/services/coffee/model"
	"coffeeshop-api/pkg/errors"
)

const coffeeLimit = 10

type storage struct {
	db *sql.DB
}

func New(db *sql.DB) *storage {
	return &storage{db: db}
}

func (s *storage) GetCoffeeInfo(req *model.GetCoffeeInfoReqStorage) (*model.GetCoffeeInfoResStorage, error) {
	var coffee model.Coffee

	if err := s.db.QueryRow(`
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

func (s *storage) ListCoffee(req *model.ListCoffeeReqStorage) (*model.ListCoffeeResStorage, error) {
	rows, err := s.db.Query(`
		SELECT
			id,
			title,
			"description",
			"image",
			"weight",
			price
		FROM api.coffee
		LIMIT $1 OFFSET $2
	`,
		coffeeLimit,
		req.Offset,
	)
	if err != nil {
		return nil, errors.Wrap(err, "get coffee list")
	}

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

	return &model.ListCoffeeResStorage{CoffeeList: coffeeList}, nil
}
