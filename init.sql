-- CREATE SCHEMA

CREATE SCHEMA IF NOT EXISTS api;

-- DROP TABLES

-- DROP TABLE IF EXISTS api.users;
-- DROP TABLE IF EXISTS api.coffee;
-- DROP TABLE IF EXISTS api.orders;
-- DROP TABLE IF EXISTS api.order_items;

-- CREATE TABLES

CREATE TABLE IF NOT EXISTS api.users (
    id              SERIAL PRIMARY KEY,
    username        VARCHAR(200) NOT NULL,
    email           VARCHAR(200) UNIQUE NOT NULL,
    phone           BIGINT NOT NULL,
    password_hash   BYTEA NOT NULL,
    deleted_at      TIMESTAMP WITHOUT TIME ZONE DEFAULT NULL,
    created_at      TIMESTAMP WITHOUT TIME ZONE DEFAULT current_timestamp NOT NULL
);

CREATE TABLE IF NOT EXISTS api.coffee (
    id              SERIAL PRIMARY KEY,
    title           VARCHAR(200) NOT NULL,
    "description"   TEXT,
    "image"         VARCHAR(1000) DEFAULT 'https://example.com/coffee.png' NOT NULL,
    "weight"        INTEGER NOT NULL,
    price           DECIMAL(10, 2) NOT NULL DEFAULT 0,
    created_at      TIMESTAMP WITHOUT TIME ZONE DEFAULT current_timestamp NOT NULL
);
COMMENT ON COLUMN api.coffee.weight IS 'Net weight in grams';

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

-- FILL TABLES

INSERT INTO api.users (
    username,
    email,
    phone,
    password_hash
) VALUES (
    'johndoe',
    'test@user.com',
    1234567890,
    E'\\xbcea32d00fc2350539eb56f2931b32371e895c7189da287743be71bdd22f8bab' -- "Password:123"
);

INSERT INTO api.coffee (
    title, "description", "weight", price
) VALUES
    ('Espresso', 'A strong and concentrated coffee made by forcing steam through finely-ground coffee beans.', 250, 4.99),
    ('Latte', 'A creamy and mild coffee drink made with espresso and steamed milk.', 400, 5.99),
    ('Cappuccino', 'An Italian coffee drink made with equal parts of espresso, steamed milk, and milk foam.', 350, 6.49),
    ('Cold Brew', 'A smooth and refreshing coffee drink made by steeping coarse coffee grounds in cold water for an extended period of time.', 500, 4.99),
    ('Mocha', 'A decadent coffee drink made with espresso, steamed milk, chocolate syrup, and whipped cream.', 450, 5.49);