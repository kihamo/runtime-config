package manager

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

type manager struct {
	store Store
}

func NewManager(store Store) *manager {
	return &manager{
		store: store,
	}
}

func (m *manager) Versions(ctx context.Context) ([]config.Version, error) {
	return m.Versions(ctx)
}

func (m *manager) VersionCurrent(ctx context.Context) (config.Version, error) {
	return nil, nil
}

func (m *manager) Values(ctx context.Context, version config.Version) (map[string]config.Value, error) {
	storeVariables, err := m.store.Variables(ctx, version)
	if err != nil {
		return nil, err
	}

	values := make(map[string]config.Value, len(storeVariables))

	for _, variable := range storeVariables {
		values[variable.Name()] = variable.Value()
	}

	return values, nil
}

func (m *manager) SetVersionChangeCallback(callback config.VersionChangeCallback) error {
	err := m.store.SetVersionChangeCallback(callback)
	if err != nil && !config.IsNotImplemented(err) {
		return err
	}

	return nil
}

func (m *manager) SetVariableChangeCallback(version config.Version, callback config.VariableChangeCallback) error {
	err := m.store.SetVariableChangeCallback(version, callback)
	if err != nil && !config.IsNotImplemented(err) {
		return err
	}

	return nil
}

func (m *manager) SetVariableChangeByNameCallback(version config.Version, name string, callback config.VariableChangeCallback) error {
	err := m.store.SetVariableChangeByNameCallback(version, name, callback)
	if err != nil && !config.IsNotImplemented(err) {
		return err
	}

	return nil
}
