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
		for i := 1; i <= 2; i++ {
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
	mp := map[string]int{"spade": 1, "diamond": 2, "club": 3, "heart": 4}
	sort.Slice(cards, func(i, j int) bool {
		if mp[cards[i].Suit] < mp[cards[j].Suit] {
			return true
		}

		if mp[cards[i].Suit] > mp[cards[j].Suit] {
			return false
		}
		return cards[i].Value < cards[j].Value
	})
}

func DelCard(v int, suit string, cards *[]Card) {
	var ind int
	for i, k := range *cards {
		if k.Suit == suit && v == k.Value {
			ind = i
		}
	}
	*cards = append((*cards)[:ind], (*cards)[ind+1:]...)
}
