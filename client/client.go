package client

import (
	"context"
	"github.com/kihamo/runtime-config/config"
)

type Client struct {
	version   config.Version
	stores    []config.Store
	variables map[string]config.Variable
}

func NewClient(ctx context.Context, version config.Version, stores ...config.Store) (*Client, error) {
	c := &Client{
		version:   version,
		stores:    stores,
		variables: make(map[string]config.Variable),
	}

	// TODO: parallel check
	for _, store := range stores {
		variables, err := store.VariableList(ctx, version)
		if err != nil {
			return nil, err
		}

		for _, variable := range variables {
			c.variables[variable.Name()] = variable
		}
	}

	return c, nil
}
