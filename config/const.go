package config

import (
	"errors"
)

var (
	ErrorNotImplemented   = errors.New("Not implemented")
	ErrorVersionNotFound  = errors.New("Version not found")
	ErrorVariableNotFound = errors.New("Variable not found")
)
