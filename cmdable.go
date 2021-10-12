package ocredis

import (
	"context"
	"time"
)

type Cmdable interface {
	ExpireAt(ctx context.Context, key string, tm time.Time) BoolCmd
	HLen(ctx context.Context, key string) IntCmd
	HGet(ctx context.Context, key, field string) StringCmd
	HSet(ctx context.Context, key, field string, value interface{}) BoolCmd
	SetNX(ctx context.Context, key string, value interface{}, expiration time.Duration) BoolCmd
	Del(ctx context.Context, keys ...string) IntCmd
	Get(ctx context.Context, key string) StringCmd
	Set(ctx context.Context, key string, value interface{}, expiration time.Duration) StatusCmd
}
