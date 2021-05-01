package validator_test

import (
	"testing"

	"dev-test/nubank-dev-test-2k21/app/entity"
	"dev-test/nubank-dev-test-2k21/app/validator"
)

func TestKindOfValidator(t *testing.T) {
	validator := validator.NewCardActiveValidator()

	if !validator.IsAccountValidator() {
		t.Errorf("validator.IsAccountValidator() expected true, got false")
	}

	if validator.IsTransactionValidator() {
		t.Errorf("validator.IsTransactionValidator() expected false, got true")
	}

	if validator.IsOperationValidator() {
		t.Errorf("validator.IsOperationValidator() expected false, got true")
	}
}
func TestCardActiveValidator(t *testing.T) {
	account := entity.NewAccount(true, 100)
	validator := validator.NewCardActiveValidator()
	violationGot := validator.GetViolation(account)

	if violationGot != nil {
		t.Errorf("validator.GetViolation() expected violation nil, got %s", violationGot.GetName())
	}
}

func TestCardNotActiveValidator(t *testing.T) {
	violationExpected := entity.NewViolationCardNotActive()
	account := entity.NewAccount(false, 100)
	validator := validator.NewCardActiveValidator()
	violationGot := validator.GetViolation(account)

	if violationExpected.GetName() != violationGot.GetName() {
		t.Errorf("validator.GetViolation() expected violation with name %s, got %s", violationExpected.GetName(), violationGot.GetName())
	}
}
