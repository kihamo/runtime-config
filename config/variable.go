package config

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
