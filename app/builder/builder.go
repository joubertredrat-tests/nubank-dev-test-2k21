package builder

import (
	"dev-test/nubank-dev-test-2k21/app/dto/input"
	"dev-test/nubank-dev-test-2k21/app/entity"
)

func CreateAccountFromCommand(accountLine input.AccountLine) entity.Account {
	return entity.NewAccount(
		accountLine.Account.ActiveCard,
		accountLine.Account.AvailableLimit,
	)
}

func CreateTransactionFromCommand(transactionLine input.TransactionLine) entity.Transaction {
	return entity.NewTransaction(
		transactionLine.Transaction.Merchant,
		transactionLine.Transaction.Amount,
		transactionLine.Transaction.Time,
	)
}
