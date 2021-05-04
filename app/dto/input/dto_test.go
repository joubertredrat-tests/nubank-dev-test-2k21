package input_test

import (
	"testing"
	"time"

	"dev-test/nubank-dev-test-2k21/app/dto/input"
)

func TestAccountLine(t *testing.T) {
	accountLine := input.AccountLine{
		Account: struct {
			ActiveCard     bool `json:"active-card"`
			AvailableLimit uint `json:"available-limit"`
		}{
			ActiveCard:     true,
			AvailableLimit: 120,
		},
	}

	if !accountLine.IsAccount() {
		t.Errorf("accountLine.IsAccount() expected true, got false")
	}

	if accountLine.IsTransaction() {
		t.Errorf("accountLine.IsTransaction() expected false, got true")
	}
}

func TestTransactionLine(t *testing.T) {
	transactionLine := input.TransactionLine{
		Transaction: struct {
			Merchant string    `json:"merchant"`
			Amount   uint      `json:"amount"`
			Time     time.Time `json:"time"`
		}{
			Merchant: "Burger King",
			Amount:   20,
			Time:     time.Now(),
		},
	}

	if transactionLine.IsAccount() {
		t.Errorf("transactionLine.IsAccount() expected false, got true")
	}

	if !transactionLine.IsTransaction() {
		t.Errorf("transactionLine.IsTransaction() expected true, got false")
	}
}

func TestOperations(t *testing.T) {
	operations := input.NewOperations()

	if len(operations.Lines) > 0 {
		t.Errorf("operations.Lines expected to have 0 lines, got %d", len(operations.Lines))
	}

	accountLine := input.AccountLine{
		Account: struct {
			ActiveCard     bool `json:"active-card"`
			AvailableLimit uint `json:"available-limit"`
		}{
			ActiveCard:     true,
			AvailableLimit: 120,
		},
	}

	operations.AddLine(accountLine)

	if len(operations.Lines) != 1 {
		t.Errorf("operations.Lines expected to have 1 line, got %d", len(operations.Lines))
	}

	transactionLine := input.TransactionLine{
		Transaction: struct {
			Merchant string    `json:"merchant"`
			Amount   uint      `json:"amount"`
			Time     time.Time `json:"time"`
		}{
			Merchant: "Burger King",
			Amount:   20,
			Time:     time.Now(),
		},
	}

	operations.AddLine(transactionLine)

	if len(operations.Lines) != 2 {
		t.Errorf("operations.Lines expected to have 2 lines, got %d", len(operations.Lines))
	}
}
