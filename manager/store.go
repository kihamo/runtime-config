package manager

import (
	"context"

	"github.com/kihamo/runtime-config/config"
)

type Store interface {
	Versions(context.Context, string) ([]config.Version, error)
	Variables(context.Context, config.Version) ([]config.Variable, error)
	SetVersionChangeCallback(config.VersionChangeCallback) error
	SetVariableChangeCallback(config.Version, config.VariableChangeCallback) error
	SetVariableChangeByNameCallback(config.Version, string, config.VariableChangeCallback) error
}
