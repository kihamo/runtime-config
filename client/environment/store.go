package environment

import (
	"github.com/kihamo/runtime-config/client"
	"github.com/kihamo/runtime-config/internal/store/environment"
)

func NewStore(prefix string) client.Store {
	return environment.NewStore(prefix)
}
