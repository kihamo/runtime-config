package config

//go:generate mockgen -destination=./mocks/mock_variable.go -package=mocks github.com/kihamo/runtime-config/config Variable

type VariableType int64

const (
	VariableTypeBool VariableType = iota
	VariableTypeInt
	VariableTypeInt64
	VariableTypeUint
	VariableTypeUint64
	VariableTypeFloat64
	VariableTypeDuration
	VariableTypeStr
)

type Variable interface {
	Name() string
	Type() VariableType
	Value() Value
	Usage() string
	Group() string
	IsOptional() bool
	IsEditable() bool
}

type VariableChangeCallback func(Variable, Value, Value)
