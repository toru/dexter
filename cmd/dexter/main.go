package main

import (
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/pelletier/go-toml"

	"github.com/toru/dexter/subscription"
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

	// TODO(toru): This should be backed by a datastore whether it's on-memory
	// or disk. Write a simple inter-changeable storage mechanism.
	subscriptions := make([]subscription.Subscription, 0, len(cfg.Endpoints))
	for _, endpoint := range cfg.Endpoints {
		sub := subscription.New()
		sub.SetFeedURL(endpoint)
		subscriptions = append(subscriptions, *sub)
	}

	fmt.Println("starting dexter")
	for range time.Tick(time.Second) {
		log.Printf("tick: %d\n", time.Now().Unix())
	}
}
