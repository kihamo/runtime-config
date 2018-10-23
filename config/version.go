package config

type Version interface {
	ID() string
	ProjectID() string
}

type VersionChangeCallback func(Version)
