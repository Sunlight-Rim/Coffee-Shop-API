-- +goose Up
CREATE TYPE api.topping AS ENUM (
    'vanilla',
    'strawberry',
    'mulberry'
);

CREATE TABLE IF NOT EXISTS api.order_items (
    id              SERIAL PRIMARY KEY,
    order_id        INTEGER NOT NULL REFERENCES api.orders(id),
    coffee_id       INTEGER NOT NULL REFERENCES api.coffee(id),
    topping         api.topping,
    created_at      TIMESTAMP WITHOUT TIME ZONE DEFAULT current_timestamp NOT NULL
);

-- +goose Down
DROP TABLE api.order_items;
DROP TYPE api.topping;
