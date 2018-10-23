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

func (c *Client) Values(ctx context.Context) (map[string]config.Value, error) {
	values := make(map[string]config.Value)

	for i := len(c.stores) - 1; i >= 0; i-- {
		storeVariables, err := c.stores[i].Variables(ctx, c.version)
		if err != nil {
			return nil, err
		}

		for _, variable := range storeVariables {
			values[variable.Name()] = variable.Value()
		}
	}

	return values, nil
}

func (c *Client) Value(ctx context.Context, name string) (config.Value, error) {
	v := internal.NewVariable(name, nil)

	for _, store := range c.stores {
		variable, err := store.VariableRead(ctx, c.version, v)

		if err != nil && config.IsVariableNotFound(err) {
			continue
		}

		return variable.Value(), err
	}

	return nil, config.ErrVariableNotFound
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

func (c *Client) SetVariableChangeByNameCallback(name string, callback config.VariableChangeCallback) error {
	for _, store := range c.stores {
		err := store.SetVariableChangeByNameCallback(c.version, name, callback)
		if err != nil && !config.IsNotImplemented(err) {
			return err
		}
	}

	return nil
}
