package main

import (
	"cmp"
	"fmt"
	"log"
	"slices"
	"strconv"
	"strings"

	"github.com/criscokid/aoc2023/internal/fileinput"
)

var cards map[rune]int = map[rune]int{
	'J': 1,
	'2': 2,
	'3': 3,
	'4': 4,
	'5': 5,
	'6': 6,
	'7': 7,
	'8': 8,
	'9': 9,
	'T': 10,
	'Q': 12,
	'K': 13,
	'A': 14,
}

const (
	HighCard = iota
	OnePair
	TwoPair
	ThreeOfAKind
	FullHouse
	FourOfAKind
	FiveOfAKind
)

var textResult map[int]string = map[int]string{
	HighCard:     "HighCard",
	OnePair:      "OnePair",
	TwoPair:      "TwoPair",
	ThreeOfAKind: "ThreeOfAKind",
	FullHouse:    "FullHouse",
	FourOfAKind:  "FourOfAKind",
	FiveOfAKind:  "FiveOfAKind",
}

type hand struct {
	cards    string
	bet      int
	strength int
}

func main() {
	file_path := "input.txt"
	lines, err := fileinput.ReadLines(file_path)
	if err != nil {
		log.Fatal(err)
	}

	hands := []hand{}
	for _, line := range lines {
		hands = append(hands, parseHand(line))
	}

	for i := 0; i < len(hands); i++ {
		str := findStrength(hands[i])
		hands[i].strength = str
	}

	slices.SortFunc(hands, func(a, b hand) int {
		return a.Compare(b)
	})

	total := 0

	for i := 0; i < len(hands); i++ {
		total += hands[i].bet * (i + 1)
	}

	fmt.Println(total)
}

func parseHand(input string) hand {
	parts := strings.Fields(input)
	bet, err := strconv.Atoi(parts[1])
	if err != nil {
		log.Fatal(err)
	}
	return hand{cards: parts[0], bet: bet, strength: 0}
}

func findStrength(h hand) int {
	counts := charCounts(h.cards)
	jCount, hasJ := counts['J']
	//Five of a kind
	if len(counts) == 1 {
		return FiveOfAKind
	}
	//high card
	if len(counts) == 5 {
		if hasJ {
			return OnePair
		}
		return HighCard
	}

	if len(counts) == 4 {
		if hasJ {
			return ThreeOfAKind
		}

		return OnePair
	}

	if len(counts) == 3 {
		for _, v := range counts {
			if v == 3 {
				if hasJ {
					return FourOfAKind
				}
				return ThreeOfAKind
			}
		}

		if hasJ {
			if jCount == 2 {
				return FourOfAKind
			}
			return FullHouse
		}

		return TwoPair
	}

	if len(counts) == 2 {
		for _, v := range counts {
			if v == 4 {
				if hasJ {
					return FiveOfAKind
				}
				return FourOfAKind
			}
		}
		if hasJ {
			return FiveOfAKind
		}
		return FullHouse
	}

	return HighCard
}

func charCounts(input string) map[rune]int {
	charMap := map[rune]int{}
	for _, r := range input {
		_, ok := charMap[r]
		if ok {
			charMap[r] += 1
		} else {
			charMap[r] = 1
		}
	}
	return charMap
}

func (h hand) Compare(other hand) int {
	cmpValue := cmp.Compare(h.strength, other.strength)
	if cmpValue != 0 {
		return cmpValue
	}

	for i := 0; i < len(h.cards); i++ {
		hValue := cards[rune(h.cards[i])]
		oValue := cards[rune(other.cards[i])]
		if hValue < oValue {
			return -1
		} else if hValue > oValue {
			return 1
		}
	}

	return 0
}
