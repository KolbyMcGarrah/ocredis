// Package ocredis instruments ocredis interactions with Open Census
package ocredis

import (
	"context"
	"time"
)

// Client represents the redis client that is used throughout each version
type Client interface {
	Get(ctx context.Context, key string) StringCmd
	Set(ctx context.Context, key string, value interface{}, expiration time.Duration) StatusCmd
	Incr(ctx context.Context, key string) IntCmd
	Ping(ctx context.Context) StatusCmd
	Del(ctx context.Context, keys ...string) IntCmd
	SetNX(ctx context.Context, key string, value interface{}, expiration time.Duration) BoolCmd
	Close(ctx context.Context) error
	LPop(ctx context.Context, key string) StringCmd
	Eval(ctx context.Context, script string, keys []string, args []string) RedisCmd
	LPush(ctx context.Context, key string, values ...interface{}) IntCmd
	RPush(ctx context.Context, key string, values ...interface{}) IntCmd
	RPop(ctx context.Context, key string) StringCmd
	HSet(ctx context.Context, key, field string, value interface{}) BoolCmd
	HGet(ctx context.Context, key, field string) StringCmd
	HLen(ctx context.Context, key string) IntCmd
	ExpireAt(ctx context.Context, key string, tm time.Time) BoolCmd
}
