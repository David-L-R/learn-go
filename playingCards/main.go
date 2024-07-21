package main

import "fmt"


func main() {
	
	var cards deck

	cards = createShuffledDeck()

	cards.saveToFile("cards")

	handSize := 5
	numPlayers := 4
	
	remainingDeck, playersHands := cards.dealHand(handSize, numPlayers)
	
	remainingDeck.print("cards")

	for _, hand := range playersHands {
		hand.print("hand")
	}

	err := remainingDeck.saveToFile("remaining_deck.txt")
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Deck saved to file successfully")
	}
}