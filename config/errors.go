package config

import (
	"errors"
)

var (
	ErrNotImplemented   = errors.New("Not implemented")
	ErrVersionNotFound  = errors.New("Version not found")
	ErrVariableNotFound = errors.New("Variable not found")
)

func IsNotImplemented(err error) bool {
	return err == ErrNotImplemented
}

func IsVersionNotFound(err error) bool {
	return err == ErrVersionNotFound
}

func IsVariableNotFound(err error) bool {
	return err == ErrVariableNotFound
}
