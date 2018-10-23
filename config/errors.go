package config

import (
	"errors"
)

// Common errors
var (
	ErrNotImplemented   = errors.New("not implemented")
	ErrVersionNotFound  = errors.New("version not found")
	ErrVariableNotFound = errors.New("variable not found")
)

// IsNotImplemented checks if the provider error equals ErrNotImplemented
func IsNotImplemented(err error) bool {
	return err == ErrNotImplemented
}

// IsVersionNotFound checks if the provider error equals ErrVersionNotFound
func IsVersionNotFound(err error) bool {
	return err == ErrVersionNotFound
}

// IsVariableNotFound checks if the provider error equals ErrVariableNotFound
func IsVariableNotFound(err error) bool {
	return err == ErrVariableNotFound
}
