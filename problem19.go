package main

import "fmt"

func isLeapYear(n int) bool {
	return n%4 == 0 && (n%100 != 0 || n%400 == 0)
}

func main() {
	
	currentYear := 1900
	currentDay := 1
	currentMonth := 1
	currentWeekDay := 1 // goes from 0 (=Sunday) to 6 (=Saturday)

	numOfSundaysOnFirst := 0

	for currentYear <= 2000 {

		currentDay++
		currentWeekDay = (currentWeekDay + 1)%7


		if currentMonth == 2 {
			if (currentDay == 29 && !isLeapYear(currentYear)) || (currentDay == 30 && isLeapYear(currentYear)) {
				currentDay = 1
				currentMonth = 3	
			}
		} else if currentMonth == 9 || currentMonth == 4 || currentMonth == 6 || currentMonth == 11 {
			if currentDay == 31 {
				currentDay = 1
				currentMonth++
			}
		} else {
			if currentDay == 32 {
				currentDay = 1
				currentMonth++

				if currentMonth == 13 {
					currentMonth = 1
					currentYear++
				}
			}
		}

		if currentWeekDay == 0 && currentDay == 1 && currentYear >= 1901 && currentYear <= 2000 {
			numOfSundaysOnFirst++
		}
	}

	fmt.Println(numOfSundaysOnFirst)
}