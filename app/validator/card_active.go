package validator

import (
	"dev-test/nubank-dev-test-2k21/app/entity"
)

type CardActiveValidator struct{}

func NewCardActiveValidator() *CardActiveValidator {
	return &CardActiveValidator{}
}

func (v *CardActiveValidator) IsAccountValidator() bool {
	return true
}

func (v *CardActiveValidator) IsTransactionValidator() bool {
	return false
}

func (v *CardActiveValidator) IsOperationValidator() bool {
	return false
}

func (v *CardActiveValidator) GetViolation(account entity.Account) *entity.Violation {
	if !account.ActiveCard {
		violation := entity.NewViolationCardNotActive()
		return &violation
	}

	return nil
}
