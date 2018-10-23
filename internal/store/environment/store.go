package environment

import (
	"context"
	"fmt"
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

func (s *Store) Versions(context.Context) ([]config.Version, error) {
	return nil, config.ErrorNotImplemented
}

func (s *Store) read(version config.Version) {
	if _, ok := s.variables[version.ProjectID()]; ok {
		return
	}

	envs := os.Environ()
	variables := make(map[string]config.Variable, len(envs))
	prefix := s.getVariableKeyPrefix(version.ProjectID())
	skipSymbols := len(prefix)

	for _, v := range envs {
		pair := strings.Split(v, keyValueSeparator)
		key := strings.ToUpper(pair[0])

		if strings.HasPrefix(key, prefix) {
			variables[strings.ToUpper(pair[0][skipSymbols:])] = internal.NewVariable(pair[0][skipSymbols:], pair[1])
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

func (s *Store) VariableCreate(context.Context, config.Version, config.Variable) error {
	return config.ErrorNotImplemented
}

func (s *Store) VariableRead(ctx context.Context, version config.Version, variable config.Variable) (config.Variable, error) {
	s.read(version)

	name := strings.ToUpper(variable.Name())

	fmt.Println(name, s.variables[version.ProjectID()])

	variable, ok := s.variables[version.ProjectID()][name]

	if !ok {
		return nil, config.ErrorVariableNotFound
	}

	return variable, nil
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

func (s *Store) SetVariableChangeCallback(config.Version, config.VariableChangeCallback) error {
	return config.ErrorNotImplemented
}

func (s *Store) SetVariableChangeByNameCallback(config.Version, string, config.VariableChangeCallback) error {
	return config.ErrorNotImplemented
}

func (s *Store) getVariableKeyPrefix(projectID string) string {
	return s.prefix + strings.ToUpper(projectID) + spaceSeparator
}
