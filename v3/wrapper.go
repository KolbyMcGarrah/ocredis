package v3

import (
	"context"
	"time"

	"github.com/KolbyMcGarrah/ocredis"
	pkgredis "gopkg.in/redis.v3"
)

// Wrap returns a wrapped redis client
func Wrap(client *pkgredis.Client, instanceName string) *Wrapper {
	return &Wrapper{
		client:       client,
		instanceName: instanceName,
	}
}

var _ ocredis.Client = &Wrapper{}

// Wrapper wraps the redis package with an instance name to be used to collect metrics.
type Wrapper struct {
	client       *pkgredis.Client
	instanceName string
}

// Get integrates the redis get command with metrics
func (w *Wrapper) Get(ctx context.Context, key string) (cmd ocredis.StringCmd) {
	var recordCallFunc = ocredis.RecordCall(ctx, "go.redis.get", w.instanceName)
	defer func() {
		recordCallFunc(cmd)
	}()

	cmd = w.client.Get(key)
	return
}

// Set integrates the redis Set command with metrics
func (w *Wrapper) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) (cmd ocredis.StatusCmd) {
	var recordCallFunc = ocredis.RecordCall(ctx, "go.redis.set", w.instanceName)
	defer func() {
		recordCallFunc(cmd)
	}()

	cmd = w.client.Set(key, value, expiration)
	return
}

// Incr integrates the redis Incr command with metrics
func (w *Wrapper) Incr(ctx context.Context, key string) (cmd ocredis.IntCmd) {
	var recordCallFunc = ocredis.RecordCall(ctx, "go.redis.incr", w.instanceName)
	defer func() {
		recordCallFunc(cmd)
	}()
	cmd = w.client.Incr(key)
	return

}

// Ping integrates the redis Ping command with metrics
func (w *Wrapper) Ping(ctx context.Context) (cmd ocredis.StatusCmd) {
	var recordCallFunc = ocredis.RecordCall(ctx, "go.redis.ping", w.instanceName)
	defer func() {
		recordCallFunc(cmd)
	}()
	cmd = w.client.Ping()
	return cmd
}
