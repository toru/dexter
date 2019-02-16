package web

import (
	"fmt"
	"log"
	"net/http"

	"github.com/toru/dexter/store"
)

const defaultPort uint = 8084

type ServerConfig struct {
	Listen string // TCP network address to listen on
	Port   uint   // TCP port to listen for Web API requests
}

// ServeWebAPI starts the Web API application.
func ServeWebAPI(cfg ServerConfig, db store.Store) error {
	log.Println("starting the web api server")
	if cfg.Port == 0 {
		log.Printf("port missing, using: %d", defaultPort)
		cfg.Port = defaultPort
	}

	// TODO(toru): TLS
	addr := fmt.Sprintf("%s:%d", cfg.Listen, cfg.Port)
	log.Fatal(http.ListenAndServe(addr, nil))

	return nil
}
