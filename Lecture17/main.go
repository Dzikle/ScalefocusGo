package main

import (
	"fmt"
)

type Card struct {
	Value int
	Suite int
}

var MCard Card

func main() {

	card1 := Card{Value: 3, Suite: 2}
	card2 := Card{Value: 14, Suite: 2}
	card3 := Card{Value: 7, Suite: 2}
	card4 := Card{Value: 11, Suite: 2}
	card5 := Card{Value: 11, Suite: 3}

	cardSlice := []Card{card1, card2, card3, card4, card5}
	fmt.Println(cardSlice)
	// MaxCards(cardSlice)

	fmt.Println(MCard)

}
