package helper

import "time"

func GetTimeFromString(timeString string) time.Time {
	timeLayout := "2006-01-02T15:04:05.000Z"
	time, _ := time.Parse(timeLayout, timeString)
	return time
}
