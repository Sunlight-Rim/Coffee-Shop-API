package model

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	jwt.RegisteredClaims

	UserID uint64 `json:"user_id"`
}

type Token struct {
	String string
	Exp    time.Time
}
