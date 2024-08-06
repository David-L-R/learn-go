package main

import (
	"fmt"
)

func main() {
    var cards deck
    
    cards = cards.fillInDeck()

    err := cards.saveToFile("cards.txt")
    if err != nil {
        fmt.Println("Error saving to file:", err)
    } else {
        fmt.Println("Deck saved to file successfully")
    }
}

