package ocredis

import "go.opencensus.io/trace"

// DefaultInstanceName is the instance name assigned when one isn't provided
const DefaultInstanceName = "default"

// TraceOption allows for managing cache trace configurations using funcitonal options
type TraceOption func(o *TraceOptions)

// TraceOptions holds configurations of the ocredis tracing middleware
// by default all options are initialized to false.
type TraceOptions struct {

	// AllowRoot, if set to true, will allow ocredis to create root spans in
	// absence of existing spans or even context.
	// Default is to not trace ocredis calls if no existing parent span is found
	// in context or when using methods not taking context.
	AllowRoot bool

	// InstanceName identifies the cache
	InstanceName string

	// DefaultAttributes will set to each span as default
	DefaultAttributes []trace.Attribute

	// Sampler to use when creating spans
	Sampler trace.Sampler

	// Setting the below options will control whether or not spans are created
	// on their call.
	Get    bool
	Set    bool
	Incr   bool
	Ping   bool
	Del    bool
	SetNX  bool
	Close  bool
	Expire bool
}

// WithAllTraceOptions enables all available traceoptions
func WithAllTraceOptions() TraceOption {
	return func(o *TraceOptions) {
		*o = AllTraceOptions
	}
}

// AllTraceOptions has all tracing options enabled
var AllTraceOptions = TraceOptions{
	Get:    true,
	Set:    true,
	Incr:   true,
	Ping:   true,
	Del:    true,
	SetNX:  true,
	Close:  true,
	Expire: true,
}

// WithAllowRoot if set to true, will allow ocredis to create root spans in
// absence of exisiting spans or even context.
// Default is to not trace redis calls if no existing parent span is found
// in context or when using methods not taking context.
func WithAllowRoot(b bool) TraceOption {
	return func(o *TraceOptions) {
		o.AllowRoot = b
	}
}

// WithInstanceName sets cache instance name.
func WithInstanceName(instanceName string) TraceOption {
	return func(o *TraceOptions) {
		o.InstanceName = instanceName
	}
}

// WithGet if true will allow tracing on the get call.
func WithGet(b bool) TraceOption {
	return func(o *TraceOptions) {
		o.Get = b
	}
}

// WithSet if true will allow tracing on the set call.
func WithSet(b bool) TraceOption {
	return func(o *TraceOptions) {
		o.Set = b
	}
}

// WithIncr if true will allow tracing on the incr call.
func WithIncr(b bool) TraceOption {
	return func(o *TraceOptions) {
		o.Incr = b
	}
}

// WithPing if true will allow tracing on the ping call.
func WithPing(b bool) TraceOption {
	return func(o *TraceOptions) {
		o.Ping = b
	}
}

// WithDel if true will allow tracing on the Del call.
func WithDel(b bool) TraceOption {
	return func(o *TraceOptions) {
		o.Del = b
	}
}

// WithSetNX if true will allow tracing on the SetNX call.
func WithSetNX(b bool) TraceOption {
	return func(o *TraceOptions) {
		o.SetNX = b
	}
}

// WithClose if true will allow tracing on the close call.
func WithClose(b bool) TraceOption {
	return func(o *TraceOptions) {
		o.Close = b
	}
}

// WithExpire if true will allow tracing on the expire call.
func WithExpire(b bool) TraceOption {
	return func(o *TraceOptions) {
		o.Expire = b
	}
}
