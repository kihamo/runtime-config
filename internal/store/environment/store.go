package environment

import (
	"context"

	"github.com/kihamo/runtime-config/config"
)

type Store struct {
	prefix string
}

func NewStore(prefix string) *Store {
	return &Store{
		prefix: prefix,
	}
}

func (s *Store) read() {

}

func (s *Store) Versions(context.Context) ([]config.Version, error) {
	return nil, config.ErrorNotImplemented
}

func (s *Store) Variables(ctx context.Context, version config.Version) ([]config.Variable, error) {
	return nil, config.ErrorNotImplemented
}

func (s *Store) VariableCreate(context.Context, config.Version, config.Variable) error {
	return config.ErrorNotImplemented
}

func (s *Store) VariableRead(ctx context.Context, version config.Version, variable config.Variable) (config.Variable, error) {
	return nil, config.ErrorNotImplemented
}

func (s *Store) VariableUpdate(context.Context, config.Version, config.Variable) error {
	return config.ErrorNotImplemented
}

func (s *Store) VariableDelete(context.Context, config.Version, config.Variable) error {
	return config.ErrorNotImplemented
}

func (s *Store) SetVersionChangeCallback(config.VersionChangeCallback) error {
	return config.ErrorNotImplemented
}

func (s *Store) SetVariableChangeCallback(config.VariableChangeCallback) error {
	return config.ErrorNotImplemented
}

func (s *Store) SetVariableChangeByNameCallback(string, config.VariableChangeCallback) error {
	return config.ErrorNotImplemented
}
