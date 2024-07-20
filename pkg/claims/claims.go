package claims

import "github.com/golang-jwt/jwt/v5"

// General tokens claims.
type Claims struct {
	jwt.RegisteredClaims

	UserID uint64 `json:"user_id"`
}
