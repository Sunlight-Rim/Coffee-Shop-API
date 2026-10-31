-- +goose Up
CREATE TYPE api.status AS ENUM (
    'cooking',
    'cancelled',
    'ready to receive',
    'received'
);

CREATE TABLE IF NOT EXISTS api.orders (
    id              SERIAL PRIMARY KEY,
    user_id         INTEGER NOT NULL REFERENCES api.users(id),
    "address"       VARCHAR(1000) NOT NULL,
    "status"        api.status DEFAULT 'cooking' NOT NULL,
    created_at      TIMESTAMP WITHOUT TIME ZONE DEFAULT current_timestamp NOT NULL
);

-- +goose Down
DROP TABLE api.orders;
DROP TYPE api.status;