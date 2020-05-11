package poll

import (
	"fmt"
	"os"
	"utils"
)

const (
	tooManyArgs = "There are too many arguments.\n"
	notEnoughArgs = "There are not enough arguments.\n"
	mustBePositiveInteger = "must be a positive integer.\n"
	mustBeGreaterThanZero = "must be greater than zero.\n"
	percentageMustBeBetween = "The percentage of voting intentions must be between 0 and 100.\n"
	maxArg = 3
	minArg = 3
)

var a sArg

type sArg struct {
	pSize int
	sSize int
	p float64
}

func printError(errorMessage string) {
	fmt.Printf("Error: %s\n", errorMessage)
}

func printErrorWithValue(valueName string, errorMessage string) {
	fmt.Printf("Error: '%s' %s\n", valueName, errorMessage)
}

// CheckHelp arg -h
func CheckHelp() bool {
	argsWithoutProg := os.Args[1:]

	for _, arg := range argsWithoutProg {
		if arg == "-h" {
			return true
		}
	}
	return false
}

func getIntegerPositiveValueGreaterThanZero(valueName string, arg string) (bool, int) {
	if !utils.IsInteger(arg) {
		printErrorWithValue(valueName, mustBePositiveInteger)
		return false, -1
	}
	integer := utils.ConvertStringToInt(arg)
	if integer <= 0 {
		printErrorWithValue(valueName, mustBeGreaterThanZero)
		return false, -1
	}
	return true, integer
}

func getPercentageValue(arg string) (bool, float64) {
	if !utils.IsPercentageValue(arg) {
		printError(percentageMustBeBetween)
		return false, -1
	}
	float := utils.ConvertStringToFloat(arg)
	return true, float
}

// CheckArgs check user input's args
func CheckArgs() bool {
	argsWithoutProg := os.Args[1:]

	if len(argsWithoutProg) < minArg {
		printError(notEnoughArgs)
		return false
	}
	if len(argsWithoutProg) > maxArg {
		printError(tooManyArgs)
		return false
	}
	valueNames := [9]string{"pSize", "sSize", "p"}
	for i, arg := range argsWithoutProg {
		valueName := valueNames[i]

		// Check and assign pSize
		if valueName == "pSize" {
			status, integer := getIntegerPositiveValueGreaterThanZero(valueName, arg)
			if !status {
				return false
			}
			a.pSize = integer
		}

		// Check and assign sSize
		if valueName == "sSize" {
			status, float := getIntegerPositiveValueGreaterThanZero(valueName, arg)
			if !status {
				return false
			}
			a.sSize = float
		}

		// Check and assign p
		if valueName == "p" {
			status, float := getPercentageValue(arg)
			if !status {
				return false
			}
			a.p = float
		}
	}
	return true
}