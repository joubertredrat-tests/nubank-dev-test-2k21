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
		// 	name:           "Test with not initialized account",
		// 	violationsData: getOperationsDataWithNotInitializedAccount(),
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
		ViolationData{
			Account:            entity.NewAccount(true, 100),
			Transaction:        entity.NewTransaction("Burger King", 20, helper.GetTimeFromString("2021-04-20T19:25:00.000Z")),
			ViolationsExpected: []*entity.Violation{},
		},
	}
}

// func getOperationsDataWithNotInitializedAccount() []ViolationData {
// 	return []ViolationData{
// 		ViolationData{
// 			Account:     entity.NewAccountEmpty(),
// 			Transaction: entity.NewTransaction("Burger King", 20, helper.GetTimeFromString("2021-04-20T19:25:00.000Z")),
// 			ViolationsExpected: []*entity.Violation{
// 				&entity.NewViolationAccountNotInitialized(),
// 			},
// 		},
// 	}
// }
