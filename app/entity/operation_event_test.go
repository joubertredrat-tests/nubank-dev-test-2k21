package entity_test

import (
	"reflect"
	"testing"

	"dev-test/nubank-dev-test-2k21/app/entity"
)

func TestOperationEvent(t *testing.T) {
	tests := []struct {
		name                  string
		accountExpected       entity.Account
		violationsExpected    []*entity.Violation
		hasViolationsExpected bool
		operationEventGot     entity.OperationEvent
	}{
		{
			name:                  "Test operation event with no violations",
			accountExpected:       entity.NewAccount(true, 120),
			violationsExpected:    entity.NewViolationsEmpty(),
			hasViolationsExpected: false,
			operationEventGot: entity.NewOperationEvent(
				entity.NewAccount(true, 120),
				entity.NewViolationsEmpty(),
			),
		},
		{
			name:            "Test operation event with one violation",
			accountExpected: entity.NewAccount(true, 120),
			violationsExpected: []*entity.Violation{
				entity.NewViolationAccountNotInitialized(),
			},
			hasViolationsExpected: true,
			operationEventGot: entity.NewOperationEvent(
				entity.NewAccount(true, 120),
				[]*entity.Violation{
					entity.NewViolationAccountNotInitialized(),
				},
			),
		},
		{
			name:            "Test operation event with multiple violations",
			accountExpected: entity.NewAccount(true, 120),
			violationsExpected: []*entity.Violation{
				entity.NewViolationInsufficientLimit(),
				entity.NewViolationDoubleTransaction(),
			},
			hasViolationsExpected: true,
			operationEventGot: entity.NewOperationEvent(
				entity.NewAccount(true, 120),
				[]*entity.Violation{
					entity.NewViolationInsufficientLimit(),
					entity.NewViolationDoubleTransaction(),
				},
			),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if !reflect.DeepEqual(test.accountExpected, test.operationEventGot.GetAccount()) {
				t.Errorf("%s operationEvent.GetAccount() expected %#v, got %#v", test.name, test.accountExpected, test.operationEventGot.GetAccount())
			}

			if !reflect.DeepEqual(test.violationsExpected, test.operationEventGot.GetViolations()) {
				t.Errorf("%s operationEvent.GetViolations() expected %#v, got %#v", test.name, test.violationsExpected, test.operationEventGot.GetViolations())
			}

			if test.hasViolationsExpected != test.operationEventGot.HasViolations() {
				t.Errorf("%s operationEvent.HasViolations() expected %v, got %v", test.name, test.hasViolationsExpected, test.operationEventGot.HasViolations())
			}
		})
	}
}
