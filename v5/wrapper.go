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

var _ ocredis.Cmdable = &Wrapper{}

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
