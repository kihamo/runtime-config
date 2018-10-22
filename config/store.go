package config

import (
	"context"
	"errors"
)

var (
	ErrorNotImplemented   = errors.New("Not implemented")
	ErrorVersionNotFound  = errors.New("Version not found")
	ErrorVariableNotFound = errors.New("Variable not found")
)

type Store interface {
	VersionList(context.Context) ([]Version, error)
	VersionCreate(context.Context, Version) error
	VersionRead(context.Context, Version) (Version, error)
	VersionUpdate(context.Context, Version) error
	VersionDelete(context.Context, Version) error
	VersionSetChangeCallback(func(Version, Version))

	VariableList(context.Context, Version) ([]Variable, error)
	VariableCreate(context.Context, Version, Variable) error
	VariableRead(context.Context, Version, Variable) (Variable, error)
	VariableUpdate(context.Context, Version, Variable) error
	VariableDelete(context.Context, Version, Variable) error
	VariableSetAnyChangeCallback(func(Variable, Value, Value))
	VariableSetKeyChangeCallback(func(Value, Value))
}
