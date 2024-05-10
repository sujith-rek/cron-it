package utils

import (
	"strings"
)

var (
	DayOfWeekMap = map[string]string{
		"0": "Sunday",
		"1": "Monday",
		"2": "Tuesday",
		"3": "Wednesday",
		"4": "Thursday",
		"5": "Friday",
		"6": "Saturday",
		"*": "Every",
	}
	MonthMap = map[string]string{
		"1":  "January",
		"2":  "February",
		"3":  "March",
		"4":  "April",
		"5":  "May",
		"6":  "June",
		"7":  "July",
		"8":  "August",
		"9":  "September",
		"10": "October",
		"11": "November",
		"12": "December",
		"*" : "Every",
	}
)


func ExtractNameFromExecString(execString string) string {
	// split the exec string by space
	execStringParts := strings.Split(execString, " ")
	// minute, hour, day, month, day of week

	output := ""

	// minute
	if execStringParts[0] == "*" {
		output += "Every minute"
	} else {
		output += "Every " + execStringParts[0] + " minute"
	}

	// hour
	if execStringParts[1] == "*" {
		output += ", every hour"
	} else {
		output += ", every " + execStringParts[1] + " hour"
	}

	// day
	if execStringParts[2] == "*" {
		output += ", every day"
	} else {
		output += ", every " + execStringParts[2] + " day"
	}

	// month
	if execStringParts[3] == "*" {
		output += ", every month"
	} else {
		output += ", every " + MonthMap[execStringParts[3]] + " month"
	}

	// day of week
	if execStringParts[4] == "*" {
		output += ", every day of the week"
	} else {
		output += ", every " + DayOfWeekMap[execStringParts[4]] + " of the week"
	}

	return output
}