package util

import "testing"

func TestStringToTime(t *testing.T) {

	convertedTime := StringToTime("1992-11-07T16:52:00")

	if convertedTime.Year() != 1992 {
		t.Errorf("Conversao do Year incorreta")
	}
	if convertedTime.Month() != 11 {
		t.Errorf("Conversao do Month incorreta")
	}
	if convertedTime.Day() != 07 {
		t.Errorf("Conversao do Day incorreta")
	}
}
