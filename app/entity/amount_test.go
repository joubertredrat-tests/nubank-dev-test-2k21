package entity_test

import (
	"testing"

	"dev-test/nubank-dev-test-2k21/app/entity"
)

func TestAmount(t *testing.T) {
	valueExpected := uint(120)
	amount := entity.NewAmount(valueExpected)

	if valueExpected != amount.GetValue() {
		t.Errorf("amount.GetValue() expected %d, got %d", valueExpected, amount.GetValue())
	}
}
