package main

import "fmt"


var	card string = "Ace of Spades"

func main() {
	cards := deck{"Ace of Spades", "Five of Diamonds"}

	cards.print()

	cards.appendCard("Six of Clubs")

	fmt.Println(cards)
}

func newCard() string {
	return "Five of Diamonds"
}