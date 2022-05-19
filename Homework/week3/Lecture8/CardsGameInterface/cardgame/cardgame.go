package cardgame

type Card struct {
	Value string
	Suite string
}
type Deck struct {
	CardDeck []Card
}

func NewDeck() Deck {
	var InitDeck []Card
	Suite := []string{"diamonds", "spades", "hearts", "clubs"}
	Values := []string{"A", "2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K"}

	for _, s := range Suite {
		for _, v := range Values {
			Card := Card{Value: v, Suite: s}
			InitDeck = append(InitDeck, Card)
			// fmt.Println(Card)
		}

	}
	NewDeck := Deck{CardDeck: InitDeck}
	return NewDeck

}

func (d *Deck) Deal() *Card {
	c := d.CardDeck[0]

	return &c
}
