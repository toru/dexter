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
	case "mysql", "mariadb":
	default:
		log.Println("unknown store")
	}
	return nil, false
}
