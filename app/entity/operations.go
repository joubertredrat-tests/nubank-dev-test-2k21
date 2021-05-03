package entity

type Operations struct {
	events []OperationEvent
}

func NewOperations() Operations {
	return Operations{
		events: []OperationEvent{},
	}
}

func (o *Operations) RegisterEvent(account Account, violations []*Violation) {
	o.events = append(
		o.events,
		NewOperationEvent(account, []*Violation{}),
	)
}

func (o *Operations) RegisterViolationEvent(account Account, violation *Violation) {
	o.events = append(
		o.events,
		NewOperationEvent(account, []*Violation{violation}),
	)
}
