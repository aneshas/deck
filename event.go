package deck

import "time"

type Event interface{}

type BaseEvent struct {
	AggregateID string
	Timestamp   time.Time
	Version     int
	ProcessID   string
}
