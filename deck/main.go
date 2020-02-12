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
// option to filter out cards

// option to add an arbitrary number of jockers
// option to construct a single deck composed of multiple decks

func main() {
	fmt.Println("------------New Card--------------------")
	n := deck.New()
	fmt.Println(n)
	fmt.Println("----------shuffled card-----------------")
	deck.ShuffleCard(n)
	fmt.Println(n)
	fmt.Println("----------sorted card-------------------")
	deck.SortCard(n)
	fmt.Println(n)
	fmt.Println("---------deleted card-------------------")
	deck.ShuffleCard(n)
	fmt.Println(n)
	deck.DelCard(1, "spade", &n)
	fmt.Println(n)
	fmt.Println("---------re-sort card-------------------")
	deck.SortCard(n)
	fmt.Println(n)
}
