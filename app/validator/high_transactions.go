package validator

import (
	"time"

	"dev-test/nubank-dev-test-2k21/app/entity"
)

type HighTransactionsValidator struct {
	TransactionsTime     []time.Time
	TransactionsAnalysis uint
	TimeIntervalSeconds  uint
}

func NewHighTransactionsValidator(transactionsAnalysis, timeIntervalSeconds uint) *HighTransactionsValidator {
	return &HighTransactionsValidator{
		TransactionsAnalysis: transactionsAnalysis,
		TimeIntervalSeconds:  timeIntervalSeconds,
	}
}

func (v *HighTransactionsValidator) IsAccountValidator() bool {
	return false
}

func (v *HighTransactionsValidator) IsTransactionValidator() bool {
	return true
}

func (v *HighTransactionsValidator) IsOperationValidator() bool {
	return false
}

func (v *HighTransactionsValidator) GetViolation(transaction entity.Transaction) *entity.Violation {
	v.registerTransactionTime(transaction.Time)
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
