package blackjack

var Lookup = map[string]int{
	"ace":   11,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
	"ten":   10,
	"jack":  10,
	"queen": 10,
	"king":  10,
	"other": 0,
}

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	return Lookup[card]
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	heldTotal := ParseCard(card1) + ParseCard(card2)
	dealerVal := ParseCard(dealerCard)

	// If you have a pair of aces, always split
	if card1 == "ace" && card2 == "ace" {
		return "P"
	}

	// If you have blackjack (21) and dealer does NOT have ace, 10, J, Q, K, you automatically win
	if heldTotal == 21 {
		if dealerVal < 10 {
			return "W"
		}
		// If dealer has ace, 10, J, Q, or K, stand and wait
		return "S"
	}

	// If cards total 17 to 20, always stand
	if heldTotal >= 17 && heldTotal <= 20 {
		return "S"
	}

	// If cards total 12 to 16, stand unless dealer has 7 or higher, then hit
	if heldTotal >= 12 && heldTotal <= 16 {
		if dealerVal >= 7 {
			return "H"
		}
		return "S"
	}

	// If cards total 11 or less, always hit
	if heldTotal <= 11 {
		return "H"
	}

	// Default fallback (shouldn't be reached)
	return "S"
}
