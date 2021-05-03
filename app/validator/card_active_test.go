package validator_test

import (
	"testing"

	"dev-test/nubank-dev-test-2k21/app/entity"
	"dev-test/nubank-dev-test-2k21/app/validator"
)

func TestCardActiveValidatorIsBreakNextCheck(t *testing.T) {
	validator := validator.NewCardActiveValidator()

	if !validator.IsBreakNextCheck() {
		t.Errorf("validator.IsBreakNextCheck() expected true, got false")
	}
}
func TestCardActiveValidator(t *testing.T) {
	validator := validator.NewCardActiveValidator()
	violationGot := validator.GetViolation(entity.NewAccount(true, 100), entity.Transaction{})

	if violationGot != nil {
		t.Errorf("validator.GetViolation() expected violation nil, got %s", violationGot.GetName())
	}
}

func TestCardNotActiveValidator(t *testing.T) {
	violationExpected := entity.NewViolationCardNotActive()
	validator := validator.NewCardActiveValidator()
	violationGot := validator.GetViolation(entity.NewAccount(false, 100), entity.Transaction{})

	if violationExpected.GetName() != violationGot.GetName() {
		t.Errorf("validator.GetViolation() expected violation with name %s, got %s", violationExpected.GetName(), violationGot.GetName())
	}
}
