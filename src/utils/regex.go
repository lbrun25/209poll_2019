package utils

import (
	"regexp"
)

// IsFloat - check if the value is a float value
func IsFloat(value string) bool {
	var re = regexp.MustCompile("[+-]?([0-9]*[.])?[0-9]+")

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