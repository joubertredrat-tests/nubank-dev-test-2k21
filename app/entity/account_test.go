package entity_test

import (
	"testing"

	"dev-test/nubank-dev-test-2k21/app/entity"
	"dev-test/nubank-dev-test-2k21/app/helper"
)

func TestAccount(t *testing.T) {
	tests := []struct {
		name                   string
		initializedExpected    bool
		activeCardExpected     bool
		availableLimitExpected entity.Amount
		accountGot             entity.Account
	}{
		{
			name:                   "Test with empty account",
			initializedExpected:    false,
			activeCardExpected:     false,
			availableLimitExpected: entity.NewAmount(0),
			accountGot:             entity.NewAccountEmpty(),
		},
		{
			name:                   "Test with account with inactive card",
			initializedExpected:    true,
			activeCardExpected:     false,
			availableLimitExpected: entity.NewAmount(120),
			accountGot:             entity.NewAccount(false, 120),
		},
		{
			name:                   "Test with account with active card",
			initializedExpected:    true,
			activeCardExpected:     true,
			availableLimitExpected: entity.NewAmount(120),
			accountGot:             entity.NewAccount(true, 120),
		},
		{
			name:                   "Test with account after subtract limit",
			initializedExpected:    true,
			activeCardExpected:     true,
			availableLimitExpected: entity.NewAmount(100),
			accountGot: entity.NewAccountSubtractLimit(
				entity.NewAccount(true, 120),
				entity.NewTransaction("Burger King", 20, helper.GetTimeFromString("2021-04-20T19:25:00.000Z")),
			),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if test.initializedExpected != test.accountGot.IsInitialized() {
				t.Errorf("%s account.IsInitialized() expected %v, got %v", test.name, test.initializedExpected, test.accountGot.IsInitialized())
			}

			if test.activeCardExpected != test.accountGot.IsActiveCard() {
				t.Errorf("%s account.IsActiveCard() expected %v, got %v", test.name, test.activeCardExpected, test.accountGot.IsActiveCard())
			}

			if test.availableLimitExpected != test.accountGot.GetAvailableLimit() {
				t.Errorf("%s account.GetAvailableLimit() expected %#v, got %#v", test.name, test.availableLimitExpected, test.accountGot.GetAvailableLimit())
			}
		})
	}
}
