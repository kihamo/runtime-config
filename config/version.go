package config

type Version interface {
	ID() string
	ProjectID() uint64
}

type VersionChangeCallback func(Version)
