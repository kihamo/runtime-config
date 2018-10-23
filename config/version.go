package config

//go:generate mockgen -destination=./mocks/mock_version.go -package=mocks github.com/kihamo/runtime-config/config Version

type Version interface {
	ID() string
	ProjectID() string
}

type VersionChangeCallback func(Version)
