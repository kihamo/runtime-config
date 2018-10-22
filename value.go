package config

import (
	"time"
)

type Value struct {
	value interface{}
}

func NewValue(value interface{}) Value {
	return Value{
		value: value,
	}
}

func (v Value) Bool() bool {
	val, _ := v.MaybeBool()
	return val
}

func (v Value) MaybeBool() (bool, error) {
	return false, nil
}

func (v Value) Int() int {
	val, _ := v.MaybeInt()
	return val
}

func (v Value) MaybeInt() (int, error) {
	return -1, nil
}

func (v Value) Int64() int64 {
	val, _ := v.MaybeInt64()
	return val
}

func (v Value) MaybeInt64() (int64, error) {
	return -1, nil
}

func (v Value) Uint() uint {
	val, _ := v.MaybeUint()
	return val
}

func (v Value) MaybeUint() (uint, error) {
	return 0, nil
}

func (v Value) Uint64() uint64 {
	val, _ := v.MaybeUint64()
	return val
}

func (v Value) MaybeUint64() (uint64, error) {
	return 0, nil
}

func (v Value) Float64() float64 {
	val, _ := v.MaybeFloat64()
	return val
}

func (v Value) MaybeFloat64() (float64, error) {
	return -1, nil
}

func (v Value) Duration() time.Duration {
	val, _ := v.MaybeDuration()
	return val
}

func (v Value) MaybeDuration() (time.Duration, error) {
	return 0, nil
}

func (v Value) String() string {
	val, _ := v.MaybeString()
	return val
}

func (v Value) MaybeString() (string, error) {
	return "", nil
}
