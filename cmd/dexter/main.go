package main

import (
	"flag"
	"log"
	"time"

	"github.com/pelletier/go-toml"

	"github.com/toru/dexter/index"
	"github.com/toru/dexter/storage"
	"github.com/toru/dexter/subscription"
	"github.com/toru/dexter/web"
)

const (
	defaultSyncInterval  = "30m"
	defaultHashAlgo      = "sha1"
	defaultStorageEngine = "memory"
)

type config struct {
	SyncInterval time.Duration  `toml:"sync_interval"` // Interval between subscription syncs
	HashAlgo     string         `toml:"hash_algo"`     // Hash algorithm for indexing
	Storage      storage.Config `toml:"storage"`       // Storage engine configuration
	Web          web.Config     `toml:"web"`           // Web API server configuration

	// Temporary hack for development purpose. Eventually a more
	// sophisticated mechanism will be provided.
	Endpoints []string // Feed URLs to pull from
}

func main() {
	var cfgPath string
	var verbose bool

	flag.StringVar(&cfgPath, "cfg", "", "Path to the config file (required)")
	flag.BoolVar(&verbose, "v", false, "Verbose output")
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
	if len(cfg.HashAlgo) == 0 {
		cfg.HashAlgo = defaultHashAlgo
	}
	if len(cfg.Storage.Engine) == 0 {
		cfg.Storage.Engine = defaultStorageEngine
	}
	if cfg.SyncInterval == 0 {
		if verbose {
			log.Printf("sync_interval missing, using: %s\n", defaultSyncInterval)
		}
		cfg.SyncInterval, err = time.ParseDuration(defaultSyncInterval)
		if err != nil {
			log.Fatal(err)
		}
	}

	db, err := storage.GetStore(cfg.Storage)
	if err != nil {
		log.Fatal(err)
	}

	// Temporary way to bootstrap subscriptions for dev purpose.
	for _, endpoint := range cfg.Endpoints {
		// TODO(toru): Use the index issuer.
		sub, err := subscription.New(endpoint)
		sub.ID = &index.SHA224DexID{}
		sub.ID.SetValueFromString(endpoint)

		if err != nil {
			log.Print(err)
		}
		if err = db.WriteSubscription(sub); err != nil {
			log.Fatal(err)
		}
	}

	if cfgTree.Has("web") {
		go web.ServeWebAPI(cfg.Web, db)
	}

	log.Printf("starting dexter with sync interval: %s\n", cfg.SyncInterval)
	for range time.Tick(cfg.SyncInterval) {
		log.Printf("tick: %d\n", time.Now().Unix())

		// TODO(toru): Concurrency
		for _, sub := range db.Subscriptions() {
			if sub.IsOffline() {
				log.Printf("skipping: %x", sub.ID)
				continue
			}
			log.Printf("syncing: %s\n", sub.FeedURL.String())
			dataFeed, err := sub.FeedSync()
			if err != nil {
				log.Print(err)
				continue
			}
			if err = db.WriteSubscription(&sub); err != nil {
				log.Print(err)
				continue
			}
			if err = db.WriteFeed(dataFeed); err != nil {
				log.Print(err)
				continue
			}
		}
	}
}
