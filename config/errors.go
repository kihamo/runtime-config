package config

import (
	"errors"
)

var (
	ErrorNotImplemented   = errors.New("Not implemented")
	ErrorVersionNotFound  = errors.New("Version not found")
	ErrorVariableNotFound = errors.New("Variable not found")
)

func IsNotImplemented(err error) bool {
	return err == ErrorNotImplemented
}

func IsVersionNotFound(err error) bool {
	return err == ErrorVersionNotFound
}

func IsVariableNotFound(err error) bool {
	return err == ErrorVariableNotFound
}
