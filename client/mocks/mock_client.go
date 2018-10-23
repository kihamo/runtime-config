// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/kihamo/runtime-config/client (interfaces: Store)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	config "github.com/kihamo/runtime-config/config"
)

// MockStore is a mock of Store interface
type MockStore struct {
	ctrl     *gomock.Controller
	recorder *MockStoreMockRecorder
}

// MockStoreMockRecorder is the mock recorder for MockStore
type MockStoreMockRecorder struct {
	mock *MockStore
}

// NewMockStore creates a new mock instance
func NewMockStore(ctrl *gomock.Controller) *MockStore {
	mock := &MockStore{ctrl: ctrl}
	mock.recorder = &MockStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockStore) EXPECT() *MockStoreMockRecorder {
	return m.recorder
}

// SetVariableChangeByNameCallback mocks base method
func (m *MockStore) SetVariableChangeByNameCallback(arg0 config.Version, arg1 string, arg2 config.VariableChangeCallback) error {
	ret := m.ctrl.Call(m, "SetVariableChangeByNameCallback", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetVariableChangeByNameCallback indicates an expected call of SetVariableChangeByNameCallback
func (mr *MockStoreMockRecorder) SetVariableChangeByNameCallback(arg0, arg1, arg2 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetVariableChangeByNameCallback", reflect.TypeOf((*MockStore)(nil).SetVariableChangeByNameCallback), arg0, arg1, arg2)
}

// SetVariableChangeCallback mocks base method
func (m *MockStore) SetVariableChangeCallback(arg0 config.Version, arg1 config.VariableChangeCallback) error {
	ret := m.ctrl.Call(m, "SetVariableChangeCallback", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetVariableChangeCallback indicates an expected call of SetVariableChangeCallback
func (mr *MockStoreMockRecorder) SetVariableChangeCallback(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetVariableChangeCallback", reflect.TypeOf((*MockStore)(nil).SetVariableChangeCallback), arg0, arg1)
}

// VariableRead mocks base method
func (m *MockStore) VariableRead(arg0 context.Context, arg1 config.Version, arg2 config.Variable) (config.Variable, error) {
	ret := m.ctrl.Call(m, "VariableRead", arg0, arg1, arg2)
	ret0, _ := ret[0].(config.Variable)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// VariableRead indicates an expected call of VariableRead
func (mr *MockStoreMockRecorder) VariableRead(arg0, arg1, arg2 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VariableRead", reflect.TypeOf((*MockStore)(nil).VariableRead), arg0, arg1, arg2)
}

// Variables mocks base method
func (m *MockStore) Variables(arg0 context.Context, arg1 config.Version) ([]config.Variable, error) {
	ret := m.ctrl.Call(m, "Variables", arg0, arg1)
	ret0, _ := ret[0].([]config.Variable)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Variables indicates an expected call of Variables
func (mr *MockStoreMockRecorder) Variables(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Variables", reflect.TypeOf((*MockStore)(nil).Variables), arg0, arg1)
}
