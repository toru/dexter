package store

type MemoryStore struct {
}

func (s MemoryStore) Name() string {
	return "Memory Store"
}
