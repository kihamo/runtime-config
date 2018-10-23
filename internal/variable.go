package internal

import (
	"github.com/kihamo/runtime-config/config"
)

type variable struct {
	name       string
	value      config.Value
	usage      string
	group      string
	isOptional bool
	isEditable bool
}

func NewVariable(name string, value interface{}) variable {
	return variable{
		name:  name,
		value: NewValue(value),
	}
}

func (v variable) Name() string {
	return v.name
}

func (v variable) Type() config.VariableType {
	return config.VariableTypeStr
}

func (v variable) Value() config.Value {
	return v.value
}

func (v variable) Usage() string {
	return v.usage
}

func (v variable) Group() string {
	return v.group
}

func (v variable) IsOptional() bool {
	return v.isOptional
}

func (v variable) IsEditable() bool {
	return v.isEditable
}
