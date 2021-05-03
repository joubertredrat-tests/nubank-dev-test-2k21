package validator

import "dev-test/nubank-dev-test-2k21/app/entity"

type AccountAlreadyInitializedValidator struct{}

func NewAccountAlreadyInitializedValidator() *AccountAlreadyInitializedValidator {
	return &AccountAlreadyInitializedValidator{}
}

func (v *AccountAlreadyInitializedValidator) IsBreakNextCheck() bool {
	return true
}

func (v *AccountAlreadyInitializedValidator) GetViolation(account entity.Account, transaction entity.Transaction) *entity.Violation {
	if account.IsInitialized() {
		return entity.NewViolationAccountAlreadyInitialized()
	}

	return nil
}
