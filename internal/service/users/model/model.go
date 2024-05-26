package model

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JWTClaims struct {
	jwt.RegisteredClaims

	UserID uint64 `json:"user_id"`
}

type Token struct {
	String string
	Exp    time.Time
}

type User struct {
	ID       uint64
	Username string
	Email    string
	Phone    uint64
}
