# AUTHORIZATION


### Create new account
```bash
curl -X POST http://localhost:1337/api/auth/signup \
    -H 'Accept: application/json' \
    -d '{
        "username": "test",
        "phone":    12345678901,
        "email":    "test@mail.com",
        "password": "Password:123"
    }'
```

### Login
```bash
curl -X POST http://localhost:1337/api/auth/signin \
    -H 'Accept: application/json' \
    -c cookies.txt \
    -d '{
            "email":    "test@mail.com",
            "password": "Password:123"
    }'
```

### Get account information
```bash
curl -X GET http://localhost:1337/api/user \
    -H 'Accept: application/json' \
    -b cookies.txt
```

### Change password
```bash
curl -X PUT http://localhost:1337/api/user/password \
    -H 'Accept: application/json' \
    -b cookies.txt \
    -d '{
            "old_password": "Password:123",
            "new_password": "newPassword:123"
    }'
```

### Refresh token
```bash
curl -X POST http://localhost:1337/api/auth/refresh \
    -b cookies.txt \
    -H 'Accept: application/json'
```


# COFFEE


### List available coffees
```bash
curl -X GET http://localhost:1337/api/coffees \
    -H 'Accept: application/json'
```

### Get coffee details
```bash
curl -X GET http://localhost:1337/api/coffees/2 \
    -H 'Accept: application/json'
```

### List available toppings
```bash
curl -X GET http://localhost:1337/api/toppings \
    -H 'Accept: application/json'
```


# ORDERS


### Create a coffee order (payment not implemented)
```bash
curl -X POST http://localhost:1337/api/orders \
    -b cookies.txt \
    -H 'Accept: application/json' \
    -d '{
            "address": "Grove st., 45",
            "items": [
                {
                    "coffee_id": 2,
                    "topping": "vanilla"
                }
            ]
    }'
```

### Connect to Server-Sent Events endpoint for get orders statuses notifications
```bash
curl -N http://localhost:1337/api/events/orders/statuses \
    -b cookies.txt \
    -H 'Accept: text/event-stream'
```

### Complete order cooking (as a coffeeshop employee)
```bash
curl -X POST http://localhost:1337/api/employee/orders/1/complete \
    -H 'Authorization: Bearer secret-employee-token' \
    -H 'Accept: application/json'
```

### ...or cancel the order (as a user)
```bash
curl -X POST http://localhost:1337/api/orders/1/cancel \
    -b cookies.txt \
    -H 'Accept: application/json'
```

### List your orders
```bash
curl -X GET http://localhost:1337/api/orders \
    -b cookies.txt \
    -H 'Accept: application/json'
```

### Get order details
```bash
curl -X GET http://localhost:1337/api/orders/1 \
    -b cookies.txt \
    -H 'Accept: application/json'
```


# DEAUTH


### Delete your account
```bash
curl -X DELETE http://localhost:1337/api/user \
    -b cookies.txt \
    -H 'Accept: application/json'
```

### Signout from active session
```bash
curl -X POST http://localhost:1337/api/auth/signout \
    -b cookies.txt \
    -H 'Accept: application/json'
```