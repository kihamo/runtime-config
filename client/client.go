package client

import (
	"context"
	"errors"

	"github.com/kihamo/runtime-config/config"
	"github.com/kihamo/runtime-config/internal"
)

type Store interface {
	Variables(context.Context, config.Version) ([]config.Variable, error)
	VariableRead(ctx context.Context, version config.Version, variable config.Variable) (config.Variable, error)
	SetVersionChangeCallback(config.VersionChangeCallback) error
	SetVariableChangeCallback(config.Version, config.VariableChangeCallback) error
	SetVariableChangeByNameCallback(config.Version, string, config.VariableChangeCallback) error
}

type Client struct {
	version config.Version
	stores  []Store
}

func NewClient(ctx context.Context, version config.Version, stores ...Store) (*Client, error) {
	if len(stores) == 0 {
		return nil, errors.New("Stores isn't set")
	}

	c := &Client{
		version: version,
		stores:  stores,
	}

	return c, nil
}

func (c *Client) Variables(ctx context.Context) ([]config.Variable, error) {
	tmp := make(map[string]config.Variable)
	totalStoresCount := len(c.stores) - 1

	for i := totalStoresCount; i >= 0; i-- {
		storeVariables, err := c.stores[i].Variables(ctx, c.version)
		if err != nil {
			return nil, err
		}

		if totalStoresCount == 0 {
			return storeVariables, nil
		}

		for _, variable := range storeVariables {
			tmp[variable.Name()] = variable
		}
	}

	variables := make([]config.Variable, 0, len(tmp))
	for _, variable := range tmp {
		variables = append(variables, variable)
	}

	return variables, nil
}

func (c *Client) GetVariable(ctx context.Context, name string) (config.Variable, error) {
	v := internal.NewVariable(name, nil)

	for _, store := range c.stores {
		variable, err := store.VariableRead(ctx, c.version, v)

		if err != nil && config.IsVariableNotFound(err) {
			continue
		}

		return variable, err
	}

	return nil, config.ErrorVariableNotFound
}

func (c *Client) SetVersionChangeCallback(callback config.VersionChangeCallback) error {
	for _, store := range c.stores {
		err := store.SetVersionChangeCallback(callback)
		if err != nil && !config.IsNotImplemented(err) {
			return err
		}
	}

	return nil
}

func (c *Client) SetVariableChangeCallback(callback config.VariableChangeCallback) error {
	for _, store := range c.stores {
		err := store.SetVariableChangeCallback(c.version, callback)
		if err != nil && !config.IsNotImplemented(err) {
			return err
		}
	}

	return nil
}

func (c *Client) SetVariableChangeByNameCallback(name string, callback func(config.Variable, config.Value, config.Value)) error {
	for _, store := range c.stores {
		err := store.SetVariableChangeByNameCallback(c.version, name, callback)
		if err != nil && !config.IsNotImplemented(err) {
			return err
		}
	}

	return nil
}
