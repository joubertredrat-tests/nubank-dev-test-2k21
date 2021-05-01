package validator

import (
	"dev-test/nubank-dev-test-2k21/app/dto/input"
	"dev-test/nubank-dev-test-2k21/app/entity"
)

type CardLimitValidator struct{}

func NewCardLimitValidator() ValidatorInterface {
	return CardLimitValidator{}
}

func (v CardLimitValidator) GetViolation(account entity.Account, transactionLine input.TransactionLine) *entity.Violation {
	if transactionLine.Transaction.Amount > account.AvailableLimit.Value {
		violation := entity.NewViolationInsufficientLimit()
		return &violation
	}

	return nil
}
