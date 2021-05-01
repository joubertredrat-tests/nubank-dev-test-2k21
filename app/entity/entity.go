package entity

type Event struct {
	Account    Account
	Violations []Violation
}

func NewEvent(account Account, violations []Violation) Event {
	return Event{
		Account:    account,
		Violations: violations,
	}
}

type Operations struct {
	Events []Event
}

func NewOperations() Operations {
	return Operations{
		Events: []Event{},
	}
}

func (o *Operations) RegisterEvent(account Account) {
	o.Events = append(
		o.Events,
		NewEvent(account, []Violation{}),
	)
}

func (o *Operations) RegisterViolationEvent(account Account, violation Violation) {
	o.Events = append(
		o.Events,
		NewEvent(account, []Violation{violation}),
	)
}

func (o *Operations) RegisterViolationsEvent(account Account, violations []Violation) {
	o.Events = append(
		o.Events,
		NewEvent(account, violations),
	)
}
