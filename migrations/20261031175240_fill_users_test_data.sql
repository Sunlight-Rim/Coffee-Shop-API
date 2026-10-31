-- +goose Up
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

-- +goose Down
DELETE FROM api.users WHERE email='test@user.com';
