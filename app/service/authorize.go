package service

import (
	"fmt"

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

func (a AuthorizeService) HandleOperations(inputOperations input.Operations) error {
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

		if operationLine.IsAccount() {
			account = builder.CreateAccountFromInputDTO(operationLine.(input.AccountLine))
			operations.RegisterEvent(account)
			continue
		}

		if operationLine.IsTransaction() {
			// violations := []entity.Violation
			// transaction := builder.CreateTransactionFromInputDTO(operationLine)

			// for _, validator := range a.Validators {
			// 	if validator.IsAccountValidator() {
			// 		validator.GetViolation(account)
			// 	}
			// }

			fmt.Printf("%+v\n", "ttt all")
		}
	}

	fmt.Printf("%+v\n", operations)

	return nil
}
