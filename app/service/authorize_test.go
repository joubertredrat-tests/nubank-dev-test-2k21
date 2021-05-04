package service_test

import (
	"reflect"
	"testing"
	"time"

	"dev-test/nubank-dev-test-2k21/app/dto/input"
	"dev-test/nubank-dev-test-2k21/app/entity"
	"dev-test/nubank-dev-test-2k21/app/helper"
	"dev-test/nubank-dev-test-2k21/app/service"
	"dev-test/nubank-dev-test-2k21/app/validator"
)

func TestAuthorizeService(t *testing.T) {
	tests := []struct {
		name                  string
		getInputOperations    func() input.Operations
		getOperationsExpected func() entity.Operations
	}{
		{
			name: "Test handle operations with no violations",
			getInputOperations: func() input.Operations {
				operations := input.NewOperations()
				operations.AddLine(input.AccountLine{
					Account: struct {
						ActiveCard     bool `json:"active-card"`
						AvailableLimit uint `json:"available-limit"`
					}{
						ActiveCard:     true,
						AvailableLimit: 120,
					},
				})
				operations.AddLine(input.TransactionLine{
					Transaction: struct {
						Merchant string    `json:"merchant"`
						Amount   uint      `json:"amount"`
						Time     time.Time `json:"time"`
					}{
						Merchant: "Burger King",
						Amount:   20,
						Time:     helper.GetTimeFromString("2021-04-20T19:25:00.000Z"),
					},
				})
				operations.AddLine(input.TransactionLine{
					Transaction: struct {
						Merchant string    `json:"merchant"`
						Amount   uint      `json:"amount"`
						Time     time.Time `json:"time"`
					}{
						Merchant: "Habib's",
						Amount:   10,
						Time:     helper.GetTimeFromString("2021-04-20T19:42:00.000Z"),
					},
				})
				operations.AddLine(input.TransactionLine{
					Transaction: struct {
						Merchant string    `json:"merchant"`
						Amount   uint      `json:"amount"`
						Time     time.Time `json:"time"`
					}{
						Merchant: "Bob's",
						Amount:   15,
						Time:     helper.GetTimeFromString("2021-04-21T07:04:00.000Z"),
					},
				})
				operations.AddLine(input.TransactionLine{
					Transaction: struct {
						Merchant string    `json:"merchant"`
						Amount   uint      `json:"amount"`
						Time     time.Time `json:"time"`
					}{
						Merchant: "Subway",
						Amount:   30,
						Time:     helper.GetTimeFromString("2021-04-21T07:15:00.000Z"),
					},
				})

				return operations
			},
			getOperationsExpected: func() entity.Operations {
				operations := entity.NewOperations()
				operations.RegisterEvent(
					entity.NewAccount(true, 120),
					entity.NewViolationsEmpty(),
				)
				operations.RegisterEvent(
					entity.NewAccount(true, 100),
					entity.NewViolationsEmpty(),
				)
				operations.RegisterEvent(
					entity.NewAccount(true, 90),
					entity.NewViolationsEmpty(),
				)
				operations.RegisterEvent(
					entity.NewAccount(true, 75),
					entity.NewViolationsEmpty(),
				)
				operations.RegisterEvent(
					entity.NewAccount(true, 45),
					entity.NewViolationsEmpty(),
				)

				return operations
			},
		},
	}

	validatorManager := validator.NewManager(
		[]validator.ValidatorInterface{
			validator.NewCardLimitValidator(),
			validator.NewHighTransactionsValidator(3, 120),
			validator.NewDoubleTransactionValidator(120),
		},
	)

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			authorizeService := service.NewAuthorizeService(validatorManager)
			operationsGot := authorizeService.HandleOperations(test.getInputOperations())

			if !reflect.DeepEqual(test.getOperationsExpected(), operationsGot) {
				t.Errorf("%s authorizeService.HandleOperations() expected %#v, got %#v", test.name, test.getOperationsExpected(), operationsGot)
			}
		})
	}
}
