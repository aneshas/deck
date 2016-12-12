package adapter

import (
	"log"
	"net/http"
	"time"

	"github.com/tonto/deck"
)

func WithPersistAndPublish(store deck.EventStore, mq deck.MQ) deck.Adapter {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// TODO - Persist and publish events attomically

			log.Printf("Done serving request at %s", time.Now())
			log.Printf("-------------")

			h.ServeHTTP(w, r)
		})
	}
}
