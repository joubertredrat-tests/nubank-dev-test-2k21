package service

import (
	"errors"
	"fmt"

	"dev-test/nubank-dev-test-2k21/app/builder"
	"dev-test/nubank-dev-test-2k21/app/dto/input"
	"dev-test/nubank-dev-test-2k21/app/entity"
	"dev-test/nubank-dev-test-2k21/app/validator"
)

var (
	ErrAccountRequired = errors.New("account not informed in first operation")
)

type AuthorizeService struct {
	Validators []validator.ValidatorInterface
}

func NewAuthorizeService(validators []validator.ValidatorInterface) AuthorizeService {
	return AuthorizeService{
		Validators: validators,
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

			fmt.Printf("%+v\n", "ttt all")
		}
	}

	fmt.Printf("%+v\n", operations)

	return nil
}
