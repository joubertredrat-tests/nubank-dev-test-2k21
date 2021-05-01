package validator

import (
	"dev-test/nubank-dev-test-2k21/app/entity"
)

type CardLimitValidator struct{}

func NewCardLimitValidator() *CardLimitValidator {
	return &CardLimitValidator{}
}

func (v *CardLimitValidator) IsAccountValidator() bool {
	return false
}

func (v *CardLimitValidator) IsTransactionValidator() bool {
	return false
}

func (v *CardLimitValidator) IsOperationValidator() bool {
	return true
}

func (v CardLimitValidator) GetViolation(account entity.Account, transaction entity.Transaction) *entity.Violation {
	if transaction.Amount.Value > account.AvailableLimit.Value {
		violation := entity.NewViolationInsufficientLimit()
		return &violation
	}

	return nil
}
