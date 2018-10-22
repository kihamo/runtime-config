package manager

type Store interface {
}

type Manager struct {
	store Store
}

func NewManager(store Store) *Manager {
	return &Manager{
		store: store,
	}
}
