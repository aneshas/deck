package adapter

import (
	"log"
	"net/http"

	"github.com/gorilla/context"
	"github.com/tonto/deck"
	"github.com/tonto/deck/respond"
)

// WithSoftValidation validates the given command
// based on command internal rules
func WithSoftValidation() deck.Adapter {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			cmd := context.Get(r, deck.COMMAND_KEY)

			command, ok := cmd.(deck.Command)
			if ok {
				err := command.Validate()
				if err != nil {
					log.Printf("Command not valid: %#v, ERROR: %v", cmd, err)
					respond.Respond(w, http.StatusBadRequest, nil, err)
					return
				}
			}

			log.Printf("Command validated: %#v", cmd)

			h.ServeHTTP(w, r)
		})
	}
}
