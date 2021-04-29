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

type Account struct {
	ActiveCard     bool
	AvailableLimit Amount
	Violations     []AccountViolation
}

type AccountViolation struct {
	Name string
}

type Transaction struct {
	Merchant   string
	Amount     Amount
	Time       time.Time
	Violations []TransactionViolation
}

type TransactionViolation struct {
	Name string
}
