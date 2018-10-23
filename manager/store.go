package manager

//go:generate mockgen -destination=./mocks/mock_manager.go -package=mocks github.com/kihamo/runtime-config/manager Store

import (
	"context"

	"github.com/kihamo/runtime-config/config"
)

type Store interface {
	Versions(context.Context) ([]config.Version, error)
	Variables(context.Context, config.Version) ([]config.Variable, error)
	SetVersionChangeCallback(config.VersionChangeCallback) error
	SetVariableChangeCallback(config.Version, config.VariableChangeCallback) error
	SetVariableChangeByNameCallback(config.Version, string, config.VariableChangeCallback) error
}
