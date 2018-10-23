// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/kihamo/runtime-config/config (interfaces: Value)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"
	time "time"

	gomock "github.com/golang/mock/gomock"
)

// MockValue is a mock of Value interface
type MockValue struct {
	ctrl     *gomock.Controller
	recorder *MockValueMockRecorder
}

// MockValueMockRecorder is the mock recorder for MockValue
type MockValueMockRecorder struct {
	mock *MockValue
}

// NewMockValue creates a new mock instance
func NewMockValue(ctrl *gomock.Controller) *MockValue {
	mock := &MockValue{ctrl: ctrl}
	mock.recorder = &MockValueMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockValue) EXPECT() *MockValueMockRecorder {
	return m.recorder
}

// Bool mocks base method
func (m *MockValue) Bool() bool {
	ret := m.ctrl.Call(m, "Bool")
	ret0, _ := ret[0].(bool)
	return ret0
}

// Bool indicates an expected call of Bool
func (mr *MockValueMockRecorder) Bool() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Bool", reflect.TypeOf((*MockValue)(nil).Bool))
}

// Duration mocks base method
func (m *MockValue) Duration() time.Duration {
	ret := m.ctrl.Call(m, "Duration")
	ret0, _ := ret[0].(time.Duration)
	return ret0
}

// Duration indicates an expected call of Duration
func (mr *MockValueMockRecorder) Duration() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Duration", reflect.TypeOf((*MockValue)(nil).Duration))
}

// Float64 mocks base method
func (m *MockValue) Float64() float64 {
	ret := m.ctrl.Call(m, "Float64")
	ret0, _ := ret[0].(float64)
	return ret0
}

// Float64 indicates an expected call of Float64
func (mr *MockValueMockRecorder) Float64() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Float64", reflect.TypeOf((*MockValue)(nil).Float64))
}

// Int mocks base method
func (m *MockValue) Int() int {
	ret := m.ctrl.Call(m, "Int")
	ret0, _ := ret[0].(int)
	return ret0
}

// Int indicates an expected call of Int
func (mr *MockValueMockRecorder) Int() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Int", reflect.TypeOf((*MockValue)(nil).Int))
}

// Int64 mocks base method
func (m *MockValue) Int64() int64 {
	ret := m.ctrl.Call(m, "Int64")
	ret0, _ := ret[0].(int64)
	return ret0
}

// Int64 indicates an expected call of Int64
func (mr *MockValueMockRecorder) Int64() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Int64", reflect.TypeOf((*MockValue)(nil).Int64))
}

// IsNil mocks base method
func (m *MockValue) IsNil() bool {
	ret := m.ctrl.Call(m, "IsNil")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsNil indicates an expected call of IsNil
func (mr *MockValueMockRecorder) IsNil() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsNil", reflect.TypeOf((*MockValue)(nil).IsNil))
}

// MaybeBool mocks base method
func (m *MockValue) MaybeBool() (bool, error) {
	ret := m.ctrl.Call(m, "MaybeBool")
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MaybeBool indicates an expected call of MaybeBool
func (mr *MockValueMockRecorder) MaybeBool() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MaybeBool", reflect.TypeOf((*MockValue)(nil).MaybeBool))
}

// MaybeDuration mocks base method
func (m *MockValue) MaybeDuration() (time.Duration, error) {
	ret := m.ctrl.Call(m, "MaybeDuration")
	ret0, _ := ret[0].(time.Duration)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MaybeDuration indicates an expected call of MaybeDuration
func (mr *MockValueMockRecorder) MaybeDuration() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MaybeDuration", reflect.TypeOf((*MockValue)(nil).MaybeDuration))
}

// MaybeFloat64 mocks base method
func (m *MockValue) MaybeFloat64() (float64, error) {
	ret := m.ctrl.Call(m, "MaybeFloat64")
	ret0, _ := ret[0].(float64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MaybeFloat64 indicates an expected call of MaybeFloat64
func (mr *MockValueMockRecorder) MaybeFloat64() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MaybeFloat64", reflect.TypeOf((*MockValue)(nil).MaybeFloat64))
}

// MaybeInt mocks base method
func (m *MockValue) MaybeInt() (int, error) {
	ret := m.ctrl.Call(m, "MaybeInt")
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MaybeInt indicates an expected call of MaybeInt
func (mr *MockValueMockRecorder) MaybeInt() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MaybeInt", reflect.TypeOf((*MockValue)(nil).MaybeInt))
}

// MaybeInt64 mocks base method
func (m *MockValue) MaybeInt64() (int64, error) {
	ret := m.ctrl.Call(m, "MaybeInt64")
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MaybeInt64 indicates an expected call of MaybeInt64
func (mr *MockValueMockRecorder) MaybeInt64() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MaybeInt64", reflect.TypeOf((*MockValue)(nil).MaybeInt64))
}

// MaybeString mocks base method
func (m *MockValue) MaybeString() (string, error) {
	ret := m.ctrl.Call(m, "MaybeString")
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MaybeString indicates an expected call of MaybeString
func (mr *MockValueMockRecorder) MaybeString() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MaybeString", reflect.TypeOf((*MockValue)(nil).MaybeString))
}

// MaybeUint mocks base method
func (m *MockValue) MaybeUint() (uint, error) {
	ret := m.ctrl.Call(m, "MaybeUint")
	ret0, _ := ret[0].(uint)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MaybeUint indicates an expected call of MaybeUint
func (mr *MockValueMockRecorder) MaybeUint() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MaybeUint", reflect.TypeOf((*MockValue)(nil).MaybeUint))
}

// MaybeUint64 mocks base method
func (m *MockValue) MaybeUint64() (uint64, error) {
	ret := m.ctrl.Call(m, "MaybeUint64")
	ret0, _ := ret[0].(uint64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MaybeUint64 indicates an expected call of MaybeUint64
func (mr *MockValueMockRecorder) MaybeUint64() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MaybeUint64", reflect.TypeOf((*MockValue)(nil).MaybeUint64))
}

// String mocks base method
func (m *MockValue) String() string {
	ret := m.ctrl.Call(m, "String")
	ret0, _ := ret[0].(string)
	return ret0
}

// String indicates an expected call of String
func (mr *MockValueMockRecorder) String() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "String", reflect.TypeOf((*MockValue)(nil).String))
}

// Uint mocks base method
func (m *MockValue) Uint() uint {
	ret := m.ctrl.Call(m, "Uint")
	ret0, _ := ret[0].(uint)
	return ret0
}

// Uint indicates an expected call of Uint
func (mr *MockValueMockRecorder) Uint() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Uint", reflect.TypeOf((*MockValue)(nil).Uint))
}

// Uint64 mocks base method
func (m *MockValue) Uint64() uint64 {
	ret := m.ctrl.Call(m, "Uint64")
	ret0, _ := ret[0].(uint64)
	return ret0
}

// Uint64 indicates an expected call of Uint64
func (mr *MockValueMockRecorder) Uint64() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Uint64", reflect.TypeOf((*MockValue)(nil).Uint64))
}
