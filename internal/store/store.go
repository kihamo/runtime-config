package store

import (
	"context"

	"github.com/kihamo/runtime-config/config"
)

type Store interface {
	VersionList(context.Context) ([]config.Version, error)
	VersionCreate(context.Context, config.Version) error
	VersionRead(context.Context, config.Version) (config.Version, error)
	VersionUpdate(context.Context, config.Version) error
	VersionDelete(context.Context, config.Version) error
	VersionSetChangeCallback(func(config.Version, config.Version))

	VariableList(context.Context, config.Version) ([]config.Variable, error)
	VariableCreate(context.Context, config.Version, config.Variable) error
	VariableRead(context.Context, config.Version, config.Variable) (config.Variable, error)
	VariableUpdate(context.Context, config.Version, config.Variable) error
	VariableDelete(context.Context, config.Version, config.Variable) error
	VariableSetAnyChangeCallback(func(config.Variable, config.Value, config.Value))
	VariableSetKeyChangeCallback(func(config.Value, config.Value))
}
