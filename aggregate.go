package deck

type Aggregate interface {
	Apply(Event)
	SetUncommited([]Event)
	ApplyUncommited()
	Seed()
}
