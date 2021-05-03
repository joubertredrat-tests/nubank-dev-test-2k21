package validator

import "dev-test/nubank-dev-test-2k21/app/entity"

type Manager struct {
	validators []ValidatorInterface
}

func NewManager(validators []ValidatorInterface) Manager {
	return Manager{
		validators: validators,
	}
}

func (m Manager) GetViolations(account entity.Account, transaction entity.Transaction) []*entity.Violation {
	violations := []*entity.Violation{}

	for _, validator := range m.validators {
		violation := validator.GetViolation(account, transaction)

		if violation != nil && validator.IsBreakNextCheck() {
			violations = append(violations, violation)
			return violations
		}

		if violation != nil {
			violations = append(violations, violation)
		}
	}

	return violations
}
