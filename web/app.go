package web

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/toru/dexter/store"
	"github.com/toru/dexter/subscription"
)

const defaultPort uint = 8084

type ServerConfig struct {
	Listen string // TCP network address to listen on
	Port   uint   // TCP port to listen for Web API requests
}

type subscriptionPresenter struct {
	ID  string `json:"id"`  // Hex representation of the ID
	URL string `json:"url"` // FeedURL as a string
}

// POST /subscriptions
// Creates a new subscription, given a "url" parameter.
func postSubscriptionsHandler(db store.Store, w http.ResponseWriter, r *http.Request) {
	feedURL := r.PostFormValue("url")
	if len(feedURL) == 0 {
		http.Error(w, strconv.Quote("url parameter missing"),
			http.StatusBadRequest)
		return
	}

	sub, err := subscription.New(feedURL)
	if err != nil {
		http.Error(w, strconv.Quote(err.Error()),
			http.StatusInternalServerError)
		return
	}
	if err := db.WriteSubscription(sub); err != nil {
		http.Error(w, strconv.Quote(err.Error()),
			http.StatusInternalServerError)
		return
	}
	w.Write([]byte(strconv.Quote("ok")))
}

// GET /subscriptions
// Renders a list of subscriptions.
func getSubscriptionsHandler(db store.Store, w http.ResponseWriter, r *http.Request) {
	subs := make([]subscriptionPresenter, 0, db.NumSubscriptions())
	for _, sub := range db.Subscriptions() {
		subs = append(subs, subscriptionPresenter{
			hex.EncodeToString(sub.ID[:]),
			sub.FeedURL.String(),
		})
	}

	buf, err := json.Marshal(subs)
	if err != nil {
		log.Print(err)
		http.Error(w, strconv.Quote("payload generation"),
			http.StatusInternalServerError)
		return
	}
	w.Write(buf)
}

// Entry point for the /subscriptions resource.
func subscriptionsHandlerFunc(db store.Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			postSubscriptionsHandler(db, w, r)
		case http.MethodGet:
			getSubscriptionsHandler(db, w, r)
		default:
			http.Error(w, strconv.Quote("not found"),
				http.StatusNotFound)
		}
	}
}

// ServeWebAPI starts the Web API application.
func ServeWebAPI(cfg ServerConfig, db store.Store) error {
	log.Println("starting the web api server")
	if cfg.Port == 0 {
		log.Printf("port missing, using: %d", defaultPort)
		cfg.Port = defaultPort
	}

	http.Handle("/subscriptions", subscriptionsHandlerFunc(db))

	// TODO(toru): TLS
	addr := fmt.Sprintf("%s:%d", cfg.Listen, cfg.Port)
	log.Fatal(http.ListenAndServe(addr, nil))

	return nil
}
