package deck

import (
	"math/rand"
	"sort"
	"time"
)

// value is the face value of the card
// suit is spade, diamond, clubs, hearts
type Card struct {
	Value int
	Suit  string
}

func New() []Card {
	var ret = []Card{}
	var suit = []string{"spade", "diamond", "club", "heart"}
	for _, v := range suit {
		for i := 1; i <= 3; i++ {
			ret = append(ret, Card{i, v})
		}
	}
	return ret
}

func ShuffleCard(cards []Card) {
	rand.Seed(time.Now().Unix())
	rand.Shuffle(len(cards), func(i, j int) {
		cards[i], cards[j] = cards[j], cards[i]
	})
}

func SortCard(cards []Card) {
	// sort.Slice(cards, func(i, j int) bool {
	// 	switch {
	// 	case cards[j].Suit == "spade" && cards[i].Suit == "spade" && cards[i].Value > cards[j].Value:
	// 		return true
	// 	case cards[j].Suit == "diamond" && cards[i].Suit == "diamond" && cards[i].Value > cards[j].Value:
	// 		return true
	// 	case cards[j].Suit == "club" && cards[i].Suit == "club" && cards[i].Value > cards[j].Value:
	// 		return true
	// 	case cards[j].Suit == "heart" && cards[i].Suit == "heart" && cards[i].Value > cards[j].Value:
	// 		return true
	// 	default:
	// 		return false
	// 	}

	// })

	sort.Slice(cards, func(i, j int) bool {
		return cards[i].Suit > cards[j].Suit
	})

}
