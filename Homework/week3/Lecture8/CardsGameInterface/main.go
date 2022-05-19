package main

import (
	"cardgame/carddraw"
	"cardgame/cardgame"
	"fmt"
)

func main() {

	Deck := cardgame.NewDeck()
	length := len(Deck.CardDeck)
	var s []cardgame.Card
	for i := 0; i < length; i++ {
		s = append(s, carddraw.DrawAllcards(&Deck)...)
		Deck.CardDeck = append(Deck.CardDeck[:0], Deck.CardDeck[1:]...)
	}

	fmt.Println(s)

}
