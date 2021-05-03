package validator

import (
	"dev-test/nubank-dev-test-2k21/app/entity"
)

type DoubleTransactionValidator struct {
	transactions        map[string][]entity.Transaction
	timeIntervalSeconds uint
}

func NewDoubleTransactionValidator(timeIntervalSeconds uint) *DoubleTransactionValidator {
	return &DoubleTransactionValidator{
		transactions:        make(map[string][]entity.Transaction),
		timeIntervalSeconds: timeIntervalSeconds,
	}
}

func (v *DoubleTransactionValidator) GetViolation(account entity.Account, transaction entity.Transaction) *entity.Violation {
	v.registerTransaction(transaction)
	if v.hasDoubleTransactions(transaction) {
		return entity.NewViolationDoubleTransaction()
	}

	return nil
}

func (v *DoubleTransactionValidator) registerTransaction(transaction entity.Transaction) {
	if _, ok := v.transactions[transaction.GetMerchant()]; !ok {
		v.transactions[transaction.GetMerchant()] = []entity.Transaction{}
	}

	v.transactions[transaction.GetMerchant()] = append(
		v.transactions[transaction.GetMerchant()],
		transaction,
	)
}

func (v *DoubleTransactionValidator) hasDoubleTransactions(transaction entity.Transaction) bool {
	if len(v.transactions[transaction.GetMerchant()]) < 2 {
		return false
	}

	keyTransactionInitial := len(v.transactions[transaction.GetMerchant()]) - 2
	keyTransactionFinal := len(v.transactions[transaction.GetMerchant()]) - 1
	transactionInitial := v.transactions[transaction.GetMerchant()][keyTransactionInitial]
	transactionFinal := v.transactions[transaction.GetMerchant()][keyTransactionFinal]

	if transactionFinal.GetAmount().GetValue() != transactionInitial.GetAmount().GetValue() {
		return false
	}

	timeDiff := transactionFinal.GetTime().Sub(transactionInitial.GetTime())
	return float64(v.timeIntervalSeconds) >= timeDiff.Seconds()
}
