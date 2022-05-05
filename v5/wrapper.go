package v5

import (
	"context"
	"time"

	"github.com/KolbyMcGarrah/ocredis"
	"go.opencensus.io/trace"
	pkgredis "gopkg.in/redis.v5"
)

// Wrap returns a wrapped redis client
func Wrap(c *pkgredis.Client, options ...ocredis.TraceOption) *Wrapper {
	o := ocredis.TraceOptions{}
	for _, option := range options {
		option(&o)
	}
	if o.InstanceName == "" {
		o.InstanceName = ocredis.DefaultInstanceName
	} else {
		o.DefaultAttributes = append(o.DefaultAttributes, trace.StringAttribute("cache.instance", o.InstanceName))
	}
	return &Wrapper{
		client:  c,
		options: o,
	}
}

var _ ocredis.Client = &Wrapper{}

// Wrapper wraps the redis package with an instance name to be used to collect metrics.
type Wrapper struct {
	client  *pkgredis.Client
	options ocredis.TraceOptions
}

func (w *Wrapper) ExpireAt(ctx context.Context, key string, tm time.Time) (cmd ocredis.BoolCmd) {
	if ocredis.AllowTrace(ctx, w.options.ExpireAt, w.options.AllowRoot) {
		span := ocredis.StartSpan(ctx, "ExpireAt", w.options)
		if span != nil {
			defer func() {
				span.EndSpanWithErr(cmd.Err())
			}()
		}
	}
	var recordCallFunc = ocredis.RecordCall(ctx, "go.redis.expireat", w.options.InstanceName)
	defer func() {
		recordCallFunc(cmd)
	}()
	cmd = w.client.ExpireAt(key, tm)
	return cmd
}

func (w *Wrapper) HLen(ctx context.Context, key string) (cmd ocredis.IntCmd) {
	if ocredis.AllowTrace(ctx, w.options.HLen, w.options.AllowRoot) {
		span := ocredis.StartSpan(ctx, "HLen", w.options)
		if span != nil {
			defer func() {
				span.EndSpanWithErr(cmd.Err())
			}()
		}
	}
	var recordCallFunc = ocredis.RecordCall(ctx, "go.redis.hlen", w.options.InstanceName)
	defer func() {
		recordCallFunc(cmd)
	}()
	cmd = w.client.HLen(key)
	return
}

func (w *Wrapper) HGet(ctx context.Context, key, field string) (cmd ocredis.StringCmd) {
	if ocredis.AllowTrace(ctx, w.options.HGet, w.options.AllowRoot) {
		span := ocredis.StartSpan(ctx, "HGet", w.options)
		if span != nil {
			defer func() {
				span.EndSpanWithErr(cmd.Err())
			}()
		}
	}
	var recordCallFunc = ocredis.RecordCall(ctx, "go.redis.hget", w.options.InstanceName)
	defer func() {
		recordCallFunc(cmd)
	}()
	cmd = w.client.HGet(key, field)
	return
}

func (w *Wrapper) HSet(ctx context.Context, key, field string, value interface{}) (cmd ocredis.BoolCmd) {
	if ocredis.AllowTrace(ctx, w.options.HSet, w.options.AllowRoot) {
		span := ocredis.StartSpan(ctx, "Hset", w.options)
		if span != nil {
			defer func() {
				span.EndSpanWithErr(cmd.Err())
			}()
		}
	}
	var recordCallFunc = ocredis.RecordCall(ctx, "go.redis.hset", w.options.InstanceName)
	defer func() {
		recordCallFunc(cmd)
	}()
	cmd = w.client.HSet(key, field, value)
	return
}

func (w *Wrapper) SetNX(ctx context.Context, key string, value interface{}, expiration time.Duration) (cmd ocredis.BoolCmd) {
	if ocredis.AllowTrace(ctx, w.options.SetNX, w.options.AllowRoot) {
		span := ocredis.StartSpan(ctx, "SetNX", w.options)
		if span != nil {
			defer func() {
				span.EndSpanWithErr(cmd.Err())
			}()
		}
	}
	var recordCallFunc = ocredis.RecordCall(ctx, "go.redis.setnx", w.options.InstanceName)
	defer func() {
		recordCallFunc(cmd)
	}()
	cmd = w.client.SetNX(key, value, expiration)
	return
}

func (w *Wrapper) Del(ctx context.Context, keys ...string) (cmd ocredis.IntCmd) {
	if ocredis.AllowTrace(ctx, w.options.Del, w.options.AllowRoot) {
		span := ocredis.StartSpan(ctx, "Del", w.options)
		if span != nil {
			defer func() {
				span.EndSpanWithErr(cmd.Err())
			}()
		}
	}
	var recordCallFunc = ocredis.RecordCall(ctx, "go.redis.del", w.options.InstanceName)
	defer func() {
		recordCallFunc(cmd)
	}()
	cmd = w.client.Del(keys...)
	return
}

func (w *Wrapper) Get(ctx context.Context, key string) (cmd ocredis.StringCmd) {
	if ocredis.AllowTrace(ctx, w.options.Get, w.options.AllowRoot) {
		span := ocredis.StartSpan(ctx, "Get", w.options)
		if span != nil {
			defer func() {
				span.EndSpanWithErr(cmd.Err())
			}()
		}
	}
	var recordCallFunc = ocredis.RecordCall(ctx, "go.redis.get", w.options.InstanceName)
	defer func() {
		recordCallFunc(cmd)
	}()
	cmd = w.client.Get(key)
	return
}

