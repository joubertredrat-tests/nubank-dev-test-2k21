package service

import (
	"errors"
	"fmt"

	"dev-test/nubank-dev-test-2k21/app/builder"
	"dev-test/nubank-dev-test-2k21/app/dto/command"
	"dev-test/nubank-dev-test-2k21/app/entity"
)

var (
	ErrAccountRequired = errors.New("account not informed in first operation")
)

type AuthorizeService struct {
	Account entity.Account
}

func NewAuthorizeService() AuthorizeService {
	return AuthorizeService{}
}

func (a AuthorizeService) HandleOperations(commandOperationsDTO command.Operations) error {
	accountLine := commandOperationsDTO.Lines[0]
	if !accountLine.IsAccount() {
		return ErrAccountRequired
	}

	account := builder.CreateAccountFromCommand(accountLine.(command.AccountLine))
	operations := entity.NewOperations()
	operations.RegisterEvent(account)

	for i := 1; i < len(commandOperationsDTO.Lines); i++ {
		operationLine := commandOperationsDTO.Lines[i]
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
			fmt.Printf("%+v\n", "aaa")
		}
	}

	fmt.Printf("%+v\n", operations)

	return nil
}
