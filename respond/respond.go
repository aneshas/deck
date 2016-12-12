package respond

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/tonto/deck"
)

func Respond(w http.ResponseWriter, code int, data interface{}, err error) {
	var resp []byte

	if err != nil {
		e := deck.Error{
			Message:   err.Error(),
			Timestamp: time.Now(),
		}
		resp, _ = json.Marshal(e)
	} else {
		resp, _ = json.Marshal(data)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(resp)
}
