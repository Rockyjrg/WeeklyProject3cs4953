# WeeklyProject3cs4953

Create a program to handle a deck of cards in Go.

Functions you need to complete in your deck_of_cards.go file.

func NewDeck() \*CardDeck //returns a standard deck with 52 cards in order ascending order (do all clubs before diamonds before hearts before spades)

(d \*CardDeck) Shuffle() //randomizes the order of cards currently in the deck
Come up with your own way of doing this! It doesn’t need to be perfect.

(d \*CardDeck) Contains(card Card) bool //returns true if card is in the deck, false if not.

(d \*CardDeck) DrawTop() Card //Removes top card from the deck and returns it

(d \*CardDeck) DrawBottom() Card //Removes bottom card from the deck and returns it

(d \*CardDeck) DrawRandom() Card //Removes and returns card from a random position in the deck.

(d \*CardDeck) CardToTop(card Card) //Place a card on top of the deck.

(d \*CardDeck) CardToBottom(card Card)//Place a card on the bottom of the deck.

(d \*CardDeck) CardToRandom(card Card) //Place a card somewhere random on the deck.

(d \*CardDeck) CardsLeft() int //Returns number of cards left in the deck
