package validator_test

import (
	"testing"
	"time"

	"dev-test/nubank-dev-test-2k21/app/dto/input"
	"dev-test/nubank-dev-test-2k21/app/entity"
	"dev-test/nubank-dev-test-2k21/app/helper"
	"dev-test/nubank-dev-test-2k21/app/validator"
)

func TestCardActiveValidator(t *testing.T) {
	account := entity.NewAccount(true, 100)
	transactionLine := input.TransactionLine{
		Transaction: struct {
			Merchant string    `json:"merchant"`
			Amount   uint      `json:"amount"`
			Time     time.Time `json:"time"`
		}{
			Merchant: "Burger King",
			Amount:   20,
			Time:     helper.GetTimeFromString("2021-04-20T19:25:00.000Z"),
		},
	}

	validator := validator.NewCardActiveValidator()
	violationGot := validator.GetViolation(account, transactionLine)

	if violationGot != nil {
		t.Errorf("validator.GetViolation() expected violation nil, got %s", violationGot.GetName())
	}
}

func TestCardNotActiveValidator(t *testing.T) {
	violationExpected := entity.NewViolationCardNotActive()

	account := entity.NewAccount(false, 100)
	transactionLine := input.TransactionLine{
		Transaction: struct {
			Merchant string    `json:"merchant"`
			Amount   uint      `json:"amount"`
			Time     time.Time `json:"time"`
		}{
			Merchant: "Burger King",
			Amount:   20,
			Time:     helper.GetTimeFromString("2021-04-20T19:25:00.000Z"),
		},
	}

	validator := validator.NewCardActiveValidator()
	violationGot := validator.GetViolation(account, transactionLine)

	if violationExpected.GetName() != violationGot.GetName() {
		t.Errorf("validator.GetViolation() expected violation with name %s, got %s", violationExpected.GetName(), violationGot.GetName())
	}
}
