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
    -d '{
            "email":    "test@mail.com",
            "password": "Password:123"
    }'
```

### Get account information
```bash
curl -X GET http://localhost:1337/api/user \
    -H 'Accept: application/json'
```

### Change password
```bash
curl -X PUT http://localhost:1337/api/user/password \
    -H 'Accept: application/json' \
    -d '{
            "old_password": "Password:123",
            "new_password": "newPassword:123"
    }'
```

### Refresh token
```bash
curl -X POST http://localhost:1337/api/auth/refresh \
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


### Create a coffee order (payment not implemented thus not needed)
```bash
curl -X POST http://localhost:1337/api/orders \
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
### Connect to Server-Sent Events endpoint for get norifications about orders statuses updates
```bash
curl -N http://localhost:1337/api/events/orders/statuses \
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
curl -X POST http://localhost:1337/api/orders/2/cancel \
    -H 'Accept: application/json'
```

### List your orders
```bash
curl -X GET http://localhost:1337/api/orders \
    -H 'Accept: application/json'
```

### Get order details
```bash
curl -X GET http://localhost:1337/api/orders/1 \
    -H 'Accept: application/json'
```


# DEAUTH


### Delete your account (revoke access tokens not implemented)
```bash
curl -X DELETE http://localhost:1337/api/user \
    -H 'Accept: application/json'
```

### Singnout from this session
```bash
curl -X POST http://localhost:1337/api/auth/signout \
    -H 'Accept: application/json'
```