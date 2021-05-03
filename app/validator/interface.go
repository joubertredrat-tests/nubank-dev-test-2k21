package validator

import (
	"dev-test/nubank-dev-test-2k21/app/entity"
)

type ValidatorInterface interface {
	GetViolation(account entity.Account, transaction entity.Transaction) *entity.Violation
}
