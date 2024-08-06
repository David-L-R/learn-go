package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

type deck []string

func createShuffledDeck() deck {
    suits := []string{"♠", "♣", "♥", "♦"}
    ranks := []string{"A", "2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K"}

    var d deck

    // Populate the deck
    for _, suit := range suits {
        for _, rank := range ranks {
            card := rank + suit
            d = append(d, card)
        }
    }

    d = d.shuffleDeck()

    return d
}

func (d deck) fillInDeck() deck {
    suits := []string{"♠", "♣", "♥", "♦"}
    ranks := []string{"A", "2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K"}

    for _, suit := range suits {
        for _, rank := range ranks {
            card := rank + suit
            d = append(d, card)
        }
    }

    // print deck
    for i, card := range d {
        fmt.Println(i, card)
    }

    return d
}

func (d deck) print(name string) {
    if len(name) > 0 {
        fmt.Println(name)
    }
    for i, card := range d {
        fmt.Println(i, card)
    }
}

func (d deck) shuffleDeck() deck {
    // Seed the random number generator
    rand.Seed(time.Now().UnixNano())

    // Shuffle the deck
    rand.Shuffle(len(d), func(i, j int) {
        d[i], d[j] = d[j], d[i]
    })
    return d
}

func (d deck) dealHand(handSize int, numPlayers int) (deck, []deck) {
    var players = make([]deck, numPlayers)

    for i := 0; i < numPlayers; i++ {
        players[i] = d[:handSize]
        d = d[handSize:]
    }

    return d, players
}

func (d deck) saveToFile(filename string) error {
    return os.WriteFile("./"+filename, []byte(strings.Join([]string(d), ",")), 0644)
}
