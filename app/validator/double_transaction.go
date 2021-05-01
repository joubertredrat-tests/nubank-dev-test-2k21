package validator

import (
	"dev-test/nubank-dev-test-2k21/app/entity"
)

type DoubleTransactionValidator struct {
	Transactions        map[string][]entity.Transaction
	TimeIntervalSeconds uint
}

func NewDoubleTransactionValidator(timeIntervalSeconds uint) *DoubleTransactionValidator {
	return &DoubleTransactionValidator{
		Transactions:        make(map[string][]entity.Transaction),
		TimeIntervalSeconds: timeIntervalSeconds,
	}
}

func (v *DoubleTransactionValidator) IsAccountValidator() bool {
	return false
}

func (v *DoubleTransactionValidator) IsTransactionValidator() bool {
	return true
}

func (v *DoubleTransactionValidator) IsOperationValidator() bool {
	return false
}

func (v *DoubleTransactionValidator) GetViolation(transaction entity.Transaction) *entity.Violation {
	v.registerTransaction(transaction)
	if v.hasDoubleTransactions(transaction) {
		violation := entity.NewViolationDoubleTransaction()
		return &violation
	}

	return nil
}

func (v *DoubleTransactionValidator) registerTransaction(transaction entity.Transaction) {
	if _, ok := v.Transactions[transaction.Merchant]; !ok {
		v.Transactions[transaction.Merchant] = []entity.Transaction{}
	}

	v.Transactions[transaction.Merchant] = append(
		v.Transactions[transaction.Merchant],
		transaction,
	)
}

func (v *DoubleTransactionValidator) hasDoubleTransactions(transaction entity.Transaction) bool {
	if len(v.Transactions[transaction.Merchant]) < 2 {
		return false
	}

	keyTransactionInitial := len(v.Transactions[transaction.Merchant]) - 2
	keyTransactionFinal := len(v.Transactions[transaction.Merchant]) - 1
	transactionInitial := v.Transactions[transaction.Merchant][keyTransactionInitial]
	transactionFinal := v.Transactions[transaction.Merchant][keyTransactionFinal]

	if transactionFinal.Amount.Value != transactionInitial.Amount.Value {
		return false
	}

	timeDiff := transactionFinal.Time.Sub(transactionInitial.Time)
	return float64(v.TimeIntervalSeconds) >= timeDiff.Seconds()
}
