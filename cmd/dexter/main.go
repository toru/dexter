package main

import (
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/pelletier/go-toml"
)

type config struct {
	// Temporary hack for development purpose. Eventually a more
	// sophisticated mechanism will be provided.
	Endpoints []string // Feed URLs to pull from
}

func main() {
	var cfgPath string
	flag.StringVar(&cfgPath, "cfg", "", "Path to the config file (required)")
	flag.Parse()

	if len(cfgPath) == 0 {
		flag.PrintDefaults()
		log.Fatal()
	}

	tree, err := toml.LoadFile(cfgPath)
	if err != nil {
		log.Fatal(err)
	}

	cfg := &config{}
	if err := tree.Unmarshal(cfg); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("loaded: %+v\n", cfg)

	fmt.Println("starting dexter")
	for range time.Tick(time.Second) {
		fmt.Printf("tick: %d\n", time.Now().Unix())
	}
}
