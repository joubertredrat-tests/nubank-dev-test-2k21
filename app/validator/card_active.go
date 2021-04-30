package validator

import (
	"dev-test/nubank-dev-test-2k21/app/dto/input"
	"dev-test/nubank-dev-test-2k21/app/entity"
)

type CardActiveValidator struct{}

func NewCardActiveValidator() ValidatorInterface {
	return &CardActiveValidator{}
}

func (v *CardActiveValidator) GetViolation(account entity.Account, transactionLine input.TransactionLine) *entity.Violation {
	if !account.ActiveCard {
		violation := entity.NewViolationCardNotActive()
		return &violation
	}

	return nil
}
