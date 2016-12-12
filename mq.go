package deck

type MQ interface {
	Publish([]Event)
}
