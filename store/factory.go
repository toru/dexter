package store

import (
	"log"
)

type Store interface {
	Name() string
}

func GetStore(name string) (Store, bool) {
	switch name {
	case "memory":
		return MemoryStore{}, true
	case "mysql", "mariadb":
		log.Println("work in progress")
	default:
		log.Println("unknown store")
	}
	return nil, false
}
