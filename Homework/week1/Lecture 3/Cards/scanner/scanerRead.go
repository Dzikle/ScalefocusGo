package scanner

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func check(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func Read() (int, int, int, int) {

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter the value of the first card: ")
	firstCardValueS, _ := reader.ReadString('\n')
	firstCardV, err := strconv.ParseFloat(strings.TrimSpace(firstCardValueS), 64)
	check(err)

	fmt.Print("Enter the value of the first card suite  :  Spade - 0 , Diamond - 1 , Heart - 2, Club -3   ")
	firstCardSuiteV, _ := reader.ReadString('\n')
	firstCardS, err := strconv.ParseFloat(strings.TrimSpace(firstCardSuiteV), 64)
	check(err)

	fmt.Print("Enter the value of the Second card: ")
	SecondCardValueV, _ := reader.ReadString('\n')
	SecondCardV, err := strconv.ParseInt(strings.TrimSpace(SecondCardValueV), 0, 64)
	check(err)

	fmt.Print("Enter the value of the first card suite :  Spade - 0 , Diamond - 1 , Heart - 2, Club -3  ")
	SecondCardSuiteS, _ := reader.ReadString('\n')
	SecondCardS, err := strconv.ParseInt(strings.TrimSpace(SecondCardSuiteS), 0, 64)
	check(err)

	return int(firstCardV), int(firstCardS), int(SecondCardV), int(SecondCardS)
}
