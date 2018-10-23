package manager

import (
	"context"

	"github.com/kihamo/runtime-config/config"
	"github.com/kihamo/runtime-config/internal"
)

type manager struct {
	store Store
}

func NewManager(store Store) *manager {
	return &manager{
		store: store,
	}
}

func (m *manager) Versions(ctx context.Context, projectID string) ([]config.Version, error) {
	return m.store.Versions(ctx, projectID)
}

func (m *manager) VersionCurrent(ctx context.Context, projectID string) (config.Version, error) {
	return nil, nil
}

func (m *manager) Variables(ctx context.Context, projectID, versionID string) ([]config.Variable, error) {
	return m.store.Variables(ctx, internal.NewVersion(projectID, versionID))
}

func (m *manager) SetVersionChangeCallback(callback config.VersionChangeCallback) error {
	err := m.store.SetVersionChangeCallback(callback)
	if err != nil && !config.IsNotImplemented(err) {
		return err
	}

	return nil
}

func (m *manager) SetVariableChangeCallback(projectID, versionID string, callback config.VariableChangeCallback) error {
	err := m.store.SetVariableChangeCallback(internal.NewVersion(projectID, versionID), callback)
	if err != nil && !config.IsNotImplemented(err) {
		return err
	}

	return nil
}

func (m *manager) SetVariableChangeByNameCallback(projectID, versionID string, name string, callback config.VariableChangeCallback) error {
	err := m.store.SetVariableChangeByNameCallback(internal.NewVersion(projectID, versionID), name, callback)
	if err != nil && !config.IsNotImplemented(err) {
		return err
	}

	return nil
}
