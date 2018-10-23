package etcd

import (
	"context"
	"fmt"
	"strings"

	"github.com/kihamo/runtime-config/config"
	"github.com/kihamo/runtime-config/internal"
	"go.etcd.io/etcd/clientv3"
)

const (
	keyPrefix = "/config/"
)

// Store implements
type Store struct {
	client *clientv3.Client
}

// NewStore creates new instance of Store from provided client
func NewStore(client *clientv3.Client) *Store {
	return &Store{client: client}
}

// NewEtcdClient creates new etcd client from provided config
func NewEtcdClient(c *Config) (*clientv3.Client, error) {
	// Validate our config
	if err := validateConfig(c); err != nil {
		return nil, fmt.Errorf("unable to validate config: %v", err)
	}

	// Setup etcd config
	c.setDefaults()
	endpoints := strings.Split(c.Endpoints, ",")
	etcdConfig := clientv3.Config{
		Endpoints:   endpoints,
		DialTimeout: c.Timeout,
	}

	// Create secure connection if certs were provided
	if c.tlsEnabled {
		tlsConfig, err := newTLSConfig(c.ServiceCert, c.ServiceKey, c.CACert)
		if err != nil {
			return nil, fmt.Errorf("unable to create TLS config: %v", err)
		}

		etcdConfig.TLS = tlsConfig
	}

	// Setup client
	etcdClient, err := clientv3.New(etcdConfig)
	if err != nil {
		return nil, fmt.Errorf("unable to create new etcd client: %v", err)
	}

	return etcdClient, nil
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
		return "", fmt.Errorf("project ID is empty")
	}

	if versionID == "" {
		return "", fmt.Errorf("version ID is empty")
	}

	return keyPrefix + projectID + "/" + versionID, nil
}

func getVariableKey(projectID, versionID, variableName string) (string, error) {
	if variableName == "" {
		return "", fmt.Errorf("variable name is empty")
	}

	pathVersion, err := getVersionKey(projectID, versionID)
	if err != nil {
		return "", err
	}

	return pathVersion + "/" + variableName, nil
}
