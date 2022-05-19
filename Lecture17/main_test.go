package main

import (
	card "card/cardgame"
	"testing"
)

func TestMaxCards(t *testing.T) {
	card1 := card.Card{Value: 3, Suite: 2}
	card2 := card.Card{Value: 4, Suite: 2}
	cardSlice := []card.Card{card1, card2}

	expectedCard := card2
	result := card.MaxCards(cardSlice)

	if result != expectedCard {
		t.Errorf("The Biger card should be %d, not %d ", expectedCard, result)
	}
}

func TestCompareCards(t *testing.T) {
	card1 := card.Card{Value: 3, Suite: 2}
	card2 := card.Card{Value: 4, Suite: 2}

	expectedCard := -1
	result := card.CompareCards(card1, card2)
	if result != expectedCard {
		t.Errorf("The Biger card should be %d, not %d ", expectedCard, result)
	}

}
