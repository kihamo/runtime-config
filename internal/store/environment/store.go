package environment

import (
	"context"
	"os"
	"strings"

	"github.com/kihamo/runtime-config/config"
	"github.com/kihamo/runtime-config/internal"
)

const (
	spaceSeparator    = "_"
	keyValueSeparator = "="
)

type Store struct {
	prefix    string
	variables map[string]map[string]config.Variable
}

func NewStore(prefix string) *Store {
	return &Store{
		prefix:    strings.ToUpper(prefix),
		variables: make(map[string]map[string]config.Variable),
	}
}

func (s *Store) Versions(context.Context, string) ([]config.Version, error) {
	return nil, config.ErrNotImplemented
}

func (s *Store) read(version config.Version) {
	if _, ok := s.variables[version.ProjectID()]; ok {
		return
	}

	envs := os.Environ()
	variables := make(map[string]config.Variable, len(envs))
	prefix := s.prefix + strings.ToUpper(version.ProjectID()) + spaceSeparator
	skipSymbols := len(prefix)

	for _, v := range envs {
		pair := strings.Split(v, keyValueSeparator)
		key := strings.ToUpper(pair[0])

		if strings.HasPrefix(key, prefix) {
			var value interface{}
			if pair[1] != "" {
				value = pair[1]
			}

			key := pair[0][skipSymbols:]
			variables[strings.ToUpper(key)] = internal.NewVariable(key, value)
		}
	}

	s.variables[version.ProjectID()] = variables
}

func (s *Store) Variables(ctx context.Context, version config.Version) ([]config.Variable, error) {
	s.read(version)

	variables := make([]config.Variable, 0, len(s.variables[version.ProjectID()]))
	for _, variable := range s.variables[version.ProjectID()] {
		variables = append(variables, variable)
	}

	return variables, nil
}

func (s *Store) VariableCreate(ctx context.Context, version config.Version, variable config.Variable) error {
	s.read(version)
	projectKey := version.ProjectID()

	if _, ok := s.variables[projectKey]; ok {
		s.variables[projectKey] = make(map[string]config.Variable, 1)
	}

	s.variables[projectKey][strings.ToUpper(variable.Name())] = variable
	return nil
}

func (s *Store) VariableRead(ctx context.Context, version config.Version, variable config.Variable) (config.Variable, error) {
	s.read(version)

	if list, ok := s.variables[version.ProjectID()]; ok {
		if variable, ok := list[strings.ToUpper(variable.Name())]; ok {
			return variable, nil
		}
	}

	return variable, config.ErrVariableNotFound
}

func (s *Store) VariableUpdate(ctx context.Context, version config.Version, variable config.Variable) error {
	s.read(version)
	projectKey := version.ProjectID()

	if list, ok := s.variables[projectKey]; ok {
		variableKey := strings.ToUpper(variable.Name())

		if _, ok = list[variableKey]; ok {
			s.variables[projectKey][variableKey] = variable
			return nil
		}
	}

	return config.ErrVariableNotFound
}

func (s *Store) VariableDelete(ctx context.Context, version config.Version, variable config.Variable) error {
	s.read(version)
	projectKey := version.ProjectID()

	if list, ok := s.variables[projectKey]; ok {
		variableKey := strings.ToUpper(variable.Name())

		if _, ok = list[variableKey]; ok {
			delete(s.variables[projectKey], variableKey)
		}
	}

	return nil
}

func (s *Store) SetVersionChangeCallback(config.VersionChangeCallback) error {
	return config.ErrNotImplemented
}

func (s *Store) SetVariableChangeCallback(config.Version, config.VariableChangeCallback) error {
	return config.ErrNotImplemented
}

func (s *Store) SetVariableChangeByNameCallback(config.Version, string, config.VariableChangeCallback) error {
	return config.ErrNotImplemented
}
