package cardgame

// For review
type Card struct {
	Value int
	Suite int
}

var MCard Card

func CompareCards(cardOne Card, cardTwo Card) int {

	if (cardOne.Value > cardTwo.Value) || (cardOne.Value == cardTwo.Value && cardOne.Suite > cardTwo.Suite) {
		return 1
	} else if (cardOne.Value < cardTwo.Value) || (cardOne.Value == cardTwo.Value && cardOne.Suite < cardTwo.Suite) {
		return -1
	} else {
		return 0
	}
}

func MaxCards(cards []Card) Card {

	for _, v := range cards {
		if CompareCards(v, MCard) > 0 {
			MCard = v
		} else {
			continue
		}
	}
	return MCard
}
