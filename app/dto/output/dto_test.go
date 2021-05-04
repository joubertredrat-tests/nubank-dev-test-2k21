package output_test

import (
	"reflect"
	"testing"

	"dev-test/nubank-dev-test-2k21/app/dto/output"
)

func TestAccountLine(t *testing.T) {
	activeCardExpected := true
	availableLimitExpected := uint(120)
	violationsExpected := []string{"my-violation"}

	accountLine := output.NewAccountLine(activeCardExpected, availableLimitExpected, violationsExpected)

	if activeCardExpected != accountLine.Account.ActiveCard {
		t.Errorf("accountLine.Account.ActiveCard expected %v, got %v", activeCardExpected, accountLine.Account.ActiveCard)
	}

	if availableLimitExpected != accountLine.Account.AvailableLimit {
		t.Errorf("accountLine.Account.AvailableLimit expected %v, got %v", availableLimitExpected, accountLine.Account.AvailableLimit)
	}

	if !reflect.DeepEqual(violationsExpected, accountLine.Account.Violations) {
		t.Errorf("accountLine.Account.Violations expected %v, got %v", violationsExpected, accountLine.Account.Violations)
	}
}
