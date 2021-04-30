package entity

import "time"

const (
	ACCOUNT_VIOLATION_ALREADY_INITIALIZED               = "account-already-initialized"
	TRANSACTION_VIOLATION_CARD_NOT_ACTIVE               = "card-not-active"
	TRANSACTION_VIOLATION_INSUFFICIENT_LIMIT            = "insufficient-limit"
	TRANSACTION_VIOLATION_HIGH_FREQUENCY_SMALL_INTERVAL = "high-frequency-small-interval"
	TRANSACTION_VIOLATION_DOUBLE_TRANSACTION            = "double-transaction"
)

type Amount struct {
	Value uint
}

func NewAmount(value uint) Amount {
	return Amount{
		Value: value,
	}
}

type Account struct {
	ActiveCard     bool
	AvailableLimit Amount
}

func NewAccount(activeCard bool, availableLimit uint) Account {
	return Account{
		ActiveCard:     activeCard,
		AvailableLimit: NewAmount(availableLimit),
	}
}

type Transaction struct {
	Merchant string
	Amount   Amount
	Time     time.Time
}

func NewTransaction(merchant string, amount uint, time time.Time) Transaction {
	return Transaction{
		Merchant: merchant,
		Amount:   NewAmount(amount),
		Time:     time,
	}
}

type Violation struct {
	name string
}

func (v Violation) GetName() string {
	return v.name
}

func NewViolationAccountAlreadyInitialized() Violation {
	return Violation{
		name: ACCOUNT_VIOLATION_ALREADY_INITIALIZED,
	}
}

func NewViolationCardNotActive() Violation {
	return Violation{
		name: TRANSACTION_VIOLATION_CARD_NOT_ACTIVE,
	}
}

func NewViolationInsufficientLimit() Violation {
	return Violation{
		name: TRANSACTION_VIOLATION_INSUFFICIENT_LIMIT,
	}
}

func NewViolationHighFrequencySmallInterval() Violation {
	return Violation{
		name: TRANSACTION_VIOLATION_HIGH_FREQUENCY_SMALL_INTERVAL,
	}
}

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

func (o *Operations) RegisterViolationEvent(account Account, violations []Violation) {
	o.Events = append(
		o.Events,
		NewEvent(account, violations),
	)
}
