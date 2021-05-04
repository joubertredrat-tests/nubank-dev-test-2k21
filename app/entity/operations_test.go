package entity_test

import (
	"testing"

	"dev-test/nubank-dev-test-2k21/app/entity"
)

func TestOperations(t *testing.T) {
	operations := entity.NewOperations()

	if len(operations.GetEvents()) > 0 {
		t.Errorf("operations.GetEvents() expected empty events, got %d", len(operations.GetEvents()))
	}

	operations.RegisterViolationEvent(
		entity.NewAccount(true, 120),
		entity.NewViolationAccountNotInitialized(),
	)

	if len(operations.GetEvents()) != 1 {
		t.Errorf("operations.GetEvents() expected 1 event, got %d", len(operations.GetEvents()))
	}

	operations.RegisterEvent(
		entity.NewAccount(true, 120),
		[]*entity.Violation{
			entity.NewViolationHighFrequencySmallInterval(),
			entity.NewViolationDoubleTransaction(),
		},
	)

	if len(operations.GetEvents()) != 2 {
		t.Errorf("operations.GetEvents() expected 2 events, got %d", len(operations.GetEvents()))
	}
}
