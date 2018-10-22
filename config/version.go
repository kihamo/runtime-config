package config

type Version interface {
	ID() string
	ProjectID() uint64
}
