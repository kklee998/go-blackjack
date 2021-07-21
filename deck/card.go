//go:generate stringer -type=Suit,Rank

package deck

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

type Suit int
type Rank int
type Card struct {
	Suit
	Rank
}

const (
	Spade Suit = iota
	Diamond
	Club
	Heart
)

var suits = [...]Suit{Spade, Diamond, Club, Heart}

const (
	_ Rank = iota
	Ace
	Two
	Three
	Four
	Five
	Six
	Seven
	Eight
	Nine
	Ten
	Jack
	Queen
	King
)

const (
	minRank = Ace
	maxRank = King
)

func (c Card) String() string {
	return fmt.Sprintf("%s of %ss", c.Rank.String(), c.Suit.String())
}

func New(opts ...func([]Card) []Card) []Card {
	var cards []Card
	for _, suit := range suits {
		for rank := minRank; rank <= maxRank; rank++ {
			cards = append(cards, Card{Suit: suit, Rank: rank})
		}
	}

	for _, opt := range opts {
		cards = opt(cards)
	}
	return cards

}

func AscSort(cards []Card) []Card {
	sort.Slice(cards, Less(cards))
	return cards
}

func Sort(less func(cards []Card) func(i, j int) bool) func([]Card) []Card {
	return func(cards []Card) []Card {
		sort.Slice(cards, less(cards))
		return cards
	}
}

func Less(cards []Card) func(i, j int) bool {
	return func(i, j int) bool {
		return absoluteRank(cards[i]) < absoluteRank(cards[j])
	}
}

func absoluteRank(c Card) int {
	return int(c.Suit)*int(maxRank) + int(c.Rank)
}

var shuffleSeed = time.Now().Unix()

func Shuffle(cards []Card) {
	rand.Seed(shuffleSeed)
	rand.Shuffle(len(cards), func(i, j int) {
		cards[i], cards[j] = cards[j], cards[i]
	})
}

func Filter(f func(card Card) bool) func([]Card) []Card {
	return func(cards []Card) []Card {
		var r []Card
		for _, c := range cards {
			if !f(c) {
				r = append(r, c)
			}
		}
		return r
	}
}

func Exists(card Card, cards []Card) bool {
	for _, c := range cards {
		if c == card {
			return true
		}
	}
	return false
}
