package validator

import "dev-test/nubank-dev-test-2k21/app/entity"

type AccountNotInitializedValidator struct{}

func NewAccountNotInitializedValidator() *AccountNotInitializedValidator {
	return &AccountNotInitializedValidator{}
}

func (v *AccountNotInitializedValidator) IsBreakNextCheck() bool {
	return true
}

func (v *AccountNotInitializedValidator) GetViolation(account entity.Account, transaction entity.Transaction) *entity.Violation {
	if !account.IsInitialized() {
		violation := entity.NewViolationAccountNotInitialized()
		return &violation
	}

	return nil
}
