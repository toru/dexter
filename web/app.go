package web

import (
	"log"
)

const defaultPort uint = 8084

type ServerConfig struct {
	Listen string // TCP network address to listen on
	Port   uint   // TCP port to listen for Web API requests
}

func ServeWebAPI(cfg ServerConfig) error {
	log.Println("starting the web api server")
	if cfg.Port == 0 {
		log.Printf("port missing, using: %d", defaultPort)
		cfg.Port = defaultPort
	}

	// TODO(toru): Start the actual http server. Always favor TLS.
	return nil
}
