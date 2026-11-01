package cache

import (
	"context"
	"fmt"
	"time"

	"coffeeshop-api/internal/services/users/model"
	"coffeeshop-api/pkg/errors"
	"coffeeshop-api/pkg/tracer"

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
func (c *cache) SaveUserRefreshToken(ctx context.Context, userID uint64, token model.Token) error {
	ctx, span := tracer.Start(ctx, "SaveUserRefreshToken")
	defer span.End()

	if err := c.redis.Set(
		ctx,
		fmt.Sprintf(refreshTokenFmt, userID, token.String),
		nil,
		time.Until(token.Exp),
	).Err(); err != nil {
		return errors.Wrap(err, "save token")
	}

	return nil
}

// RevokeUserRefreshToken removes refresh token by user ID.
func (c *cache) RevokeUserRefreshToken(ctx context.Context, userID uint64, token string) error {
	ctx, span := tracer.Start(ctx, "RevokeUserRefreshToken")
	defer span.End()

	if err := c.redis.GetDel(
		ctx,
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
func (c *cache) RevokeAllUserRefreshTokens(ctx context.Context, userID uint64) ([]string, error) {
	ctx, span := tracer.Start(ctx, "RevokeAllUserRefreshTokens")
	defer span.End()

	var tokens []string

	if err := c.redis.Keys(
		ctx,
		fmt.Sprintf(refreshTokenFmt, userID, "*"),
	).ScanSlice(&tokens); err != nil {
		return nil, errors.Wrap(err, "get all tokens")
	}

	if len(tokens) == 0 {
		return nil, nil
	}

	if err := c.redis.Del(
		ctx,
		tokens...,
	).Err(); err != nil {
		return nil, errors.Wrap(err, "revoke all tokens")
	}

	return tokens, nil
}
