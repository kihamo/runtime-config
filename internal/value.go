package internal

import (
	"fmt"
	"strconv"
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

func (v value) IsNil() bool {
	return v.value == nil
}

func (v value) Bool() bool {
	val, _ := v.MaybeBool()
	return val
}

func (v value) MaybeBool() (bool, error) {
	raw, err := v.MaybeString()
	if err != nil {
		return false, err
	}

	val, err := strconv.ParseBool(raw)
	if err != nil {
		return false, fmt.Errorf("value cannot be parsed as bool: %s", err)
	}

	return val, nil
}

func (v value) Int() int {
	val, _ := v.MaybeInt()
	return val
}

func (v value) MaybeInt() (int, error) {
	raw, err := v.MaybeString()
	if err != nil {
		return 0, err
	}

	val, err := strconv.Atoi(raw)
	if err != nil {
		return 0, fmt.Errorf("value cannot be parsed as int: %s", err)
	}

	return val, nil
}

func (v value) Int64() int64 {
	val, _ := v.MaybeInt64()
	return val
}

func (v value) MaybeInt64() (int64, error) {
	raw, err := v.MaybeString()
	if err != nil {
		return 0, err
	}

	val, err := strconv.ParseInt(raw, 10, 64)
	if err != nil {
		return 0, fmt.Errorf("value cannot be parsed as int64: %s", err)
	}

	return val, nil
}

func (v value) Uint() uint {
	val, _ := v.MaybeUint()
	return val
}

func (v value) MaybeUint() (uint, error) {
	raw, err := v.MaybeString()
	if err != nil {
		return 0, err
	}

	val, err := strconv.ParseUint(raw, 10, 32)
	if err != nil {
		return 0, fmt.Errorf("value cannot be parsed as uint: %s", err)
	}

	return uint(val), nil
}

func (v value) Uint64() uint64 {
	val, _ := v.MaybeUint64()
	return val
}

func (v value) MaybeUint64() (uint64, error) {
	raw, err := v.MaybeString()
	if err != nil {
		return 0, err
	}

	val, err := strconv.ParseUint(raw, 10, 64)
	if err != nil {
		return 0, fmt.Errorf("value cannot be parsed as uint64: %s", err)
	}

	return val, nil
}

func (v value) Float64() float64 {
	val, _ := v.MaybeFloat64()
	return val
}

func (v value) MaybeFloat64() (float64, error) {
	raw, err := v.MaybeString()
	if err != nil {
		return 0, err
	}

	val, err := strconv.ParseFloat(raw, 64)
	if err != nil {
		return 0, fmt.Errorf("value cannot be parsed as float64: %s", err)
	}

	return val, nil
}

func (v value) Duration() time.Duration {
	val, _ := v.MaybeDuration()
	return val
}

func (v value) MaybeDuration() (time.Duration, error) {
	raw, err := v.MaybeString()
	if err != nil {
		return 0, err
	}

	val, err := time.ParseDuration(raw)
	if err != nil {
		return 0, fmt.Errorf("value cannot be parsed as duration: %s", err)
	}

	return val, nil
}

func (v value) String() string {
	val, _ := v.MaybeString()
	return val
}

func (v value) MaybeString() (string, error) {
	if v.IsNil() {
		return "", nil
	}

	return fmt.Sprintf("%s", v.value), nil
}

func (v value) GoString() string {
	return v.String()
}
