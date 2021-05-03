package validator_test

import (
	"fmt"
	"reflect"
	"testing"

	"dev-test/nubank-dev-test-2k21/app/entity"
	"dev-test/nubank-dev-test-2k21/app/helper"
	"dev-test/nubank-dev-test-2k21/app/validator"
)

type ViolationData struct {
	Account            entity.Account
	Transaction        entity.Transaction
	ViolationsExpected []*entity.Violation
}

func TestManager(t *testing.T) {
	tests := []struct {
		name           string
		violationsData []ViolationData
		getManager     func() validator.Manager
	}{
		{
			name:           "Test validation with no violations",
			violationsData: getOperationsDataWithNoViolations(),
		},
		{
			name:           "Test validation with no card limit violations",
			violationsData: getOperationsDataWithNoCardLimitViolations(),
		},
		{
			name:           "Test validation with no high transaction in small interval violations",
			violationsData: getOperationsDataWithHighTransactionsViolations(),
		},
		{
			name:           "Test validation with double transaction violations",
			violationsData: getOperationsDataWithDoubleTransactionViolations(),
		},
		{
			name:           "Test validation with multiple violations",
			violationsData: getOperationsDataWithMultipleViolations(),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			manager := validator.NewManager(
				[]validator.ValidatorInterface{
					validator.NewCardLimitValidator(),
					validator.NewHighTransactionsValidator(3, 120),
					validator.NewDoubleTransactionValidator(120),
				},
			)

			for _, violationData := range test.violationsData {
				violationsGot := manager.GetViolations(violationData.Account, violationData.Transaction)

				if !reflect.DeepEqual(violationData.ViolationsExpected, violationsGot) {
					t.Errorf(
						"%s expected %s violations, got %s",
						test.name,
						fmt.Sprintf("%+v", getViolationsNames(violationData.ViolationsExpected)),
						fmt.Sprintf("%+v", getViolationsNames(violationsGot)),
					)
				}
			}
		})
	}
}

func getViolationsNames(violations []*entity.Violation) []string {
	names := []string{}

	for _, violation := range violations {
		names = append(names, violation.GetName())
	}

	return names
}

func getOperationsDataWithNoViolations() []ViolationData {
	return []ViolationData{
		{
			Account:            entity.NewAccount(true, 200),
			Transaction:        entity.NewTransaction("Burger King", 20, helper.GetTimeFromString("2021-04-20T19:25:00.000Z")),
			ViolationsExpected: []*entity.Violation{},
		},
		{
			Account:            entity.NewAccount(true, 180),
			Transaction:        entity.NewTransaction("Habib's", 20, helper.GetTimeFromString("2021-04-20T19:42:00.000Z")),
			ViolationsExpected: []*entity.Violation{},
		},
		{
			Account:            entity.NewAccount(true, 160),
			Transaction:        entity.NewTransaction("Bob's", 20, helper.GetTimeFromString("2021-04-21T07:04:00.000Z")),
			ViolationsExpected: []*entity.Violation{},
		},
		{
			Account:            entity.NewAccount(true, 140),
			Transaction:        entity.NewTransaction("Subway", 20, helper.GetTimeFromString("2021-04-21T07:15:00.000Z")),
			ViolationsExpected: []*entity.Violation{},
		},
		{
			Account:            entity.NewAccount(true, 120),
			Transaction:        entity.NewTransaction("Domino's Pizza", 20, helper.GetTimeFromString("2021-04-21T09:26:00.000Z")),
			ViolationsExpected: []*entity.Violation{},
		},
	}
}

func getOperationsDataWithNoCardLimitViolations() []ViolationData {
	return []ViolationData{
		{
			Account:            entity.NewAccount(true, 50),
			Transaction:        entity.NewTransaction("Burger King", 20, helper.GetTimeFromString("2021-04-20T19:25:00.000Z")),
			ViolationsExpected: []*entity.Violation{},
		},
		{
			Account:            entity.NewAccount(true, 30),
			Transaction:        entity.NewTransaction("Habib's", 20, helper.GetTimeFromString("2021-04-20T19:42:00.000Z")),
			ViolationsExpected: []*entity.Violation{},
		},
		{
			Account:     entity.NewAccount(true, 10),
			Transaction: entity.NewTransaction("Bob's", 20, helper.GetTimeFromString("2021-04-21T07:04:00.000Z")),
			ViolationsExpected: []*entity.Violation{
				entity.NewViolationInsufficientLimit(),
			},
		},
		{
			Account:     entity.NewAccount(true, 10),
			Transaction: entity.NewTransaction("Subway", 20, helper.GetTimeFromString("2021-04-21T07:15:00.000Z")),
			ViolationsExpected: []*entity.Violation{
				entity.NewViolationInsufficientLimit(),
			},
		},
		{
			Account:     entity.NewAccount(true, 10),
			Transaction: entity.NewTransaction("Domino's Pizza", 20, helper.GetTimeFromString("2021-04-21T09:26:00.000Z")),
			ViolationsExpected: []*entity.Violation{
				entity.NewViolationInsufficientLimit(),
			},
		},
	}
}

