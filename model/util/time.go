package util

import "time"

const layout = "2006-01-02T15:04:05"

// StringToTime blablabla
func StringToTime(value string) time.Time {
	convertedTime, _ := time.Parse(layout, "1992-11-07T16:52:00")
	return convertedTime
}
