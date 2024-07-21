package main


func main() {
	
	var cards deck

	cards = createShuffledDeck()

	handSize := 5
	numPlayers := 4
	
	remainingDeck, playersHands := cards.dealHand(handSize, numPlayers)
	
	remainingDeck.print("cards")

	for _, hand := range playersHands {
		hand.print("hand")
	}
}