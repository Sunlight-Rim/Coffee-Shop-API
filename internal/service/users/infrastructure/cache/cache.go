package cache

import (
	"context"
	"fmt"
	"time"

	"coffeeshop-api/internal/service/users/model"
	"coffeeshop-api/pkg/errors"

	"github.com/redis/go-redis/v9"
)

const refreshTokenFmt = "refresh_%d_%s"

type cache struct {
	redis *redis.Client
}

func New(redis *redis.Client) *cache {
	return &cache{redis: redis}
}

// SaveUserRefreshToken saves refresh token by user ID with expiration.
func (c *cache) SaveUserRefreshToken(userID uint64, token *model.Token) error {
	if err := c.redis.Set(
		context.TODO(),
		fmt.Sprintf(refreshTokenFmt, userID, token.Token),
		nil,
		time.Until(token.Exp),
	).Err(); err != nil {
		return errors.Wrap(err, "store token")
	}

	return nil
}

// RevokeUserRefreshToken removes refresh token by user ID.
func (c *cache) RevokeUserRefreshToken(userID uint64, token string) error {
	if err := c.redis.GetDel(
		context.TODO(),
		fmt.Sprintf(refreshTokenFmt, userID, token),
	).Err(); err != nil {
		if errors.Is(err, redis.Nil) {
			return errors.Wrap(errors.InvalidToken, "token not found")
		}

		return errors.Wrap(err, "revoke token")
	}

	return nil
}

// RevokeAllUserRefreshTokens removes all refresh tokens by user ID.
func (c *cache) RevokeAllUserRefreshTokens(userID uint64) ([]string, error) {
	var tokens []string

	if err := c.redis.Keys(
		context.TODO(),
		fmt.Sprintf(refreshTokenFmt, userID, "*"),
	).ScanSlice(&tokens); err != nil {
		return nil, errors.Wrap(err, "get all tokens")
	}

	if err := c.redis.Del(
		context.TODO(),
		tokens...,
	).Err(); err != nil {
		return nil, errors.Wrap(err, "revoke all tokens")
	}

	return tokens, nil
}
