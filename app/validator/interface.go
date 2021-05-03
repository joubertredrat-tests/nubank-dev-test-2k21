package validator

import (
	"dev-test/nubank-dev-test-2k21/app/entity"
)

type ValidatorInterface interface {
	IsBreakNextCheck() bool
	GetViolation(account entity.Account, transaction entity.Transaction) *entity.Violation
}
