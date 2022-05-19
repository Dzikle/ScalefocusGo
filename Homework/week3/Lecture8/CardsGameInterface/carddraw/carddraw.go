package carddraw

import "cardgame/cardgame"

type dealer interface {
	Deal() *cardgame.Card
}

func DrawAllcards(dealer dealer) []cardgame.Card {

	var s []cardgame.Card

	s = append(s, *dealer.Deal())
	return s

}
