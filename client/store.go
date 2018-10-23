package client

//go:generate mockgen -destination=./mocks/mock_client.go -package=mocks github.com/kihamo/runtime-config/client Store

import (
	"context"

	"github.com/kihamo/runtime-config/config"
)

type Store interface {
	Variables(context.Context, config.Version) ([]config.Variable, error)
	VariableRead(ctx context.Context, version config.Version, variable config.Variable) (config.Variable, error)
	SetVariableChangeCallback(config.Version, config.VariableChangeCallback) error
	SetVariableChangeByNameCallback(config.Version, string, config.VariableChangeCallback) error
}
