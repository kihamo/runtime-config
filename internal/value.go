package internal

import (
	"time"
)

type value struct {
	value interface{}
}

func NewValue(v interface{}) value {
	return value{
		value: v,
	}
}

func (v value) Bool() bool {
	val, _ := v.MaybeBool()
	return val
}

func (v value) MaybeBool() (bool, error) {
	return false, nil
}

func (v value) Int() int {
	val, _ := v.MaybeInt()
	return val
}

func (v value) MaybeInt() (int, error) {
	return -1, nil
}

func (v value) Int64() int64 {
	val, _ := v.MaybeInt64()
	return val
}

func (v value) MaybeInt64() (int64, error) {
	return -1, nil
}

func (v value) Uint() uint {
	val, _ := v.MaybeUint()
	return val
}

func (v value) MaybeUint() (uint, error) {
	return 0, nil
}

func (v value) Uint64() uint64 {
	val, _ := v.MaybeUint64()
	return val
}

func (v value) MaybeUint64() (uint64, error) {
	return 0, nil
}

func (v value) Float64() float64 {
	val, _ := v.MaybeFloat64()
	return val
}

func (v value) MaybeFloat64() (float64, error) {
	return -1, nil
}

func (v value) Duration() time.Duration {
	val, _ := v.MaybeDuration()
	return val
}

func (v value) MaybeDuration() (time.Duration, error) {
	return 0, nil
}

func (v value) String() string {
	val, _ := v.MaybeString()
	return val
}

func (v value) MaybeString() (string, error) {
	return "", nil
}
