package deck

type Command interface {
	Execute() ([]Event, error)
	Validate() error
	GetAggregateID() string
}
