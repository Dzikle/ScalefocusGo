package main

//Pull request
import (
	"fmt"
	"time"
)

//For Review
func daysInMonth(month int, year int) (int, int, int, bool) {

	if 0 < month && month < 13 {

		switch month {
		case 4, 6, 9, 11:
			return 30, month, year, true

		case 1, 3, 5, 7, 8, 10, 12:
			return 31, month, year, true

		case 2:
			// Apparently The rule is that if the year is divisible by 100 and not divisible by 400, leap year is skipped
			if year%4 == 0 && year%100 != 0 || year%400 == 0 {
				return 29, month, year, true
			} else {
				return 28, month, year, true
			}
		}
	}
	return 0, month, year, false
}
func Check(err bool) {
	if !err {
		fmt.Println("You entered invalid Date!!!")
	}
}

func main() {

	days, month, year, err := daysInMonth(ReadFromConsole())

	Check(err)

	fmt.Printf("%v in the year %v has %v days", time.Month(month), year, days)

}
