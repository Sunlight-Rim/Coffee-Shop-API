-- +goose Up
CREATE TABLE IF NOT EXISTS api.users (
    id              SERIAL PRIMARY KEY,
    username        VARCHAR(200) NOT NULL,
    email           VARCHAR(200) UNIQUE NOT NULL,
    phone           BIGINT NOT NULL,
    password_hash   BYTEA NOT NULL,
    deleted_at      TIMESTAMP WITHOUT TIME ZONE DEFAULT NULL,
    created_at      TIMESTAMP WITHOUT TIME ZONE DEFAULT current_timestamp NOT NULL
);

-- +goose Down
DROP TABLE api.users;