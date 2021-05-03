package entity

type OperationEvent struct {
	account    Account
	violations []*Violation
}

func NewOperationEvent(account Account, violations []*Violation) OperationEvent {
	return OperationEvent{
		account:    account,
		violations: violations,
	}
}

func (o OperationEvent) GetAccount() Account {
	return o.account
}

func (o OperationEvent) GetViolations() []*Violation {
	return o.violations
}

func (o OperationEvent) HasViolations() bool {
	return len(o.violations) > 0
}
