package web

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/toru/dexter/index"
	"github.com/toru/dexter/store"
	"github.com/toru/dexter/subscription"
)

const defaultPort = 8081

// ServerConfig holds the Web API server settings.
type ServerConfig struct {
	Listen string // TCP network address to listen on
	Port   uint   // TCP port to listen for Web API requests
}

type entryPresenter struct {
	ID      string `json:"id"`      // Entry ID
	Summary string `json:"summary"` // Entry Summary
}

type feedPresenter struct {
	ID             string `json:"id"`              // Feed ID
	SubscriptionID string `json:"subscription_id"` // Subscription ID
	Title          string `json:"title"`           // Feed Title
}

type subscriptionPresenter struct {
	ID  string `json:"id"`  // Hex representation of the ID
	URL string `json:"url"` // FeedURL as a string
}

func splitPath(path string) []string {
	return strings.FieldsFunc(path, func(c rune) bool {
		return c == '/'
	})
}

func render400(w http.ResponseWriter, reason string) {
	http.Error(w, strconv.Quote(reason), http.StatusBadRequest)
}

func render404(w http.ResponseWriter) {
	http.Error(w, strconv.Quote("not found"), http.StatusNotFound)
}

// Entry point for the /feeds resource.
func feedsResourceHandlerFunc(db store.Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tokens := splitPath(r.URL.Path)
		numTokens := len(tokens)

		if r.Method != http.MethodGet || numTokens > 3 {
			render404(w)
			return
		}
		if numTokens > 1 {
			if len(tokens[1]) != index.DexHexIDLen {
				render404(w)
				return
			}
		}

		if numTokens == 1 {
			getFeedsHandler(db, w, r)
		} else if numTokens == 2 {
			getFeedHandler(db, tokens[1], w, r)
		} else {
			if tokens[2] != "entries" {
				render404(w)
				return
			}
			getFeedEntriesHandler(db, tokens[1], w, r)
		}
	}
}

// Entry point for the /subscriptions resource.
func subscriptionsResourceHandlerFunc(db store.Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			postSubscriptionsHandler(db, w, r)
		case http.MethodGet:
			getSubscriptionsHandler(db, w, r)
		default:
			render404(w)
		}
	}
}

// GET /feeds/:id/entries
// Renders a list of entries associated to the given feed ID.
func getFeedEntriesHandler(db store.Store, feedID string, w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("unimplemented"))
}

// GET /feeds/:id
// Renders a feed that was found with the given ID.
func getFeedHandler(db store.Store, id string, w http.ResponseWriter, r *http.Request) {
	xid, err := index.NewDexIDFromHexDigest(id)
	if err != nil {
		render400(w, "invalid feed id")
		return
	}
	f, ok := db.Feed(xid)
	if !ok {
		render404(w)
		return
	}
	subID := index.DexIDToHexDigest(f.SubscriptionID())
	rv := feedPresenter{f.ID(), subID, f.Title()}
	buf, err := json.Marshal(rv)
	if err != nil {
		log.Print(err)
		http.Error(w, strconv.Quote("payload generation"),
			http.StatusInternalServerError)
		return
	}
	w.Write(buf)
}

// GET /feeds
// Renders a list of feeds.
func getFeedsHandler(db store.Store, w http.ResponseWriter, r *http.Request) {
	feeds := make([]feedPresenter, 0, db.NumSubscriptions())
	for _, f := range db.Feeds() {
		rawSubID := f.SubscriptionID()
		feeds = append(feeds, feedPresenter{
			f.ID(),
			index.DexIDToHexDigest(rawSubID),
			f.Title(),
		})
	}

	buf, err := json.Marshal(feeds)
	if err != nil {
		log.Print(err)
		http.Error(w, strconv.Quote("payload generation"),
			http.StatusInternalServerError)
		return
	}
	w.Write(buf)
}

// POST /subscriptions
// Creates a new subscription, given a "url" parameter.
func postSubscriptionsHandler(db store.Store, w http.ResponseWriter, r *http.Request) {
	feedURL := r.PostFormValue("url")
	if len(feedURL) == 0 {
		render400(w, "url parameter missing")
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
			index.DexIDToHexDigest(sub.ID),
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

// ServeWebAPI starts the Web API application.
func ServeWebAPI(cfg ServerConfig, db store.Store) error {
	log.Println("starting the web api server")
	if cfg.Port == 0 {
		log.Printf("port missing, using: %d", defaultPort)
		cfg.Port = defaultPort
	}

	// TODO(toru): Duplicate route definition just to work around the
	// trailing slash enforcement is silly. Implement our own ServeHTTP().
	http.Handle("/feeds", feedsResourceHandlerFunc(db))
	http.Handle("/feeds/", feedsResourceHandlerFunc(db))
	http.Handle("/subscriptions", subscriptionsResourceHandlerFunc(db))

	// TODO(toru): TLS
	addr := fmt.Sprintf("%s:%d", cfg.Listen, cfg.Port)
	log.Fatal(http.ListenAndServe(addr, nil))

	return nil
}
