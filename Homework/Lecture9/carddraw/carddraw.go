package carddraw

import (
	"cardgame/cardgame"
	"errors"
)

type dealer interface {
	Deal() *cardgame.Card
}

func DrawAllcards(dealer dealer) ([]cardgame.Card, error) {

	var s []cardgame.Card
	c := *dealer.Deal()
	if c.Value != "" {
		s = append(s, *dealer.Deal())
		return s, nil
	} else {
		err := errors.New("The Deck is empty")
		return s, err
	}
}
