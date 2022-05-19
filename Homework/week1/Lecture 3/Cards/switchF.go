package main

import "fmt"

func CardSwitchComp(value int) {

	switch {
	case value == 0:
		fmt.Println("The cards are equal from Switch")
	case value == 1:
		fmt.Println("The first card is greater from Switch")
	case value == -1:
		fmt.Println("The second card is greater from Switch")
	default:
		fmt.Println("Invalid")
	}

}
