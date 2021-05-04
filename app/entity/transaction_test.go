package entity_test

import (
	"testing"
	"time"

	"dev-test/nubank-dev-test-2k21/app/entity"
)

func TestTransaction(t *testing.T) {
	merchantExpected := "Burger King"
	amountExpected := entity.NewAmount(20)
	timeExpected := time.Now()

	transactionGot := entity.NewTransaction(merchantExpected, amountExpected.GetValue(), timeExpected)

	if merchantExpected != transactionGot.GetMerchant() {
		t.Errorf("transaction.GetMerchant() expected %v, got %v", merchantExpected, transactionGot.GetMerchant())
	}

	if amountExpected != transactionGot.GetAmount() {
		t.Errorf("transaction.GetAmount() expected %#v, got %#v", amountExpected, transactionGot.GetAmount())
	}

	if timeExpected != transactionGot.GetTime() {
		t.Errorf("transaction.GetTime() expected %#v, got %#v", timeExpected, transactionGot.GetTime())
	}
}