func getOperationsDataWithHighTransactionsViolations() []ViolationData {
	return []ViolationData{
		{
			Account:            entity.NewAccount(true, 200),
			Transaction:        entity.NewTransaction("Burger King", 20, helper.GetTimeFromString("2021-04-20T19:25:00.000Z")),
			ViolationsExpected: []*entity.Violation{},
		},
		{
			Account:            entity.NewAccount(true, 180),
			Transaction:        entity.NewTransaction("Habib's", 20, helper.GetTimeFromString("2021-04-20T19:26:00.000Z")),
			ViolationsExpected: []*entity.Violation{},
		},
		{
			Account:     entity.NewAccount(true, 180),
			Transaction: entity.NewTransaction("Bob's", 20, helper.GetTimeFromString("2021-04-20T19:26:00.000Z")),
			ViolationsExpected: []*entity.Violation{
				entity.NewViolationHighFrequencySmallInterval(),
			},
		},
		{
			Account:            entity.NewAccount(true, 160),
			Transaction:        entity.NewTransaction("Subway", 20, helper.GetTimeFromString("2021-04-21T07:15:00.000Z")),
			ViolationsExpected: []*entity.Violation{},
		},
		{
			Account:            entity.NewAccount(true, 140),
			Transaction:        entity.NewTransaction("Domino's Pizza", 20, helper.GetTimeFromString("2021-04-21T09:26:00.000Z")),
			ViolationsExpected: []*entity.Violation{},
		},
	}
}

func getOperationsDataWithDoubleTransactionViolations() []ViolationData {
	return []ViolationData{
		{
			Account:            entity.NewAccount(true, 200),
			Transaction:        entity.NewTransaction("Burger King", 20, helper.GetTimeFromString("2021-04-20T19:25:00.000Z")),
			ViolationsExpected: []*entity.Violation{},
		},
		{
			Account:            entity.NewAccount(true, 180),
			Transaction:        entity.NewTransaction("Habib's", 20, helper.GetTimeFromString("2021-04-20T19:42:00.000Z")),
			ViolationsExpected: []*entity.Violation{},
		},
		{
			Account:            entity.NewAccount(true, 160),
			Transaction:        entity.NewTransaction("Bob's", 20, helper.GetTimeFromString("2021-04-21T07:04:00.000Z")),
			ViolationsExpected: []*entity.Violation{},
		},
		{
			Account:     entity.NewAccount(true, 160),
			Transaction: entity.NewTransaction("Bob's", 20, helper.GetTimeFromString("2021-04-21T07:05:00.000Z")),
			ViolationsExpected: []*entity.Violation{
				entity.NewViolationDoubleTransaction(),
			},
		},
		{
			Account:            entity.NewAccount(true, 140),
			Transaction:        entity.NewTransaction("Domino's Pizza", 20, helper.GetTimeFromString("2021-04-21T09:26:00.000Z")),
			ViolationsExpected: []*entity.Violation{},
		},
	}
}

func getOperationsDataWithMultipleViolations() []ViolationData {
	return []ViolationData{
		{
			Account:            entity.NewAccount(true, 500),
			Transaction:        entity.NewTransaction("Burger King", 20, helper.GetTimeFromString("2021-04-20T19:25:00.000Z")),
			ViolationsExpected: []*entity.Violation{},
		},
		{
			Account:     entity.NewAccount(true, 480),
			Transaction: entity.NewTransaction("Burger King", 20, helper.GetTimeFromString("2021-04-20T19:25:00.000Z")),
			ViolationsExpected: []*entity.Violation{
				entity.NewViolationDoubleTransaction(),
			},
		},
		{
			Account:            entity.NewAccount(true, 480),
			Transaction:        entity.NewTransaction("Bob's", 20, helper.GetTimeFromString("2021-04-21T09:58:00.000Z")),
			ViolationsExpected: []*entity.Violation{},
		},
		{
			Account:            entity.NewAccount(true, 460),
			Transaction:        entity.NewTransaction("Subway", 300, helper.GetTimeFromString("2021-04-21T09:58:00.000Z")),
			ViolationsExpected: []*entity.Violation{},
		},
		{
			Account:     entity.NewAccount(true, 160),
			Transaction: entity.NewTransaction("Pizza Hut", 300, helper.GetTimeFromString("2021-04-21T09:59:00.000Z")),
			ViolationsExpected: []*entity.Violation{
				entity.NewViolationInsufficientLimit(),
				entity.NewViolationHighFrequencySmallInterval(),
			},
		},
		{
			Account:            entity.NewAccount(true, 160),
			Transaction:        entity.NewTransaction("Yogoberry", 60, helper.GetTimeFromString("2021-04-21T20:21:00.000Z")),
			ViolationsExpected: []*entity.Violation{},
		},
		{
			Account:            entity.NewAccount(true, 100),
			Transaction:        entity.NewTransaction("Mcdonald's", 60, helper.GetTimeFromString("2021-04-21T20:21:00.000Z")),
			ViolationsExpected: []*entity.Violation{},
		},
		{
			Account:     entity.NewAccount(true, 40),
			Transaction: entity.NewTransaction("Mcdonald's", 60, helper.GetTimeFromString("2021-04-21T20:23:00.000Z")),
			ViolationsExpected: []*entity.Violation{
				entity.NewViolationInsufficientLimit(),
				entity.NewViolationHighFrequencySmallInterval(),
				entity.NewViolationDoubleTransaction(),
			},
		},
		{
			Account:            entity.NewAccount(true, 40),
			Transaction:        entity.NewTransaction("Starbucks", 10, helper.GetTimeFromString("2021-04-21T20:25:00.000Z")),
			ViolationsExpected: []*entity.Violation{},
		},
	}
}
