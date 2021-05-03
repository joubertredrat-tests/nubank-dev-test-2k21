package validator_test

import (
	"testing"

	"dev-test/nubank-dev-test-2k21/app/entity"
	"dev-test/nubank-dev-test-2k21/app/validator"
)

func TestAccountAlreadyInitializedValidatorIsBreakNextCheck(t *testing.T) {
	validator := validator.NewAccountAlreadyInitializedValidator()

	if !validator.IsBreakNextCheck() {
		t.Errorf("validator.IsBreakNextCheck() expected true, got false")
	}
}

func TestAccountNotAlreadyInitializedValidator(t *testing.T) {
	account := entity.NewAccountEmpty()
	validator := validator.NewAccountAlreadyInitializedValidator()
	violationGot := validator.GetViolation(account, entity.Transaction{})

	if violationGot != nil {
		t.Errorf("validator.GetViolation() expected violation nil, got %s", violationGot.GetName())
	}
}

func TestAccountAlreadyInitializedValidator(t *testing.T) {
	violationExpected := entity.NewViolationAccountAlreadyInitialized()

	account := entity.NewAccount(true, 100)
	validator := validator.NewAccountAlreadyInitializedValidator()
	violationGot := validator.GetViolation(account, entity.Transaction{})

	if violationExpected.GetName() != violationGot.GetName() {
		t.Errorf("validator.GetViolation() expected violation with name %s, got %s", violationExpected.GetName(), violationGot.GetName())
	}
}
