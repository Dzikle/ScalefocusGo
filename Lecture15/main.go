package main

import (
	"fmt"
	"log"
	"sort"
	"time"
)

type SortBy []time.Time

func (a SortBy) Len() int           { return len(a) }
func (a SortBy) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a SortBy) Less(i, j int) bool { return a[i].Before(a[j]) }

func sortDates(format string, dates ...string) ([]string, error) {

	var sortDates []string
	var timeDates []time.Time
	//String to Time.Time
	for _, dateString := range dates {
		dateParsed, err := time.Parse(format, dateString)
		if err != nil {
			log.Fatal(err)
		}
		timeDates = append(timeDates, dateParsed)
	}
	// Sorting
	sort.Sort(SortBy(timeDates))
	//Time.Time to String
	for _, date := range timeDates {
		sortDates = append(sortDates, date.String())
	}

	return sortDates, nil
}

func main() {

	dates := []string{"Sep-14-2008", "Dec-03-2021", "Mar-18-2022", "Oct-21-2015"}
	format := "Jan-02-2006"

	sortD, err := sortDates(format, dates...)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(sortD)

}
