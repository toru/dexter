package main

import (
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/pelletier/go-toml"

	"github.com/toru/dexter/subscription"
)

const defaultSyncInterval string = "30m"

type webConfig struct {
	Listen string // TCP network address to listen on
	Port   uint   // TCP port to listen for Web API requests
}

type config struct {
	SyncInterval time.Duration `toml:"sync_interval"` // Interval between subscription syncs
	Web          webConfig     `toml:"web"`           // Web API server configuration

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

	cfgTree, err := toml.LoadFile(cfgPath)
	if err != nil {
		log.Fatal(err)
	}

	cfg := &config{}
	if err := cfgTree.Unmarshal(cfg); err != nil {
		log.Fatal(err)
	}

	if cfg.SyncInterval == 0 {
		log.Printf("SyncInterval missing, using: %s\n", defaultSyncInterval)
		cfg.SyncInterval, err = time.ParseDuration(defaultSyncInterval)
		if err != nil {
			log.Fatal(err)
		}
	}

	if cfgTree.Has("web") {
		// TODO(toru): Start the Web API server
	}

	// TODO(toru): This should be backed by a datastore whether it's on-memory
	// or disk. Write a simple inter-changeable storage mechanism.
	subscriptions := make([]subscription.Subscription, 0, len(cfg.Endpoints))
	for _, endpoint := range cfg.Endpoints {
		sub := subscription.New()
		sub.SetFeedURL(endpoint)
		subscriptions = append(subscriptions, *sub)
	}

	fmt.Printf("starting dexter with sync interval: %s\n", cfg.SyncInterval)
	for range time.Tick(cfg.SyncInterval) {
		log.Printf("tick: %d\n", time.Now().Unix())

		// TODO(toru): Concurrency
		for _, sub := range subscriptions {
			if err := sub.Sync(); err != nil {
				// Crash for dev-purpose
				log.Fatal(err)
			}
			log.Printf("sync'd: %s\n", sub.FeedURL.String())
		}
	}
}
