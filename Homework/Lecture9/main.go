package main

import (
	"cardgame/carddraw"
	"cardgame/cardgame"
	"fmt"
	"log"
)

func main() {

	Deck := cardgame.NewDeck()
	length := len(Deck.CardDeck)

	var s []cardgame.Card
	for i := 0; i < length; i++ {
		ds, err := carddraw.DrawAllcards(&Deck)
		if err == nil {
			s = append(s, ds...)
			Deck.CardDeck = append(Deck.CardDeck[:0], Deck.CardDeck[1:]...)
		} else {
			log.Fatal(fmt.Errorf(err.Error()))
		}
	}
	fmt.Println(s)

}
