package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ReadFromConsole() (month, year int) {

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter the number of month betwee 1 - 12: ")
	monthS, _ := reader.ReadString('\n')
	m, _ := strconv.ParseInt(strings.TrimSpace(monthS), 10, 32)
	month = int(m)

	fmt.Print("Enter the year: ")
	yearS, _ := reader.ReadString('\n')
	y, _ := strconv.ParseInt(strings.TrimSpace(yearS), 10, 32)
	year = int(y)
	return month, year
}
