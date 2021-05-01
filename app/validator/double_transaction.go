package validator

import (
	"dev-test/nubank-dev-test-2k21/app/dto/input"
	"dev-test/nubank-dev-test-2k21/app/entity"
)

const INTERVAL_SECONDS = 120

type DoubleTransactionValidator struct {
	Transactions        map[string][]entity.Transaction
	TimeIntervalSeconds uint
}

func NewDoubleTransactionValidator(timeIntervalSeconds uint) ValidatorInterface {
	return &DoubleTransactionValidator{
		TimeIntervalSeconds: timeIntervalSeconds,
	}
}

func (v *DoubleTransactionValidator) GetViolation(account entity.Account, transactionLine input.TransactionLine) *entity.Violation {
	v.registerTransaction(transactionLine)
	if v.hasDoubleTransactions(transactionLine) {
		violation := entity.NewViolationDoubleTransaction()
		return &violation
	}

	return nil
}

func (v *DoubleTransactionValidator) registerTransaction(transactionLine input.TransactionLine) {
	if _, ok := v.Transactions[transactionLine.Transaction.Merchant]; !ok {
		v.Transactions[transactionLine.Transaction.Merchant] = []entity.Transaction{}
	}

	v.Transactions[transactionLine.Transaction.Merchant] = append(
		v.Transactions[transactionLine.Transaction.Merchant],
		entity.NewTransaction(
			transactionLine.Transaction.Merchant,
			transactionLine.Transaction.Amount,
			transactionLine.Transaction.Time,
		),
	)
}

func (v *DoubleTransactionValidator) hasDoubleTransactions(transactionLine input.TransactionLine) bool {
	if len(v.Transactions[transactionLine.Transaction.Merchant]) < 2 {
		return false
	}

	keyTransactionInitial := len(v.Transactions[transactionLine.Transaction.Merchant]) - 2
	keyTransactionFinal := len(v.Transactions[transactionLine.Transaction.Merchant]) - 1
	transactionInitial := v.Transactions[transactionLine.Transaction.Merchant][keyTransactionInitial]
	transactionFinal := v.Transactions[transactionLine.Transaction.Merchant][keyTransactionFinal]

	if transactionFinal.Amount.Value != transactionInitial.Amount.Value {
		return false
	}

	timeDiff := transactionFinal.Time.Sub(transactionInitial.Time)

	return INTERVAL_SECONDS >= timeDiff.Seconds()
}
