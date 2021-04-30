package helper_test

import (
	"testing"

	"dev-test/nubank-dev-test-2k21/app/helper"
)

func TestGetTimeFromString(t *testing.T) {
	timeStringExpected := "2021-04-30T19:25:00.000Z"
	time := helper.GetTimeFromString(timeStringExpected)
	timeStringGot := time.Format("2006-01-02T15:04:05.000Z")

	if timeStringExpected != timeStringGot {
		t.Errorf("helper.GetTimeFromString() expected %s time, got %s", timeStringExpected, timeStringGot)
	}
}
