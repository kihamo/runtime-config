package etcd

import (
	"context"

	"github.com/kihamo/runtime-config/config"
	"github.com/kihamo/runtime-config/internal"
	"github.com/pkg/errors"
	"go.etcd.io/etcd/clientv3"
)

const (
	keyPrefix = "/config/"
)

// Store implements
type Store struct {
	client *clientv3.Client
}

// NewStore creates new instance of Store
func NewStore(client *clientv3.Client) *Store {
	return &Store{
		client: client,
	}
}

func (s *Store) Versions(context.Context) ([]config.Version, error) {
	return nil, config.ErrNotImplemented
}

func (s *Store) Variables(ctx context.Context, version config.Version) ([]config.Variable, error) {
	key, err := getVersionKey(version.ProjectID(), version.ID())
	if err != nil {
		return nil, err
	}

	response, err := s.client.Get(ctx, key)
	if err != nil {
		return nil, err
	}

	variables := make([]config.Variable, response.Count, response.Count)

	for i, kv := range response.Kvs {
		variables[i] = internal.NewVariable(string(kv.Key), string(kv.Value))
	}

	return variables, nil
}

func (s *Store) VariableCreate(context.Context, config.Version, config.Variable) error {
	return config.ErrNotImplemented
}

func (s *Store) VariableRead(ctx context.Context, version config.Version, variable config.Variable) (config.Variable, error) {
	key, err := getVariableKey(version.ProjectID(), version.ID(), variable.Name())
	if err != nil {
		return nil, err
	}

	response, err := s.client.Get(ctx, key)
	if err != nil {
		return nil, err
	}

	if response.Count == 0 {
		return nil, config.ErrVariableNotFound
	}

	return variable, err
}

func (s *Store) VariableUpdate(context.Context, config.Version, config.Variable) error {
	return config.ErrNotImplemented
}

func (s *Store) VariableDelete(context.Context, config.Version, config.Variable) error {
	return config.ErrNotImplemented
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

func getVersionKey(projectID, versionID string) (string, error) {
	if projectID == "" {
		return "", errors.New("project ID is empty")
	}

	if versionID == "" {
		return "", errors.New("version ID is empty")
	}

	return keyPrefix + projectID + "/" + versionID, nil
}

func getVariableKey(projectID, versionID, variableName string) (string, error) {
	if variableName == "" {
		return "", errors.New("variable name is empty")
	}

	pathVersion, err := getVersionKey(projectID, versionID)
	if err != nil {
		return "", err
	}

	return pathVersion + "/" + variableName, nil
}
