package store

import (
	"context"

	"github.com/kihamo/runtime-config/config"
)

type Store interface {
	Versions(context.Context) ([]config.Version, error)
	Variables(context.Context, config.Version) ([]config.Variable, error)
	VariableCreate(context.Context, config.Version, config.Variable) error
	VariableRead(context.Context, config.Version, config.Variable) (config.Variable, error)
	VariableUpdate(context.Context, config.Version, config.Variable) error
	VariableDelete(context.Context, config.Version, config.Variable) error
	SetVersionChangeCallback(config.VersionChangeCallback) error
	SetVariableChangeCallback(config.VariableChangeCallback) error
	SetVariableChangeByNameCallback(string, config.VariableChangeCallback) error
}
