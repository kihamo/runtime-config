package manager

import (
	"context"

	"github.com/kihamo/runtime-config/config"
)

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

func (m *manager) Variables(ctx context.Context, version config.Version) ([]config.Variable, error) {
	return m.store.Variables(ctx, version)
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
