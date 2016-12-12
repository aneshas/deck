package adapter

import (
	"encoding/json"
	"log"
	"net/http"
	"reflect"
	"time"

	"github.com/gorilla/context"
	"github.com/tonto/deck"
	"github.com/tonto/deck/respond"
)

// WithCommandHandler executes provided http handler func with user defined
// custom code, prior to executing the actual command and commiting evts to aggregate
func WithCommandHandler(handlerFunc http.HandlerFunc) deck.Adapter {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			handlerFunc(w, r)

			cmd := context.Get(r, deck.COMMAND_KEY)
			command, ok := cmd.(deck.Command)
			if !ok {
				respond.Respond(w, http.StatusInternalServerError, nil, deck.AggregateNotSetError)
				return
			}

			aggr := context.Get(r, deck.AGGREGATE_KEY)
			aggregate, ok := aggr.(deck.Aggregate)
			if !ok {
				respond.Respond(w, http.StatusInternalServerError, nil, deck.CommandNotSetError)
				return
			}

			log.Printf("Executing command: %#v on aggregate %#v", cmd, aggr)

			newEvts, err := command.Execute()
			if err != nil {
				respond.Respond(w, http.StatusOK, nil, err)
				return
			}

			aggregate.SetUncommited(newEvts)
			// TODO - Should we apply uncommited changes here ???
			aggregate.ApplyUncommited()

			log.Printf("Command executed successfully: %#v", cmd)

			h.ServeHTTP(w, r)
		})
	}
}

// WithCommand creates a new adapter which instantiates
// the given command, decodes it from json and sets it to context
func WithCommand(command interface{}) deck.Adapter {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			log.Printf("-------------")
			log.Printf("New incomming request %s", time.Now())

			cmdPtr := reflect.New(reflect.TypeOf(command))
			cmd := cmdPtr.Interface().(deck.Command)

			err := json.NewDecoder(r.Body).Decode(&cmd)
			if err != nil {
				log.Printf("Could not json decode new command: %#v, ERROR: %v", cmd, err)
				respond.Respond(w, http.StatusInternalServerError, nil, err)
				return
			}

			context.Set(r, deck.COMMAND_KEY, cmd)

			log.Printf("Instantiating new command: %#v", cmd)

			h.ServeHTTP(w, r)
		})
	}
}
