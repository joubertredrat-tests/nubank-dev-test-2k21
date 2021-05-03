package validator

import (
	"dev-test/nubank-dev-test-2k21/app/entity"
)

type CardActiveValidator struct{}

func NewCardActiveValidator() *CardActiveValidator {
	return &CardActiveValidator{}
}

func (v *CardActiveValidator) IsBreakNextCheck() bool {
	return true
}

func (v *CardActiveValidator) GetViolation(account entity.Account, transaction entity.Transaction) *entity.Violation {
	if !account.IsActiveCard() {
		violation := entity.NewViolationCardNotActive()
		return &violation
	}

	return nil
}
