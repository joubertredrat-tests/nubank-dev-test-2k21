package validator_test

import (
	"testing"

	"dev-test/nubank-dev-test-2k21/app/entity"
	"dev-test/nubank-dev-test-2k21/app/helper"
	"dev-test/nubank-dev-test-2k21/app/validator"
)

func TestKindOfCardLimitValidator(t *testing.T) {
	validator := validator.NewCardLimitValidator()

	if validator.IsAccountValidator() {
		t.Errorf("validator.IsAccountValidator() expected false, got true")
	}

	if validator.IsTransactionValidator() {
		t.Errorf("validator.IsTransactionValidator() expected false, got true")
	}

	if !validator.IsOperationValidator() {
		t.Errorf("validator.IsOperationValidator() expected true, got false")
	}
}

func TestCardLimitValidator(t *testing.T) {
	account := entity.NewAccount(true, 100)
	transaction := entity.NewTransaction("Burger King", 20, helper.GetTimeFromString("2021-04-20T19:25:00.000Z"))
	validator := validator.NewCardLimitValidator()
	violationGot := validator.GetViolation(account, transaction)

	if violationGot != nil {
		t.Errorf("validator.GetViolation() expected violation nil, got %s", violationGot.GetName())
	}
}

func TestCardNoLimitValidator(t *testing.T) {
	violationExpected := entity.NewViolationInsufficientLimit()

	account := entity.NewAccount(true, 100)
	transaction := entity.NewTransaction("Burger King", 120, helper.GetTimeFromString("2021-04-20T19:25:00.000Z"))
	validator := validator.NewCardLimitValidator()
	violationGot := validator.GetViolation(account, transaction)

	if violationExpected.GetName() != violationGot.GetName() {
		t.Errorf("validator.GetViolation() expected violation with name %s, got %s", violationExpected.GetName(), violationGot.GetName())
	}
}
