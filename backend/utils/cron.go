package utils

import (
	"strconv"
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
	MonthDays = map[string]int{
		"1":  31,
		"2":  29,
		"3":  31,
		"4":  30,
		"5":  31,
		"6":  30,
		"7":  31,
		"8":  31,
		"9":  30,
		"10": 31,
		"11": 30,
		"12": 31,
	}
)


func ExtractNameFromExecString(execString string) string {

	execString = CleanCronString(execString)

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

func CleanCronString(cronString string) string {
	// trim
	cronString = strings.TrimSpace(cronString)
	// replace tabs with space
	cronString = strings.ReplaceAll(cronString, "\t", " ")
	// replace newlines with space
	cronString = strings.ReplaceAll(cronString, "\n", " ")
	// replace multiple spaces with single space
	cronString = strings.Join(strings.Fields(cronString), " ")
	return cronString
}

func ValidateCronString(cronString string) bool {
	
	cronString = CleanCronString(cronString)

	// split the cron string by space
	cronStringParts := strings.Split(cronString, " ")

	// check if the length of the cron string is 5
	if len(cronStringParts) != 5 {
		return false
	}

	// its okay but nope
	if cronStringParts[0] == "*" {
		return false
	} 

	if cronStringParts[1] > "23" {
		return false
	}

	if cronStringParts[3] > "12" {
		return false
	}

	if cronStringParts[4] > "6" {
		return false
	}

	days, err := strconv.Atoi(cronStringParts[2])

	if err != nil {
		return false
	}

	if days > MonthDays[cronStringParts[3]] {
		return false
	}

	return true

}