package validator_test

import (
	"testing"
	"time"

	"dev-test/nubank-dev-test-2k21/app/dto/input"
	"dev-test/nubank-dev-test-2k21/app/entity"
	"dev-test/nubank-dev-test-2k21/app/helper"
	"dev-test/nubank-dev-test-2k21/app/validator"
)

func TestNoViolation(t *testing.T) {
	account := entity.NewAccount(true, 100)
	transactionLine := input.TransactionLine{
		Transaction: struct {
			Merchant string    `json:"merchant"`
			Amount   uint      `json:"amount"`
			Time     time.Time `json:"time"`
		}{
			Merchant: "Burger King",
			Amount:   80,
			Time:     time.Now(),
		},
	}

	validator := validator.NewHighTransactionsValidator(3, 120)
	violationGot := validator.GetViolation(account, transactionLine)

	if violationGot != nil {
		t.Errorf("validator.GetViolation(account, transactionLine) expected nil, got *entity.Violation")
	}
}

func TestHighTransactionsValidator(t *testing.T) {
	account := entity.NewAccount(true, 100)

	tests := []struct {
		name               string
		transactionsLine   []input.TransactionLine
		violationsExpected int
	}{
		{
			name:               "Test with 2 transactions line",
			transactionsLine:   getTwoTransactionsLine(),
			violationsExpected: 0,
		},
		{
			name:               "Test with 14 transactions line",
			transactionsLine:   getFourteenTransactionsLine(),
			violationsExpected: 0,
		},
		{
			name:               "Test with 7 transactions with 1 violation",
			transactionsLine:   getSevenTransactionsLineWithOneViolation(),
			violationsExpected: 1,
		},
		{
			name:               "Test with 14 transactions with 3 violations",
			transactionsLine:   getFourteenTransactionsLineWithThreeViolations(),
			violationsExpected: 3,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			violationsGot := 0
			validator := validator.NewHighTransactionsValidator(3, 120)
			for _, transactionLine := range test.transactionsLine {
				violationGot := validator.GetViolation(account, transactionLine)
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

func getTwoTransactionsLine() []input.TransactionLine {
	transactionsLine := []input.TransactionLine{}

	transactionsLine = append(
		transactionsLine,
		getTransactionLine("Burger King", 20, "2021-04-20T19:25:00.000Z"),
		getTransactionLine("Habib's", 20, "2021-04-20T19:42:00.000Z"),
	)

	return transactionsLine
}

func getFourteenTransactionsLine() []input.TransactionLine {
	transactionsLine := []input.TransactionLine{}

	transactionsLine = append(
		transactionsLine,
		getTransactionLine("Burger King", 20, "2021-04-20T19:25:00.000Z"),
		getTransactionLine("Habib's", 20, "2021-04-20T19:42:00.000Z"),
		getTransactionLine("Bob's", 20, "2021-04-21T07:04:00.000Z"),
		getTransactionLine("Subway", 20, "2021-04-21T07:15:00.000Z"),
		getTransactionLine("Domino's Pizza", 20, "2021-04-21T09:26:00.000Z"),
		getTransactionLine("Pizza Hut", 20, "2021-04-21T12:30:00.000Z"),
		getTransactionLine("Burger King", 20, "2021-04-21T15:17:00.000Z"),
		getTransactionLine("Bob's", 20, "2021-04-21T17:33:00.000Z"),
		getTransactionLine("Yogoberry", 20, "2021-04-21T20:21:00.000Z"),
		getTransactionLine("Subway", 20, "2021-04-21T23:01:00.000Z"),
		getTransactionLine("Starbucks", 20, "2021-04-22T06:10:00.000Z"),
		getTransactionLine("Burger King", 20, "2021-04-22T09:22:00.000Z"),
		getTransactionLine("Burger King", 20, "2021-04-22T09:57:00.000Z"),
		getTransactionLine("Mcdonald's", 20, "2021-04-22T14:05:00.000Z"),
	)

	return transactionsLine
}

func getSevenTransactionsLineWithOneViolation() []input.TransactionLine {
	transactionsLine := []input.TransactionLine{}

	transactionsLine = append(
		transactionsLine,
		getTransactionLine("Burger King", 20, "2021-04-20T19:25:00.000Z"),
		getTransactionLine("Habib's", 20, "2021-04-20T19:42:00.000Z"),
		getTransactionLine("Bob's", 20, "2021-04-21T07:04:00.000Z"),
		getTransactionLine("Subway", 20, "2021-04-21T07:05:00.000Z"),
		getTransactionLine("Domino's Pizza", 20, "2021-04-21T07:05:00.000Z"),
		getTransactionLine("Pizza Hut", 20, "2021-04-21T12:30:00.000Z"),
		getTransactionLine("Burger King", 20, "2021-04-21T15:17:00.000Z"),
	)

	return transactionsLine
}

func getFourteenTransactionsLineWithThreeViolations() []input.TransactionLine {
	transactionsLine := []input.TransactionLine{}

	transactionsLine = append(
		transactionsLine,
		getTransactionLine("Burger King", 20, "2021-04-20T19:25:00.000Z"),
		getTransactionLine("Habib's", 20, "2021-04-20T19:42:00.000Z"),
		getTransactionLine("Bob's", 20, "2021-04-20T19:42:00.000Z"),
		getTransactionLine("Subway", 20, "2021-04-20T19:43:00.000Z"),
		getTransactionLine("Domino's Pizza", 20, "2021-04-21T07:06:00.000Z"),
		getTransactionLine("Pizza Hut", 20, "2021-04-21T12:30:00.000Z"),
		getTransactionLine("Burger King", 20, "2021-04-21T12:31:00.000Z"),
		getTransactionLine("Bob's", 20, "2021-04-21T12:31:00.000Z"),
		getTransactionLine("Yogoberry", 20, "2021-04-21T20:21:00.000Z"),
		getTransactionLine("Subway", 20, "2021-04-21T23:01:00.000Z"),
		getTransactionLine("Starbucks", 20, "2021-04-22T06:10:00.000Z"),
		getTransactionLine("Burger King", 20, "2021-04-22T09:22:00.000Z"),
		getTransactionLine("Burger King", 20, "2021-04-22T09:22:00.000Z"),
		getTransactionLine("Mcdonald's", 20, "2021-04-22T09:22:00.000Z"),
	)

	return transactionsLine
}

func getTransactionLine(merchant string, amount uint, timeString string) input.TransactionLine {
	return input.TransactionLine{
		Transaction: struct {
			Merchant string    `json:"merchant"`
			Amount   uint      `json:"amount"`
			Time     time.Time `json:"time"`
		}{
			Merchant: merchant,
			Amount:   amount,
			Time:     helper.GetTimeFromString(timeString),
		},
	}
}
