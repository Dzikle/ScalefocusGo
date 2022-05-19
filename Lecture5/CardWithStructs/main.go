package main

// For review
import "fmt"

type Card struct {
	Value int
	Suite int
}

var mCard Card

func compareCards(cardOne Card, cardTwo Card) int {

	if (cardOne.Value > cardTwo.Value) || (cardOne.Value == cardTwo.Value && cardOne.Suite > cardTwo.Suite) {
		return 1
	} else if (cardOne.Value < cardTwo.Value) || (cardOne.Value == cardTwo.Value && cardOne.Suite < cardTwo.Suite) {
		return -1
	} else {
		return 0
	}
}

func maxCards(cards []Card) Card {

	for _, v := range cards {
		if compareCards(v, mCard) > 0 {
			mCard = v
		} else {
			continue
		}
	}
	return mCard
}

func main() {

	card1 := Card{Value: 3, Suite: 2}
	card2 := Card{Value: 14, Suite: 2}
	card3 := Card{Value: 7, Suite: 2}
	card4 := Card{Value: 11, Suite: 2}
	card5 := Card{Value: 11, Suite: 3}

	cardSlice := []Card{card1, card2, card3, card4, card5}

	maxCards(cardSlice)

	fmt.Println(mCard)

}
