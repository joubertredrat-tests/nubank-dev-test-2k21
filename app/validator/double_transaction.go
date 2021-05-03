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

func (v *DoubleTransactionValidator) IsBreakNextCheck() bool {
	return false
}

func (v *DoubleTransactionValidator) GetViolation(account entity.Account, transaction entity.Transaction) *entity.Violation {
	v.registerTransaction(transaction)
	if v.hasDoubleTransactions(transaction) {
		return entity.NewViolationDoubleTransaction()
	}

	return nil
}

func (v *DoubleTransactionValidator) registerTransaction(transaction entity.Transaction) {
	if _, ok := v.Transactions[transaction.GetMerchant()]; !ok {
		v.Transactions[transaction.GetMerchant()] = []entity.Transaction{}
	}

	v.Transactions[transaction.GetMerchant()] = append(
		v.Transactions[transaction.GetMerchant()],
		transaction,
	)
}

func (v *DoubleTransactionValidator) hasDoubleTransactions(transaction entity.Transaction) bool {
	if len(v.Transactions[transaction.GetMerchant()]) < 2 {
		return false
	}

	keyTransactionInitial := len(v.Transactions[transaction.GetMerchant()]) - 2
	keyTransactionFinal := len(v.Transactions[transaction.GetMerchant()]) - 1
	transactionInitial := v.Transactions[transaction.GetMerchant()][keyTransactionInitial]
	transactionFinal := v.Transactions[transaction.GetMerchant()][keyTransactionFinal]

	if transactionFinal.GetAmount().GetValue() != transactionInitial.GetAmount().GetValue() {
		return false
	}

	timeDiff := transactionFinal.GetTime().Sub(transactionInitial.GetTime())
	return float64(v.TimeIntervalSeconds) >= timeDiff.Seconds()
}
