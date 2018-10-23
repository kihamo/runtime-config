package config

//go:generate mockgen -destination=./mocks/mock_value.go -package=mocks github.com/kihamo/runtime-config/config Value

import (
	"time"
)

type Value interface {
	IsNil() bool

	Bool() bool
	MaybeBool() (bool, error)

	Int() int
	MaybeInt() (int, error)

	Int64() int64
	MaybeInt64() (int64, error)

	Uint() uint
	MaybeUint() (uint, error)

	Uint64() uint64
	MaybeUint64() (uint64, error)

	Float64() float64
	MaybeFloat64() (float64, error)

	Duration() time.Duration
	MaybeDuration() (time.Duration, error)

	String() string
	MaybeString() (string, error)
}
