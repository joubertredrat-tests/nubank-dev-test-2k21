package validator

import (
	"dev-test/nubank-dev-test-2k21/app/entity"
)

type CardLimitValidator struct{}

func NewCardLimitValidator() *CardLimitValidator {
	return &CardLimitValidator{}
}

func (v *CardLimitValidator) IsBreakNextCheck() bool {
	return false
}

func (v *CardLimitValidator) GetViolation(account entity.Account, transaction entity.Transaction) *entity.Violation {
	if transaction.GetAmount().GetValue() > account.GetAvailableLimit().GetValue() {
		return entity.NewViolationInsufficientLimit()
	}

	return nil
}
