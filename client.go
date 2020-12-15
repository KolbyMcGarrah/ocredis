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
}
