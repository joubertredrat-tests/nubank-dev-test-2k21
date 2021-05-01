package validator_test

import (
	"testing"

	"dev-test/nubank-dev-test-2k21/app/entity"
	"dev-test/nubank-dev-test-2k21/app/helper"
	"dev-test/nubank-dev-test-2k21/app/validator"
)

func TestKindOfHighTransactionsValidator(t *testing.T) {
	validator := validator.NewHighTransactionsValidator(3, 120)

	if validator.IsAccountValidator() {
		t.Errorf("validator.IsAccountValidator() expected false, got true")
	}

	if !validator.IsTransactionValidator() {
		t.Errorf("validator.IsTransactionValidator() expected true, got false")
	}

	if validator.IsOperationValidator() {
		t.Errorf("validator.IsOperationValidator() expected false, got true")
	}
}

func TestHighTransactionsValidator(t *testing.T) {
	tests := []struct {
		name               string
		violationsExpected int
		getTransactions    func() []entity.Transaction
	}{
		{
			name:               "Test with 2 transactions",
			violationsExpected: 0,
			getTransactions: func() []entity.Transaction {
				transactions := []entity.Transaction{}

				transactions = append(
					transactions,
					entity.NewTransaction("Burger King", 20, helper.GetTimeFromString("2021-04-20T19:25:00.000Z")),
					entity.NewTransaction("Habib's", 20, helper.GetTimeFromString("2021-04-20T19:42:00.000Z")),
				)

				return transactions
			},
		},
		{
			name:               "Test with 14 transactions",
			violationsExpected: 0,
			getTransactions: func() []entity.Transaction {
				transactions := []entity.Transaction{}

				transactions = append(
					transactions,
					entity.NewTransaction("Burger King", 20, helper.GetTimeFromString("2021-04-20T19:25:00.000Z")),
					entity.NewTransaction("Habib's", 20, helper.GetTimeFromString("2021-04-20T19:42:00.000Z")),
					entity.NewTransaction("Bob's", 20, helper.GetTimeFromString("2021-04-21T07:04:00.000Z")),
					entity.NewTransaction("Subway", 20, helper.GetTimeFromString("2021-04-21T07:15:00.000Z")),
					entity.NewTransaction("Domino's Pizza", 20, helper.GetTimeFromString("2021-04-21T09:26:00.000Z")),
					entity.NewTransaction("Pizza Hut", 20, helper.GetTimeFromString("2021-04-21T12:30:00.000Z")),
					entity.NewTransaction("Burger King", 20, helper.GetTimeFromString("2021-04-21T15:17:00.000Z")),
					entity.NewTransaction("Bob's", 20, helper.GetTimeFromString("2021-04-21T17:33:00.000Z")),
					entity.NewTransaction("Yogoberry", 20, helper.GetTimeFromString("2021-04-21T20:21:00.000Z")),
					entity.NewTransaction("Subway", 20, helper.GetTimeFromString("2021-04-21T23:01:00.000Z")),
					entity.NewTransaction("Starbucks", 20, helper.GetTimeFromString("2021-04-22T06:10:00.000Z")),
					entity.NewTransaction("Burger King", 20, helper.GetTimeFromString("2021-04-22T09:22:00.000Z")),
					entity.NewTransaction("Burger King", 20, helper.GetTimeFromString("2021-04-22T09:57:00.000Z")),
					entity.NewTransaction("Mcdonald's", 20, helper.GetTimeFromString("2021-04-22T14:05:00.000Z")),
				)

				return transactions
			},
		},
		{
			name:               "Test with 7 transactions with 1 violation",
			violationsExpected: 1,
			getTransactions: func() []entity.Transaction {
				transactions := []entity.Transaction{}

				transactions = append(
					transactions,
					entity.NewTransaction("Burger King", 20, helper.GetTimeFromString("2021-04-20T19:25:00.000Z")),
					entity.NewTransaction("Habib's", 20, helper.GetTimeFromString("2021-04-20T19:42:00.000Z")),
					entity.NewTransaction("Bob's", 20, helper.GetTimeFromString("2021-04-21T07:04:00.000Z")),
					entity.NewTransaction("Subway", 20, helper.GetTimeFromString("2021-04-21T07:05:00.000Z")),
					entity.NewTransaction("Domino's Pizza", 20, helper.GetTimeFromString("2021-04-21T07:05:00.000Z")),
					entity.NewTransaction("Pizza Hut", 20, helper.GetTimeFromString("2021-04-21T12:30:00.000Z")),
					entity.NewTransaction("Burger King", 20, helper.GetTimeFromString("2021-04-21T15:17:00.000Z")),
				)

				return transactions
			},
		},
		{
			name:               "Test with 14 transactions with 3 violations",
			violationsExpected: 3,
			getTransactions: func() []entity.Transaction {
				transactions := []entity.Transaction{}

				transactions = append(
					transactions,
					entity.NewTransaction("Burger King", 20, helper.GetTimeFromString("2021-04-20T19:25:00.000Z")),
					entity.NewTransaction("Habib's", 20, helper.GetTimeFromString("2021-04-20T19:42:00.000Z")),
					entity.NewTransaction("Bob's", 20, helper.GetTimeFromString("2021-04-20T19:42:00.000Z")),
					entity.NewTransaction("Subway", 20, helper.GetTimeFromString("2021-04-20T19:43:00.000Z")),
					entity.NewTransaction("Domino's Pizza", 20, helper.GetTimeFromString("2021-04-21T07:06:00.000Z")),
					entity.NewTransaction("Pizza Hut", 20, helper.GetTimeFromString("2021-04-21T12:30:00.000Z")),
					entity.NewTransaction("Burger King", 20, helper.GetTimeFromString("2021-04-21T12:31:00.000Z")),
					entity.NewTransaction("Bob's", 20, helper.GetTimeFromString("2021-04-21T12:31:00.000Z")),
					entity.NewTransaction("Yogoberry", 20, helper.GetTimeFromString("2021-04-21T20:21:00.000Z")),
					entity.NewTransaction("Subway", 20, helper.GetTimeFromString("2021-04-21T23:01:00.000Z")),
					entity.NewTransaction("Starbucks", 20, helper.GetTimeFromString("2021-04-22T06:10:00.000Z")),
					entity.NewTransaction("Burger King", 20, helper.GetTimeFromString("2021-04-22T09:22:00.000Z")),
					entity.NewTransaction("Burger King", 20, helper.GetTimeFromString("2021-04-22T09:22:00.000Z")),
					entity.NewTransaction("Mcdonald's", 20, helper.GetTimeFromString("2021-04-22T09:22:00.000Z")),
				)

				return transactions
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			violationsGot := 0
			validator := validator.NewHighTransactionsValidator(3, 120)
			for _, transaction := range test.getTransactions() {
				violationGot := validator.GetViolation(transaction)
				if violationGot != nil {
					violationsGot++
				}
			}

			if test.violationsExpected != violationsGot {
				t.Errorf("%s expected %d violations, got %d", test.name, test.violationsExpected, violationsGot)
			}
		})
	}
}

func TestViolationHighFrequencySmallInterval(t *testing.T) {
	violationExpected := entity.NewViolationHighFrequencySmallInterval()

	transactionOne := entity.NewTransaction("Burger King", 20, helper.GetTimeFromString("2021-04-20T19:25:00.000Z"))
	transactionTwo := entity.NewTransaction("Habib's", 20, helper.GetTimeFromString("2021-04-20T19:25:00.000Z"))
	transactionThree := entity.NewTransaction("Bob's", 20, helper.GetTimeFromString("2021-04-20T19:25:00.000Z"))

	validator := validator.NewHighTransactionsValidator(3, 120)
	validator.GetViolation(transactionOne)
	validator.GetViolation(transactionTwo)
	violationGot := validator.GetViolation(transactionThree)

	if violationExpected.GetName() != violationGot.GetName() {
		t.Errorf("validator.GetViolation() expected violation with name %s, got %s", violationExpected.GetName(), violationGot.GetName())
	}
}
