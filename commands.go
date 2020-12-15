package ocredis

import "fmt"

// Cmd interface matches the Cmd struct returned by redis clients
type Cmd interface {
	Err() error
	fmt.Stringer
}

// StatusCmd interface matches the StatusCmd struct returned by redis clients
type StatusCmd interface {
	Err() error
	Result() (string, error)
	String() string
	Val() string
}

// StringCmd interface matches the StringCmd struct returned by redis clients
type StringCmd interface {
	Bytes() ([]byte, error)
	Err() error
	Float64() (float64, error)
	Int64() (int64, error)
	Result() (string, error)
	Scan(val interface{}) error
	String() string
	Uint64() (uint64, error)
	Val() string
}

// IntCmd interface matches the IntCmd struct returned by redis clients
type IntCmd interface {
	Val() int64
	String() string
	Err() error
	Result() (int64, error)
}

// BoolCmd interface matches the BoolCmd struct returned by redis clients
type BoolCmd interface {
	Err() error
	Result() (bool, error)
	String() string
	Val() bool
}
