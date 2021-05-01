package validator

import (
	"time"

	"dev-test/nubank-dev-test-2k21/app/dto/input"
	"dev-test/nubank-dev-test-2k21/app/entity"
)

type HighTransactionsValidator struct {
	TransactionsTime     []time.Time
	TransactionsAnalysis uint
	TimeIntervalSeconds  uint
}

func NewHighTransactionsValidator(transactionsAnalysis, timeIntervalSeconds uint) ValidatorInterface {
	return &HighTransactionsValidator{
		TransactionsAnalysis: transactionsAnalysis,
		TimeIntervalSeconds:  timeIntervalSeconds,
	}
}

func (v *HighTransactionsValidator) GetViolation(account entity.Account, transactionLine input.TransactionLine) *entity.Violation {
	v.registerTransactionTime(transactionLine.Transaction.Time)
	if v.hasHighFrequency() {
		violation := entity.NewViolationHighFrequencySmallInterval()
		return &violation
	}

	return nil
}

func (v *HighTransactionsValidator) registerTransactionTime(t time.Time) {
	v.TransactionsTime = append(v.TransactionsTime, t)
}

func (v *HighTransactionsValidator) hasHighFrequency() bool {
	transactionsTime := v.getTransactionsTimeForAnalysis()

	if len(transactionsTime) < int(v.TransactionsAnalysis-1) {
		return false
	}

	timeInitial := transactionsTime[0]
	timeFinal := transactionsTime[len(transactionsTime)-1]
	timeDiff := timeFinal.Sub(timeInitial)

	return float64(v.TimeIntervalSeconds) >= timeDiff.Seconds()
}

func (v *HighTransactionsValidator) getTransactionsTimeForAnalysis() []time.Time {
	if len(v.TransactionsTime) <= int(v.TransactionsAnalysis) {
		return v.TransactionsTime
	}

	sliceCut := len(v.TransactionsTime) - int(v.TransactionsAnalysis)
	return v.TransactionsTime[sliceCut:]
}
