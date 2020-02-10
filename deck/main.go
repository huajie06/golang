package main

import (
	"deck/deck"
	"fmt"
)

// TODO
// package named Card, which will be exported
// function New create a deck of cards

// add sort functions to sort the cards
// option to shuffle the cards
// option to add an arbitrary number of jockers
// option to filter out cards
// option to construct a single deck composed of multiple decks

func main() {
	n := deck.New()
	fmt.Println(n)

	fmt.Println("----------------------------------------")
	deck.ShuffleCard(n)
	fmt.Println(n)

	fmt.Println("----------------------------------------")
	deck.SortCard(n)
	fmt.Println(n)
}
