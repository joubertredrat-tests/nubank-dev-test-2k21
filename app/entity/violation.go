package entity

const (
	ACCOUNT_VIOLATION_NOT_INITIALIZED                   = "account-not-initialized"
	ACCOUNT_VIOLATION_ALREADY_INITIALIZED               = "account-already-initialized"
	TRANSACTION_VIOLATION_CARD_NOT_ACTIVE               = "card-not-active"
	TRANSACTION_VIOLATION_INSUFFICIENT_LIMIT            = "insufficient-limit"
	TRANSACTION_VIOLATION_HIGH_FREQUENCY_SMALL_INTERVAL = "high-frequency-small-interval"
	TRANSACTION_VIOLATION_DOUBLE_TRANSACTION            = "double-transaction"
)

type Violation struct {
	name string
}

func NewViolationAccountNotInitialized() Violation {
	return Violation{
		name: ACCOUNT_VIOLATION_NOT_INITIALIZED,
	}
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

func NewViolationDoubleTransaction() Violation {
	return Violation{
		name: TRANSACTION_VIOLATION_DOUBLE_TRANSACTION,
	}
}

func (v Violation) GetName() string {
	return v.name
}
