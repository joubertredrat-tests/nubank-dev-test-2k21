package validator

import (
	"time"

	"dev-test/nubank-dev-test-2k21/app/entity"
)

type HighTransactionsValidator struct {
	transactionsTime     []time.Time
	transactionsAnalysis uint
	timeIntervalSeconds  uint
}

func NewHighTransactionsValidator(transactionsAnalysis, timeIntervalSeconds uint) *HighTransactionsValidator {
	return &HighTransactionsValidator{
		transactionsAnalysis: transactionsAnalysis,
		timeIntervalSeconds:  timeIntervalSeconds,
	}
}

func (v *HighTransactionsValidator) GetViolation(account entity.Account, transaction entity.Transaction) *entity.Violation {
	v.registerTransactionTime(transaction.GetTime())
	if v.hasHighFrequency() {
		return entity.NewViolationHighFrequencySmallInterval()
	}

	return nil
}

func (v *HighTransactionsValidator) registerTransactionTime(t time.Time) {
	v.transactionsTime = append(v.transactionsTime, t)
}

func (v *HighTransactionsValidator) hasHighFrequency() bool {
	transactionsTime := v.getTransactionsTimeForAnalysis()

	if int(v.transactionsAnalysis) > len(transactionsTime) {
		return false
	}

	timeInitial := transactionsTime[0]
	timeFinal := transactionsTime[len(transactionsTime)-1]
	timeDiff := timeFinal.Sub(timeInitial)

	return float64(v.timeIntervalSeconds) >= timeDiff.Seconds()
}

func (v *HighTransactionsValidator) getTransactionsTimeForAnalysis() []time.Time {
	if int(v.transactionsAnalysis) >= len(v.transactionsTime) {
		return v.transactionsTime
	}

	sliceCut := len(v.transactionsTime) - int(v.transactionsAnalysis)
	return v.transactionsTime[sliceCut:]
}
