package deck

import (
	"errors"
	"fmt"
	"time"
)

type Error struct {
	Message   string    `json:"message"`
	Timestamp time.Time `json:"timestamp"`
}

var ParseCommandError = errors.New("Error parsing command.")
var UnrecognizedCommandError = errors.New("Command not recognized.")
var CommandNotSetError = errors.New("Command not set to context.")
var AggregateNotSetError = errors.New("Aggregate not set to context.")

func NewCommandValidationError(e string) error {
	return errors.New(fmt.Sprintf("Command failed to validate! REASON: %s", e))
}
