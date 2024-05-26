-- CREATE SCHEMA

CREATE SCHEMA IF NOT EXISTS api;

-- DROP TABLES

-- DROP TABLE IF EXISTS api.users;
-- DROP TABLE IF EXISTS api.coffee;
-- DROP TABLE IF EXISTS api.orders;
-- DROP TABLE IF EXISTS api.order_items;

-- CREATE TABLES

-- users

CREATE TABLE IF NOT EXISTS api.users (
    id              SERIAL PRIMARY KEY,
    username        VARCHAR(200) NOT NULL,
    email           VARCHAR(200) UNIQUE NOT NULL,
    phone           BIGINT NOT NULL,
    password_hash   BYTEA NOT NULL,
    deleted_at      TIMESTAMP WITHOUT TIME ZONE DEFAULT NULL,
    created_at      TIMESTAMP WITHOUT TIME ZONE DEFAULT CURRENT_TIMESTAMP NOT NULL
);

-- coffee

CREATE TABLE IF NOT EXISTS api.coffee (
    id              SERIAL PRIMARY KEY,
    title           VARCHAR(200) NOT NULL,
    description     TEXT,
    image           VARCHAR(1000) DEFAULT 'https://cdn-icons-png.freepik.com/128/1047/1047503.png' NOT NULL,
    weight          INTEGER NOT NULL,
    price           DECIMAL(10, 2) NOT NULL DEFAULT 0,
    created_at      TIMESTAMP WITHOUT TIME ZONE DEFAULT CURRENT_TIMESTAMP NOT NULL
);
COMMENT ON COLUMN api.coffee.weight IS 'Net weight in grams';

-- orders

CREATE TYPE status AS ENUM (
    'wait payment',
    'cancelled',
    'queued',
    'wait receiving'
    'received'
);

CREATE TABLE IF NOT EXISTS api.orders (
    id              SERIAL PRIMARY KEY,
    user_id         INTEGER NOT NULL REFERENCES api.users(id),
    address         VARCHAR(1000) NOT NULL,
    paid            BOOLEAN NOT NULL,
    status          status NOT NULL,
    created_at      TIMESTAMP WITHOUT TIME ZONE DEFAULT CURRENT_TIMESTAMP NOT NULL
);

-- order_items

CREATE TYPE topping AS ENUM (
    'vanilla',
    'strawberry',
    'mulberry'
);

CREATE TABLE IF NOT EXISTS api.order_items (
    id              SERIAL PRIMARY KEY,
    order_id        INTEGER NOT NULL REFERENCES api.orders(id),
    coffee_id       INTEGER NOT NULL REFERENCES api.coffee(id),
    topping         topping,
    created_at      TIMESTAMP WITHOUT TIME ZONE DEFAULT CURRENT_TIMESTAMP NOT NULL
);

-- FILL TABLES

INSERT INTO api.users (
    username,
    email,
    phone,
    password_hash
) VALUES (
    'johndoe',
    'me@mail.com',
    82641845273,
    E'\\x5E884898DA28047151D0E56F8DC6292773603D0D6AABBDD62A11EF721D1542D8' -- "password"
);

INSERT INTO api.coffee (
    title, description, weight, price
) VALUES
    ('Espresso', 'A strong and concentrated coffee made by forcing steam through finely-ground coffee beans.', 250, 4.99),
    ('Latte', 'A creamy and mild coffee drink made with espresso and steamed milk.', 400, 5.99),
    ('Cappuccino', 'An Italian coffee drink made with equal parts of espresso, steamed milk, and milk foam.', 350, 6.49),
    ('Cold Brew', 'A smooth and refreshing coffee drink made by steeping coarse coffee grounds in cold water for an extended period of time.', 500, 4.99),
    ('Mocha', 'A decadent coffee drink made with espresso, steamed milk, chocolate syrup, and whipped cream.', 450, 5.49);