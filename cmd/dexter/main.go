package main

import (
	"flag"
	"log"
	"time"

	"github.com/pelletier/go-toml"

	"github.com/toru/dexter/store"
	"github.com/toru/dexter/subscription"
	"github.com/toru/dexter/web"
)

const defaultSyncInterval string = "30m"

type config struct {
	SyncInterval time.Duration    `toml:"sync_interval"` // Interval between subscription syncs
	Web          web.ServerConfig `toml:"web"`           // Web API server configuration

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
		log.Printf("sync_interval missing, using: %s\n", defaultSyncInterval)
		cfg.SyncInterval, err = time.ParseDuration(defaultSyncInterval)
		if err != nil {
			log.Fatal(err)
		}
	}

	db, err := store.GetStore("memory")
	if err != nil {
		log.Fatal(err)
	}

	// Temporary way to bootstrap subscriptions for dev purpose.
	for _, endpoint := range cfg.Endpoints {
		sub, err := subscription.New(endpoint)
		if err != nil {
			log.Error(err)
		}
		if err = db.WriteSubscription(sub); err != nil {
			log.Fatal(err)
		}
	}

	if cfgTree.Has("web") {
		web.ServeWebAPI(cfg.Web)
	}

	log.Printf("starting dexter with sync interval: %s\n", cfg.SyncInterval)
	for range time.Tick(cfg.SyncInterval) {
		log.Printf("tick: %d\n", time.Now().Unix())

		// TODO(toru): Concurrency
		for _, sub := range db.Subscriptions() {
			if err := sub.Sync(); err != nil {
				// Crash for dev-purpose
				log.Fatal(err)
			}
			log.Printf("sync'd: %s\n", sub.FeedURL.String())
		}
	}
}
