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
			name:           "Test with no violations",
			violationsData: getOperationsDataWithNoViolations(),
			getManager: func() validator.Manager {
				return validator.NewManager(
					[]validator.ValidatorInterface{
						validator.NewAccountNotInitializedValidator(),
					},
				)
			},
		},
		// {
		// 	name:           "Test with not initialized account violations",
		// 	violationsData: getOperationsDataWithNotInitializedAccount(),
		// 	getManager: func() validator.Manager {
		// 		return validator.NewManager(
		// 			[]validator.ValidatorInterface{
		// 				validator.NewAccountNotInitializedValidator(),
		// 			},
		// 		)
		// 	},
		// },
		// {
		// 	name:           "Test with already initialized account violations",
		// 	violationsData: getOperationsDataWithAlreadyInitializedAccount(),
		// 	getManager: func() validator.Manager {
		// 		return validator.NewManager(
		// 			[]validator.ValidatorInterface{
		// 				validator.NewAccountNotInitializedValidator(),
		// 			},
		// 		)
		// 	},
		// },
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			manager := test.getManager()

			for _, violationData := range test.violationsData {
				violationsGot := manager.GetViolations(violationData.Account, violationData.Transaction)

				if !reflect.DeepEqual(violationData.ViolationsExpected, violationsGot) {
					t.Errorf(
						"%s expected %s violations, got %s",
						test.name,
						fmt.Sprintf("%+v", violationData.ViolationsExpected),
						fmt.Sprintf("%+v", violationsGot),
					)
				}
			}
		})
	}
}

func getOperationsDataWithNoViolations() []ViolationData {
	return []ViolationData{
		{
			Account:            entity.NewAccount(true, 100),
			Transaction:        entity.NewTransaction("Burger King", 20, helper.GetTimeFromString("2021-04-20T19:25:00.000Z")),
			ViolationsExpected: []*entity.Violation{},
		},
	}
}

// func getOperationsDataWithNotInitializedAccount() []ViolationData {
// 	return []ViolationData{
// 		{
// 			Account:     entity.NewAccountEmpty(),
// 			Transaction: entity.NewTransaction("Burger King", 20, helper.GetTimeFromString("2021-04-20T19:25:00.000Z")),
// 			ViolationsExpected: []*entity.Violation{
// 				entity.NewViolationAccountNotInitialized(),
// 			},
// 		},
// 		{
// 			Account:     entity.NewAccountEmpty(),
// 			Transaction: entity.NewTransaction("Habib's", 20, helper.GetTimeFromString("2021-04-20T19:42:00.000Z")),
// 			ViolationsExpected: []*entity.Violation{
// 				entity.NewViolationAccountNotInitialized(),
// 			},
// 		},
// 		{
// 			Account:            entity.NewAccount(true, 100),
// 			Transaction:        entity.NewTransaction("Bob's", 20, helper.GetTimeFromString("2021-04-21T07:04:00.000Z")),
// 			ViolationsExpected: []*entity.Violation{},
// 		},
// 		{
// 			Account:            entity.NewAccount(true, 100),
// 			Transaction:        entity.NewTransaction("Subway", 20, helper.GetTimeFromString("2021-04-21T07:15:00.000Z")),
// 			ViolationsExpected: []*entity.Violation{},
// 		},
// 	}
// }

// func getOperationsDataWithAlreadyInitializedAccount() []ViolationData {
// 	return []ViolationData{
// 		{
// 			Account:     entity.NewAccount(true, 100),
// 			Transaction: entity.NewTransaction("Burger King", 20, helper.GetTimeFromString("2021-04-20T19:25:00.000Z")),
// 			ViolationsExpected: []*entity.Violation{
// 				entity.NewViolationAccountNotInitialized(),
// 			},
// 		},
// 		{
// 			Account:     entity.NewAccount(true, 100),
// 			Transaction: entity.NewTransaction("Habib's", 20, helper.GetTimeFromString("2021-04-20T19:42:00.000Z")),
// 			ViolationsExpected: []*entity.Violation{
// 				entity.NewViolationAccountNotInitialized(),
// 			},
// 		},
// 		{
// 			Account:            entity.NewAccount(true, 100),
// 			Transaction:        entity.NewTransaction("Bob's", 20, helper.GetTimeFromString("2021-04-21T07:04:00.000Z")),
// 			ViolationsExpected: []*entity.Violation{},
// 		},
// 		{
// 			Account:            entity.NewAccount(true, 100),
// 			Transaction:        entity.NewTransaction("Subway", 20, helper.GetTimeFromString("2021-04-21T07:15:00.000Z")),
// 			ViolationsExpected: []*entity.Violation{},
// 		},
// 	}
// }
