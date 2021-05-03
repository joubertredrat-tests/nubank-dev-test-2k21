package service

import (
	"dev-test/nubank-dev-test-2k21/app/builder"
	"dev-test/nubank-dev-test-2k21/app/dto/input"
	"dev-test/nubank-dev-test-2k21/app/entity"
	"dev-test/nubank-dev-test-2k21/app/validator"
)

type AuthorizeService struct {
	validatorManager validator.Manager
}

func NewAuthorizeService(validatorManager validator.Manager) AuthorizeService {
	return AuthorizeService{
		validatorManager: validatorManager,
	}
}

func (a AuthorizeService) HandleOperations(inputOperations input.Operations) entity.Operations {
	operations := entity.NewOperations()
	account := entity.NewAccountEmpty()

	for _, operationLine := range inputOperations.Lines {
		if operationLine.IsAccount() && account.IsInitialized() {
			operations.RegisterViolationEvent(account, entity.NewViolationAccountAlreadyInitialized())

			continue
		}

		if operationLine.IsTransaction() && !account.IsInitialized() {
			operations.RegisterViolationEvent(account, entity.NewViolationAccountNotInitialized())

			continue
		}

		if operationLine.IsTransaction() && !account.IsActiveCard() {
			operations.RegisterViolationEvent(account, entity.NewViolationCardNotActive())

			continue
		}

		if operationLine.IsAccount() {
			account = builder.CreateAccountFromInputDTO(operationLine.(input.AccountLine))
			operations.RegisterEvent(account, entity.NewViolationsEmpty())
			continue
		}

		if operationLine.IsTransaction() {
			transaction := builder.CreateTransactionFromInputDTO(operationLine)
			violations := a.validatorManager.GetViolations(account, transaction)

			if 1 > len(violations) {
				account = entity.NewAccountSubtractLimit(account, transaction)
			}

			operations.RegisterEvent(account, a.validatorManager.GetViolations(account, transaction))
		}
	}

	return operations
}
