package updater

import (
	"context"
	"fmt"

	"github.com/kihamo/runtime-config/config"
	"github.com/kihamo/runtime-config/internal/store"
)

type Updater struct {
	version config.Version
	store   store.Store
}

func NewUpdater(ctx context.Context, version config.Version, store store.Store) (*Updater, error) {
	// TODO: create version if needed

	u := &Updater{
		version: version,
		store:   store,
	}

	return u, nil
}

func (u *Updater) AddVariable(ctx context.Context, variable config.Variable) error {
	variable, err := u.store.VariableRead(ctx, u.version, variable)

	fmt.Println(variable, err)

	if err != nil {
		fmt.Println(err, config.ErrVariableNotFound)

		if err != config.ErrVariableNotFound {
			return err
		}

		fmt.Println(err, "222")
	}

	fmt.Println(variable, err)

	return nil
}
