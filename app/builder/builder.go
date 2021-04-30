package builder

import (
	"dev-test/nubank-dev-test-2k21/app/dto/command"
	"dev-test/nubank-dev-test-2k21/app/entity"
)

func CreateAccountFromCommand(accountLine command.AccountLine) entity.Account {
	return entity.NewAccount(
		accountLine.Account.ActiveCard,
		accountLine.Account.AvailableLimit,
	)
}

func CreateTransactionFromCommand(transactionLine command.TransactionLine) entity.Transaction {
	return entity.NewTransaction(
		transactionLine.Transaction.Merchant,
		transactionLine.Transaction.Amount,
		transactionLine.Transaction.Time,
	)
}
