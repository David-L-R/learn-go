package main

import "fmt"

type deck []string

func (deck deck) print() {
	for i, card := range deck {
		fmt.Println(i, card)
	}
}

func (deck *deck) appendCard (card string)  {
	*deck = append(*deck, card)
}