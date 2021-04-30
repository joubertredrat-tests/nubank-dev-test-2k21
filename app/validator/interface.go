package validator

import (
	"dev-test/nubank-dev-test-2k21/app/dto/input"
	"dev-test/nubank-dev-test-2k21/app/entity"
)

type ValidatorInterface interface {
	GetViolation(account entity.Account, transactionLine input.TransactionLine) *entity.Violation
}
