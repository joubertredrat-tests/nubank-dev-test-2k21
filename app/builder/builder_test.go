package builder_test

import (
	"dev-test/nubank-dev-test-2k21/app/builder"
	"dev-test/nubank-dev-test-2k21/app/dto/input"
	"dev-test/nubank-dev-test-2k21/app/dto/output"
	"dev-test/nubank-dev-test-2k21/app/entity"
	"reflect"
	"testing"
	"time"
)

func TestCreateAccountFromInputDTO(t *testing.T) {
	accountLine := input.AccountLine{
		Account: struct {
			ActiveCard     bool `json:"active-card"`
			AvailableLimit uint `json:"available-limit"`
		}{
			ActiveCard:     true,
			AvailableLimit: 120,
		},
	}

	accountExpected := entity.NewAccount(true, 120)
	accountGot := builder.CreateAccountFromInputDTO(accountLine)

	if !reflect.DeepEqual(accountExpected, accountGot) {
		t.Errorf("expected account %+v, got %+v", accountExpected, accountGot)
	}
}

func TestCreateTransactionFromInputDTO(t *testing.T) {
	timeNow := time.Now()
	transactionLine := input.TransactionLine{
		Transaction: struct {
			Merchant string    `json:"merchant"`
			Amount   uint      `json:"amount"`
			Time     time.Time `json:"time"`
		}{
			Merchant: "Burger King",
			Amount:   20,
			Time:     timeNow,
		},
	}

	transactionExpected := entity.NewTransaction("Burger King", 20, timeNow)
	transactionGot := builder.CreateTransactionFromInputDTO(transactionLine)

	if !reflect.DeepEqual(transactionExpected, transactionGot) {
		t.Errorf("expected transaction %+v, got %+v", transactionExpected, transactionGot)
	}
}

func TestCreateAccountOutputDTOFromOperationEvent(t *testing.T) {
	account := entity.NewAccount(true, 120)
	violation := entity.NewViolationAccountNotInitialized()

	accountOutputExpected := output.NewAccountLine(
		account.IsActiveCard(),
		account.GetAvailableLimit().GetValue(),
		[]string{violation.GetName()},
	)

	operationEvent := entity.NewOperationEvent(account, []*entity.Violation{violation})
	accountOutputGot := builder.CreateAccountOutputDTOFromOperationEvent(operationEvent)

	if !reflect.DeepEqual(accountOutputExpected, accountOutputGot) {
		t.Errorf("expected accountOutput %+v, got %+v", accountOutputExpected, accountOutputGot)
	}
}
