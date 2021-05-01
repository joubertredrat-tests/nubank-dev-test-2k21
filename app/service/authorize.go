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
	accountLine := inputOperations.Lines[0]
	if !accountLine.IsAccount() {
		return ErrAccountRequired
	}

	account := builder.CreateAccountFromCommand(accountLine.(input.AccountLine))
	operations := entity.NewOperations()
	operations.RegisterEvent(account)

	for i := 1; i < len(inputOperations.Lines); i++ {
		operationLine := inputOperations.Lines[i]
		if operationLine.IsAccount() {

			operations.RegisterViolationEvent(
				account,
				[]entity.Violation{
					entity.NewViolationAccountAlreadyInitialized(),
				},
			)
			continue
		}

		if operationLine.IsTransaction() {
			fmt.Printf("%+v\n", "ttt")
		}
	}

	fmt.Printf("%+v\n", operations)

	return nil
}
