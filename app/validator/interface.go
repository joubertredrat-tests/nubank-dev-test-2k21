package validator

import (
	"dev-test/nubank-dev-test-2k21/app/entity"
)

type ValidatorInterface interface {
	IsAccountValidator() bool
	IsTransactionValidator() bool
	IsOperationValidator() bool
}

type AccountValidatorInterface interface {
	GetViolation(account entity.Account) *entity.Violation
}

type TransactionValidatorInterface interface {
	GetViolation(transaction entity.Transaction) *entity.Violation
}

type OperationValidatorInterface interface {
	GetViolation(account entity.Account, transaction entity.Transaction) *entity.Violation
}
