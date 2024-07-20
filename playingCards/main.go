package main

import "fmt"


var	card string = "Ace of Spades"

func main() {
	// var card string = "Ace of Spades"
	// card = "Five of Diamonds"

	cards := deck{"Ace of Spades", "Five of Diamonds"}

	for i, card := range cards {
		fmt.Println(i, card)
	}

	cards = append(cards, "Six of Clubs")

	fmt.Println(card)
}

func newCard() string {
	return "Five of Diamonds"
}