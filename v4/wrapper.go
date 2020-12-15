package v4

import (
	"context"
	"time"

	"github.com/KolbyMcGarrah/ocredis"
	pkgredis "gopkg.in/redis.v4"
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

// Del integrates the redis Del command with metrics
func (w *Wrapper) Del(ctx context.Context, keys ...string) (cmd ocredis.IntCmd) {
	var recordCallFunc = ocredis.RecordCall(ctx, "go.redis.del", w.instanceName)
	defer func() {
		recordCallFunc(cmd)
	}()
	cmd = w.client.Del(keys...)
	return
}

// SetNX integrates the redis SetNX command with metrics
func (w *Wrapper) SetNX(ctx context.Context, key string, value interface{}, expiration time.Duration) (cmd ocredis.BoolCmd) {
	var recordCallFunc = ocredis.RecordCall(ctx, "go.redis.setnx", w.instanceName)
	defer func() {
		recordCallFunc(cmd)
	}()
	cmd = w.client.SetNX(key, value, expiration)
	return
}

// Close integrates the redis Close command with metrics
func (w *Wrapper) Close(ctx context.Context) (err error) {
	var recordCallFunc = ocredis.RecordCall(ctx, "go.redis.close", w.instanceName)
	defer func() {
		// Pass in a blank cmd because there is no command type returned from close
		recordCallFunc(&pkgredis.Cmd{})
	}()
	err = w.client.Close()
	return
}

// Expire integrates the redis Expire command with metrics
func (w *Wrapper) Expire(ctx context.Context, key string, expiration time.Duration) (cmd ocredis.BoolCmd) {
	var recordCallFunc = ocredis.RecordCall(ctx, "go.redis.expire", w.instanceName)
	defer func() {
		recordCallFunc(cmd)
	}()
	cmd = w.client.Expire(key, expiration)
	return
}
