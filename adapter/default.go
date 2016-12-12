package adapter

import (
	"net/http"
	"reflect"

	"github.com/tonto/deck"
)

func WithDefaultStack(tCmd, tAggr reflect.Type, store deck.EventStore, mq deck.MQ) deck.Adapter {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// TODO - Persist and publish events attomically
			h.ServeHTTP(w, r)
		})
	}
}
