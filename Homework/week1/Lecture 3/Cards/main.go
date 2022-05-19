package main

import (
	"card/scanner"
	"fmt"
)

type Cardsuite = int

const (
	club Cardsuite = iota
	diamond
	heart
	spade
)
//for pull request
func compareCards(cardOneVal, cardOneSuit, CardTwoVal, CardTwoSuit int) int {
	result := 0
	if cardOneVal > CardTwoVal {
		result = 1
		fmt.Println()
		return result
	}
	if cardOneVal == CardTwoVal && cardOneSuit > cardOneVal {
		result = 1
		return result
	} else if cardOneVal == CardTwoVal && cardOneSuit == cardOneVal {
		result = 0
		return result
	} else {
		result = -1
		return result
	}
}
func ResultPrint(cardCompar int) {
	if cardCompar > 0 {
		fmt.Println("The First Card is greater")
	} else if cardCompar < 0 {
		fmt.Println("The Second Card is greater")
	} else {
		fmt.Println("The cards are equal")
	}
}
func check(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
func main() {

	ResultPrint(compareCards(scanner.Read()))

	CardSwitchComp(compareCards(scanner.Read()))

}
