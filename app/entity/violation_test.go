package entity_test

import (
	"dev-test/nubank-dev-test-2k21/app/entity"
	"reflect"
	"testing"
)

func TestViolation(t *testing.T) {
	tests := []struct {
		name                  string
		violationNameExpected string
		violationGot          *entity.Violation
	}{
		{
			name:                  "Test with account violation not initialized",
			violationNameExpected: entity.ACCOUNT_VIOLATION_NOT_INITIALIZED,
			violationGot:          entity.NewViolationAccountNotInitialized(),
		},
		{
			name:                  "Test with account violation already initialized",
			violationNameExpected: entity.ACCOUNT_VIOLATION_ALREADY_INITIALIZED,
			violationGot:          entity.NewViolationAccountAlreadyInitialized(),
		},
		{
			name:                  "Test with transaction violation card not active",
			violationNameExpected: entity.TRANSACTION_VIOLATION_CARD_NOT_ACTIVE,
			violationGot:          entity.NewViolationCardNotActive(),
		},
		{
			name:                  "Test with transaction violation insufficient limit",
			violationNameExpected: entity.TRANSACTION_VIOLATION_INSUFFICIENT_LIMIT,
			violationGot:          entity.NewViolationInsufficientLimit(),
		},
		{
			name:                  "Test with transaction violation high frequency small interval",
			violationNameExpected: entity.TRANSACTION_VIOLATION_HIGH_FREQUENCY_SMALL_INTERVAL,
			violationGot:          entity.NewViolationHighFrequencySmallInterval(),
		},
		{
			name:                  "Test with transaction violation double transaction",
			violationNameExpected: entity.TRANSACTION_VIOLATION_DOUBLE_TRANSACTION,
			violationGot:          entity.NewViolationDoubleTransaction(),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if test.violationNameExpected != test.violationGot.GetName() {
				t.Errorf("%s violation.GetName() expected %v, got %v", test.name, test.violationNameExpected, test.violationGot.GetName())
			}
		})
	}
}

func TestNewViolationsEmpty(t *testing.T) {
	violationsExpected := []*entity.Violation{}
	violationsGot := entity.NewViolationsEmpty()

	if !reflect.DeepEqual(violationsExpected, violationsGot) {
		t.Errorf("entity.NewViolationsEmpty() expected %#v, got %#v", violationsExpected, violationsGot)
	}
}
