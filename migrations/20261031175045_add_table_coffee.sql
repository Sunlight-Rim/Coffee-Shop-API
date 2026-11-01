-- +goose Up
CREATE TABLE IF NOT EXISTS api.coffee (
    id              SERIAL PRIMARY KEY,
    title           VARCHAR(200) UNIQUE NOT NULL,
    "description"   TEXT,
    "image"         VARCHAR(1000) DEFAULT 'https://example.com/coffee.png' NOT NULL,
    "weight"        INTEGER NOT NULL,
    price           DECIMAL(10, 2) NOT NULL,
    deleted_at      TIMESTAMP WITHOUT TIME ZONE,
    created_at      TIMESTAMP WITHOUT TIME ZONE DEFAULT current_timestamp NOT NULL
);
COMMENT ON COLUMN api.coffee.weight IS 'Net weight in grams';

-- +goose Down
DROP TABLE api.coffee;