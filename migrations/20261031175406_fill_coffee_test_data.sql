-- +goose Up
INSERT INTO api.coffee (
    title, "description", "weight", price
) VALUES
    ('Espresso', 'A strong and concentrated coffee made by forcing steam through finely-ground coffee beans.', 250, 4.99),
    ('Latte', 'A creamy and mild coffee drink made with espresso and steamed milk.', 400, 5.99),
    ('Cappuccino', 'An Italian coffee drink made with equal parts of espresso, steamed milk, and milk foam.', 350, 6.49),
    ('Cold Brew', 'A smooth and refreshing coffee drink made by steeping coarse coffee grounds in cold water for an extended period of time.', 500, 4.99),
    ('Mocha', 'A decadent coffee drink made with espresso, steamed milk, chocolate syrup, and whipped cream.', 450, 5.49);

-- +goose Down
DELETE FROM api.coffee WHERE title IN ('Espresso', 'Latte', 'Cappuccino', 'Cold Brew', 'Mocha');