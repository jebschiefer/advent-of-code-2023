package day07

import (
	"slices"
	"strconv"
	"strings"
)

type Hand struct {
	hand     string
	bid      int
	handType HandType
}

type HandType int

const (
	HighCard HandType = iota + 1
	OnePair
	TwoPair
	ThreeOfAKind
	FullHouse
	FourOfAKind
	FiveOfAKind
)

func CalculateTotalWinnings(lines []string) int {
	winnings := 0
	maxRank := len(lines)

	hands := []Hand{}

	for _, line := range lines {
		parts := strings.Split(line, " ")

		hand := parts[0]
		handType := GetHandType(hand)

		bid, err := strconv.Atoi(parts[1])

		if err != nil {
			panic(err)
		}

		hands = append(hands, Hand{
			hand:     hand,
			bid:      bid,
			handType: handType,
		})
	}

	// Sort hands descending (best first)
	slices.SortFunc(hands, func(a Hand, b Hand) int {
		if a.handType < b.handType {
			return 1
		} else if a.handType > b.handType {
			return -1
		}

		for i := range a.hand {
			aCard := GetCardValue(a.hand[i])
			bCard := GetCardValue(b.hand[i])

			if aCard < bCard {
				return 1
			} else if aCard > bCard {
				return -1
			}
		}

		return 0
	})

	for i, hand := range hands {
		rank := maxRank - i
		winning := rank * hand.bid
		winnings += winning
	}

	return winnings
}

func GetHandType(hand string) HandType {
	counts := map[string]int{}
	cards := strings.Split(hand, "")

	for _, card := range cards {
		if _, ok := counts[card]; ok {
			counts[card]++
		} else {
			counts[card] = 1
		}
	}

	countList := []int{}

	for _, count := range counts {
		countList = append(countList, count)
	}

	slices.Sort(countList)
	slices.Reverse(countList)

	highest := countList[0]
	secondHighest := 0

	if len(countList) > 1 {
		secondHighest = countList[1]
	}

	if highest == 5 {
		return FiveOfAKind
	} else if highest == 4 {
		return FourOfAKind
	} else if highest == 3 && secondHighest == 2 {
		return FullHouse
	} else if highest == 3 {
		return ThreeOfAKind
	} else if highest == 2 && secondHighest == 2 {
		return TwoPair
	} else if highest == 2 {
		return OnePair
	}

	return HighCard
}

func GetCardValue(card byte) int {
	switch card {
	case 'A':
		return 14
	case 'K':
		return 13
	case 'Q':
		return 12
	case 'J':
		return 11
	case 'T':
		return 10
	default:
		return int(card) - '0'
	}
}
