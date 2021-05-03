package validator_test

import (
	"testing"

	"dev-test/nubank-dev-test-2k21/app/entity"
	"dev-test/nubank-dev-test-2k21/app/validator"
)

func TestAccountNotInitializedValidatorIsBreakNextCheck(t *testing.T) {
	validator := validator.NewAccountNotInitializedValidator()

	if !validator.IsBreakNextCheck() {
		t.Errorf("validator.IsBreakNextCheck() expected true, got false")
	}
}

func TestAccountInitializedValidator(t *testing.T) {
	account := entity.NewAccount(true, 100)
	validator := validator.NewAccountNotInitializedValidator()
	violationGot := validator.GetViolation(account, entity.Transaction{})

	if violationGot != nil {
		t.Errorf("validator.GetViolation() expected violation nil, got %s", violationGot.GetName())
	}
}

func TestAccountNotInitializedValidator(t *testing.T) {
	violationExpected := entity.NewViolationAccountNotInitialized()

	account := entity.NewAccountEmpty()
	validator := validator.NewAccountNotInitializedValidator()
	violationGot := validator.GetViolation(account, entity.Transaction{})

	if violationExpected.GetName() != violationGot.GetName() {
		t.Errorf("validator.GetViolation() expected violation with name %s, got %s", violationExpected.GetName(), violationGot.GetName())
	}
}
