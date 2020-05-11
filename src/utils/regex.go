package utils

import (
	"regexp"
)

// IsPercentageValue - check if the value is a percentage value
func IsPercentageValue(value string) bool {
	var re = regexp.MustCompile("((100)|(\\d{1,2}(.\\d*)?))")

	match := re.FindString(value)
	if len(value) != len(match) {
		return false
	}
	return true
}

// IsInteger - check if the value is a positive integer
func IsInteger(arg string) bool {
	var re = regexp.MustCompile("[-+]?\\d+")

	match := re.FindString(arg)
	if len(arg) != len(match) {
		return false
	}
	return true
}