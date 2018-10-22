package config

import (
	"github.com/kihamo/runtime-config/config"
)

type Variable struct {
	name       string
	value      config.Value
	usage      string
	group      string
	isOptional bool
	isEditable bool
}

func NewVariable(name string, value interface{}) Variable {
	return Variable{
		name:  name,
		value: NewValue(value),
	}
}

func (v Variable) Name() string {
	return v.name
}

func (v Variable) Type() config.VariableType {
	return config.VariableTypeStr
}

func (v Variable) Value() config.Value {
	return v.value
}

func (v Variable) Usage() string {
	return v.usage
}

func (v Variable) Group() string {
	return v.group
}

func (v Variable) IsOptional() bool {
	return v.isOptional
}

func (v Variable) IsEditable() bool {
	return v.isEditable
}