func (w *Wrapper) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) (cmd ocredis.StatusCmd) {
	if ocredis.AllowTrace(ctx, w.options.Set, w.options.AllowRoot) {
		span := ocredis.StartSpan(ctx, "Set", w.options)
		if span != nil {
			defer func() {
				span.EndSpanWithErr(cmd.Err())
			}()
		}
	}
	var recordCallFunc = ocredis.RecordCall(ctx, "go.redis.set", w.options.InstanceName)
	defer func() {
		recordCallFunc(cmd)
	}()
	cmd = w.client.Set(key, value, expiration)
	return
}

func (w *Wrapper) LPush(ctx context.Context, key string, values ...interface{}) (cmd ocredis.IntCmd) {
	if ocredis.AllowTrace(ctx, w.options.LPush, w.options.AllowRoot) {
		span := ocredis.StartSpan(ctx, "LPush", w.options)
		if span != nil {
			defer func() {
				span.EndSpanWithErr(cmd.Err())
			}()
		}
	}
	var recordCallFunc = ocredis.RecordCall(ctx, "go.redis.lpush", w.options.InstanceName)
	defer func() {
		recordCallFunc(cmd)
	}()
	cmd = w.client.LPush(key, values...)
	return
}
func (w *Wrapper) RPush(ctx context.Context, key string, values ...interface{}) (cmd ocredis.IntCmd) {
	if ocredis.AllowTrace(ctx, w.options.RPush, w.options.AllowRoot) {
		span := ocredis.StartSpan(ctx, "RPush", w.options)
		if span != nil {
			defer func() {
				span.EndSpanWithErr(cmd.Err())
			}()
		}
	}
	var recordCallFunc = ocredis.RecordCall(ctx, "go.redis.rpush", w.options.InstanceName)
	defer func() {
		recordCallFunc(cmd)
	}()
	cmd = w.client.RPush(key, values...)
	return
}
func (w *Wrapper) RPop(ctx context.Context, key string) (cmd ocredis.StringCmd) {
	if ocredis.AllowTrace(ctx, w.options.RPop, w.options.AllowRoot) {
		span := ocredis.StartSpan(ctx, "RPop", w.options)
		if span != nil {
			defer func() {
				span.EndSpanWithErr(cmd.Err())
			}()
		}
	}
	var recordCallFunc = ocredis.RecordCall(ctx, "go.redis.rpop", w.options.InstanceName)
	defer func() {
		recordCallFunc(cmd)
	}()
	cmd = w.client.RPop(key)
	return
}

// Close integrates the redis Close command with metrics
func (w *Wrapper) Close(ctx context.Context) (err error) {
	if ocredis.AllowTrace(ctx, w.options.Close, w.options.AllowRoot) {
		span := ocredis.StartSpan(ctx, "Close", w.options)
		if span != nil {
			defer func() {
				span.EndSpanWithErr(err)
			}()
		}
	}
	var recordCallFunc = ocredis.RecordCall(ctx, "go.redis.close", w.options.InstanceName)
	defer func() {
		// Pass in a blank cmd because there is no command type returned from close
		recordCallFunc(&pkgredis.Cmd{})
	}()
	err = w.client.Close()
	return
}

// Eval integrates the redis Eval command with metrics
func (w *Wrapper) Eval(ctx context.Context, script string, keys []string, args []string) (cmd ocredis.RedisCmd) {
	if ocredis.AllowTrace(ctx, w.options.Eval, w.options.AllowRoot) {
		span := ocredis.StartSpan(ctx, "go.redis.eval", w.options)
		if span != nil {
			defer func() {
				span.EndSpanWithErr(cmd.Err())
			}()
		}
	}
	var recordCallFunc = ocredis.RecordCall(ctx, "go.redis.eval", w.options.InstanceName)
	defer func() {
		recordCallFunc(cmd)
	}()
	cmd = w.client.Eval(script, keys, args)
	return
}

// Incr integrates the redis Incr command with metrics
func (w *Wrapper) Incr(ctx context.Context, key string) (cmd ocredis.IntCmd) {
	if ocredis.AllowTrace(ctx, w.options.Incr, w.options.AllowRoot) {
		span := ocredis.StartSpan(ctx, "Incr", w.options)
		if span != nil {
			defer func() {
				span.EndSpanWithErr(cmd.Err())
			}()
		}
	}
	var recordCallFunc = ocredis.RecordCall(ctx, "go.redis.incr", w.options.InstanceName)
	defer func() {
		recordCallFunc(cmd)
	}()
	cmd = w.client.Incr(key)
	return

}

// LPop integrates the redis LPOP command with metrics
func (w *Wrapper) LPop(ctx context.Context, key string) (cmd ocredis.StringCmd) {
	if ocredis.AllowTrace(ctx, w.options.LPop, w.options.AllowRoot) {
		span := ocredis.StartSpan(ctx, "go.redis.lpop", w.options)
		if span != nil {
			defer func() {
				span.EndSpanWithErr(cmd.Err())
			}()
		}
	}
	var recordCallFunc = ocredis.RecordCall(ctx, "go.redis.lpop", w.options.InstanceName)
	defer func() {
		recordCallFunc(cmd)
	}()
	cmd = w.client.LPop(key)
	return
}

// Ping integrates the redis Ping command with metrics
func (w *Wrapper) Ping(ctx context.Context) (cmd ocredis.StatusCmd) {
	if ocredis.AllowTrace(ctx, w.options.Ping, w.options.AllowRoot) {
		span := ocredis.StartSpan(ctx, "Ping", w.options)
		if span != nil {
			defer func() {
				span.EndSpanWithErr(cmd.Err())
			}()
		}
	}
	var recordCallFunc = ocredis.RecordCall(ctx, "go.redis.ping", w.options.InstanceName)
	defer func() {
		recordCallFunc(cmd)
	}()
	cmd = w.client.Ping()
	return cmd
}
