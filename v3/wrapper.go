package v3

import (
	"context"
	"time"

	"github.com/KolbyMcGarrah/ocredis"
	"go.opencensus.io/trace"
	pkgredis "gopkg.in/redis.v3"
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

// Get integrates the redis get command with metrics
func (w *Wrapper) Get(ctx context.Context, key string) (cmd ocredis.StringCmd) {
	if ocredis.AllowTrace(ctx, w.options.Get, w.options.AllowRoot) {
		span := ocredis.StartSpan(ctx, "Get", w.options)
		if span != nil {
			defer span.EndSpanWithErr(cmd.Err())
		}
	}
	var recordCallFunc = ocredis.RecordCall(ctx, "go.redis.get", w.options.InstanceName)
	defer func() {
		recordCallFunc(cmd)
	}()

	cmd = w.client.Get(key)
	return
}

// Set integrates the redis Set command with metrics
func (w *Wrapper) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) (cmd ocredis.StatusCmd) {
	if ocredis.AllowTrace(ctx, w.options.Set, w.options.AllowRoot) {
		span := ocredis.StartSpan(ctx, "Set", w.options)
		if span != nil {
			defer span.EndSpanWithErr(cmd.Err())
		}
	}
	var recordCallFunc = ocredis.RecordCall(ctx, "go.redis.set", w.options.InstanceName)
	defer func() {
		recordCallFunc(cmd)
	}()

	cmd = w.client.Set(key, value, expiration)
	return
}

// Incr integrates the redis Incr command with metrics
func (w *Wrapper) Incr(ctx context.Context, key string) (cmd ocredis.IntCmd) {
	if ocredis.AllowTrace(ctx, w.options.Incr, w.options.AllowRoot) {
		span := ocredis.StartSpan(ctx, "Incr", w.options)
		if span != nil {
			defer span.EndSpanWithErr(cmd.Err())
		}
	}
	var recordCallFunc = ocredis.RecordCall(ctx, "go.redis.incr", w.options.InstanceName)
	defer func() {
		recordCallFunc(cmd)
	}()
	cmd = w.client.Incr(key)
	return

}

// Ping integrates the redis Ping command with metrics
func (w *Wrapper) Ping(ctx context.Context) (cmd ocredis.StatusCmd) {
	if ocredis.AllowTrace(ctx, w.options.Ping, w.options.AllowRoot) {
		span := ocredis.StartSpan(ctx, "Ping", w.options)
		if span != nil {
			defer span.EndSpanWithErr(cmd.Err())
		}
	}
	var recordCallFunc = ocredis.RecordCall(ctx, "go.redis.ping", w.options.InstanceName)
	defer func() {
		recordCallFunc(cmd)
	}()
	cmd = w.client.Ping()
	return cmd
}
