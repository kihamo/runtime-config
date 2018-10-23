package client

import (
	"context"

	rc "github.com/kihamo/runtime-config"

	"github.com/kihamo/runtime-config/config"
)

type Store interface {
	Variables(context.Context, config.Version) ([]config.Variable, error)
	VariableRead(ctx context.Context, version config.Version, variable config.Variable) (config.Variable, error)
	SetVersionChangeCallback(config.VersionChangeCallback) error
	SetVariableChangeCallback(config.Version, config.VariableChangeCallback) error
	SetVariableChangeByNameCallback(config.Version, string, config.VariableChangeCallback) error
}

type Client struct {
	version   config.Version
	stores    []Store
	variables map[string]config.Variable
}

func NewClient(ctx context.Context, version config.Version, stores ...Store) (*Client, error) {
	c := &Client{
		version:   version,
		stores:    stores,
		variables: make(map[string]config.Variable),
	}

	// TODO: parallel check
	for _, s := range stores {
		variables, err := s.Variables(ctx, version)
		if err != nil {
			return nil, err
		}

		for _, variable := range variables {
			c.variables[variable.Name()] = variable
		}
	}

	return c, nil
}

func (c *Client) Variables(ctx context.Context) ([]config.Variable, error) {
	// TODO: merge
	return c.stores[0].Variables(ctx, c.version)
}

func (c *Client) GetVariable(ctx context.Context, name string) (config.Variable, error) {
	v := rc.NewVariable(name, nil)

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
