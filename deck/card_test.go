package deck

import (
	"fmt"
	"testing"
)

func ExampleCard() {
	fmt.Println(Card{Rank: Ace, Suit: Heart})
	fmt.Println(Card{Rank: Two, Suit: Spade})
	fmt.Println(Card{Rank: Nine, Suit: Diamond})
	fmt.Println(Card{Rank: Jack, Suit: Club})

	// Output:
	// Ace of Hearts
	// Two of Spades
	// Nine of Diamonds
	// Jack of Clubs
}

func TestNew(t *testing.T) {
	cards := New()
	// 13 ranks * 4 suits
	if len(cards) != 13*4 {
		t.Error("Wrong number of cards in a new deck.")
	}
}

func TestAscSort(t *testing.T) {
	cards := New(AscSort)
	exp := Card{Rank: Ace, Suit: Spade}
	if cards[0] != exp {
		t.Error("Expected Ace of Spades as first card. Received:", cards[0])
	}
}

func TestSort(t *testing.T) {
	// Sort descending
	var dscSort = func(cards []Card) func(i, j int) bool {
		return func(i, j int) bool {
			return absoluteRank(cards[i]) > absoluteRank(cards[j])
		}

	}
	cards := New(Sort(dscSort))
	exp := Card{Rank: King, Suit: Heart}
	if cards[0] != exp {
		t.Error("Expected King of Hearts as first card. Received:", cards[0])
	}

}

func TestShuffle(t *testing.T) {
	cards := New(AscSort)
	// deterministic shuffle
	shuffleSeed = 0
	Shuffle(cards)

	first := Card{Suit: Diamond, Rank: Seven}
	second := Card{Suit: Club, Rank: Jack}

	if cards[0] != first {
		t.Errorf("Wanted %v, got %v instead", cards[0], first)
	}

	if cards[1] != second {
		t.Errorf("Wanted %v, got %v instead", cards[1], second)
	}
}

func TestExists(t *testing.T) {
	cards := New(AscSort)
	want := Card{Rank: Ace, Suit: Spade}
	if !Exists(want, cards) {
		t.Errorf("Card %v not found in deck: %v", want, cards)
	}
}

func TestFilter(t *testing.T) {
	want := Card{Rank: Ace, Suit: Spade}
	filterFunc := func(card Card) bool {
		return card == want
	}
	cards := New(Filter(filterFunc))
	if Exists(want, cards) {
		t.Errorf("Card %v should be removed from deck: %v", want, cards)
	}

}
