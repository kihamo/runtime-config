package client

import (
	"context"

	"github.com/kihamo/runtime-config/config"
)

type Store interface {
	Variables(context.Context, config.Version) ([]config.Variable, error)
	SetVersionChangeCallback(config.VersionChangeCallback) error
	SetVariableChangeCallback(config.VariableChangeCallback) error
	SetVariableChangeByNameCallback(string, config.VariableChangeCallback) error
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
		err := store.SetVariableChangeCallback(callback)
		if err != nil && !config.IsNotImplemented(err) {
			return err
		}
	}

	return nil
}

func (c *Client) SetVariableChangeByNameCallback(name string, callback func(config.Variable, config.Value, config.Value)) error {
	for _, store := range c.stores {
		err := store.SetVariableChangeByNameCallback(name, callback)
		if err != nil && !config.IsNotImplemented(err) {
			return err
		}
	}

	return nil
}
