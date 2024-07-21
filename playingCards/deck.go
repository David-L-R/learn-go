package main

import (
	"fmt"
	"math/rand"
	"time"
)


type deck []string

func (deck deck) print() {
	for i, card := range deck {
		fmt.Println(i, card)
	}
}

func (deck *deck) appendCard (card string)  {
	*deck = append(*deck, card)
}

func createShuffledDeck() []string {
	suits := []string{"♠", "♣", "♥", "♦"}
	ranks := []string{"A", "2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K"}

	var deck []string
	
	// Populate the deck
	for _, suit := range suits {
		for _, rank := range ranks {
			card := rank + suit
			deck = append(deck, card)
		}
	}
	
	//Seed the random number generator
	rand.Seed(time.Now().UnixNano())

	// Shuffle the deck
	rand.Shuffle(len(deck), func(i, j int) {
		deck[i], deck[j] = deck[j], deck[i]
	})
	
	return deck
}

