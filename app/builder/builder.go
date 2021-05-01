package builder

import (
	"dev-test/nubank-dev-test-2k21/app/dto/input"
	"dev-test/nubank-dev-test-2k21/app/entity"
)

func CreateAccountFromInputDTO(operationLine input.OperationLine) entity.Account {
	return entity.NewAccount(
		operationLine.(input.AccountLine).Account.ActiveCard,
		operationLine.(input.AccountLine).Account.AvailableLimit,
	)
}

func CreateTransactionFromInputDTO(operationLine input.OperationLine) entity.Transaction {
	return entity.NewTransaction(
		operationLine.(input.TransactionLine).Transaction.Merchant,
		operationLine.(input.TransactionLine).Transaction.Amount,
		operationLine.(input.TransactionLine).Transaction.Time,
	)
}
