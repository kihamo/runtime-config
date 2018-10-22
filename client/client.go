package client

import (
	"context"

	"github.com/kihamo/runtime-config/config"
)

type Store interface {
	VariableList(context.Context, config.Version) ([]config.Variable, error)
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
		variables, err := s.VariableList(ctx, version)
		if err != nil {
			return nil, err
		}

		for _, variable := range variables {
			c.variables[variable.Name()] = variable
		}
	}

	return c, nil
}
