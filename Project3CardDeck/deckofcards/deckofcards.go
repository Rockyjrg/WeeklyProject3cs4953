package deckofcards

import (
	"math/rand/v2"
)

// globals for suits and values
// go lets us declare these even if we don't use them jokerface.png
var suits []string = []string{"Clubs", "Diamonds", "Hearts", "Spades"}
var values []string = []string{"2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K", "A"}

// Card represents a single playing card
type Card struct {
	Suit  string
	Value string
}

// CardDeck represents a deck of playing cards
type CardDeck struct {
	Cards []Card
}

// NewDeck initializes a new deck of cards in standard order
func NewDeck() *CardDeck {
	cards := []Card{}
	for i := 0; i < len(suits); i++ {

		for j := 0; j < len(values); j++ {
			card := Card{
				Suit:  suits[i],
				Value: values[j],
			}

			cards = append(cards, card)
		}

	}
	return &CardDeck{Cards: cards}
}

// Shuffle randomizes the order of the cards in the deck
func (d *CardDeck) Shuffle() {
	for i := len(d.Cards) - 1; i > 0; i-- {
		randomize := rand.IntN(i + 1)
		d.Cards[i], d.Cards[randomize] = d.Cards[randomize], d.Cards[i]
	}
}

// Contains checks if the deck contains a specific card
func (d *CardDeck) Contains(card Card) bool {
	for _, currentCard := range d.Cards {
		if card == currentCard {
			return true
		}
	}
	return false
}

// DrawTop removes and returns the top card from the deck
func (d *CardDeck) DrawTop() Card {
	firstCard := d.Cards[0]
	d.Cards = d.Cards[1:]
	return firstCard
}

// DrawBottom removes and returns the bottom card from the deck
func (d *CardDeck) DrawBottom() Card {
	lastCard := d.Cards[len(d.Cards)-1]
	d.Cards = d.Cards[:len(d.Cards)-1]
	return lastCard
}

// DrawRandom removes and returns a card from a random position in the deck
func (d *CardDeck) DrawRandom() Card {
	index := rand.IntN(len(d.Cards))
	randomCard := d.Cards[index]
	d.Cards = append(d.Cards[:index], d.Cards[index+1:]...)
	return randomCard
}

// CardToTop places a card on top of the deck
func (d *CardDeck) CardToTop(card Card) {
	//increase size with a dummy element
	d.Cards = append(d.Cards, Card{})

	//shift elements to the right by one
	copy(d.Cards[1:], d.Cards[0:len(d.Cards)-1])

	//place new card at index[0]
	d.Cards[0] = card
}

// CardToBottom places a card on the bottom of the deck
func (d *CardDeck) CardToBottom(card Card) {

	d.Cards = append(d.Cards, card)

}

// CardToRandom places a card at a random position in the deck
func (d *CardDeck) CardToRandom(card Card) {
	index := rand.IntN(len(d.Cards))
	d.Cards = append(d.Cards, Card{})
	copy(d.Cards[index+1:], d.Cards[index:])
	d.Cards[index] = card
}

// CardsLeft returns the number of cards left in the deck
func (d *CardDeck) CardsLeft() int {
	return len(d.Cards)
}
