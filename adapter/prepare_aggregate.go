package adapter

import (
	"log"
	"net/http"
	"reflect"

	"github.com/gorilla/context"
	"github.com/tonto/deck"
)

// WithPrepareAggregate takes event store interface and aggregate type
// It will instantiate the aggregate, roll previous events over it and set it to context
func WithPrepareAggregate(aggregate interface{}, store deck.EventStore) deck.Adapter {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			aggrPtr := reflect.New(reflect.TypeOf(aggregate))
			aggr := aggrPtr.Interface().(deck.Aggregate)

			log.Printf("Seeding aggregate: %#v", aggr)
			aggr.Seed()

			// TODO
			// 1 - Get aggregate id from cmd
			// 2 - Read past events from event store and apply them to aggregate

			context.Set(r, deck.AGGREGATE_KEY, aggr)

			h.ServeHTTP(w, r)
		})
	}
}
