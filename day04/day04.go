package day04

import (
	"math"
	"regexp"
	"strconv"
	"strings"
)

type Card struct {
	label                string
	matchingNumbers      []int
	matchingNumbersCount int

	next *Card
}

type CardList struct {
	head  *Card
	tail  *Card
	count int
}

func (cardList *CardList) Append(label string, matchingNumbers []int) {
	card := &Card{
		label:                label,
		matchingNumbers:      matchingNumbers,
		matchingNumbersCount: len(matchingNumbers),
	}

	if cardList.head == nil {
		cardList.head = card
		cardList.count = 1
	} else {
		current := cardList.tail
		current.next = card
		cardList.count++
	}

	cardList.tail = card
}

func (cardList *CardList) Get(label string) *Card {
	current := cardList.head

	if current.label == label {
		return current
	}

	for current.next != nil {
		current = current.next

		if current.label == label {
			return current
		}
	}

	return nil
}

func GetPointsPerCard(lines []string) []int {
	points := []int{}

	ProcessCards(lines, func(card *Card) {
		cardPoints := calculatePoints(card.matchingNumbers)
		points = append(points, cardPoints)
	})

	return points
}

func CountInstancesOfCards(lines []string) int {
	cards := CardList{}

	ProcessCards(lines, func(card *Card) {
		cards.Append(card.label, card.matchingNumbers)
	})

	currentCard := cards.head

	for currentCard != nil {
		cardToAdd := cards.Get(currentCard.label)

		for i := 0; i < currentCard.matchingNumbersCount; i++ {
			cardToAdd = cardToAdd.next

			if cardToAdd != nil {
				cards.Append(cardToAdd.label, cardToAdd.matchingNumbers)
			}
		}

		currentCard = currentCard.next
	}

	return cards.count
}

func ProcessCards(lines []string, handleCard func(card *Card)) {
	for _, line := range lines {
		matchingNumbers := []int{}

		cardLabel := getCardLabel(line)
		cardNumberPattern := regexp.MustCompile(`^Card\s+\d+:`)
		cardValue := cardNumberPattern.ReplaceAllString(line, "")
		cardParts := strings.Split(cardValue, "|")

		re := regexp.MustCompile(`\d+`)

		winningNumbersList := re.FindAllString(cardParts[0], -1)
		ourNumbersList := re.FindAllString(cardParts[1], -1)

		winningNumbers := sliceToMap(winningNumbersList)

		for _, ourNumber := range ourNumbersList {
			if matchingNumber, ok := winningNumbers[ourNumber]; ok {
				matchingNumbers = append(matchingNumbers, matchingNumber)
			}
		}

		card := &Card{
			label:           cardLabel,
			matchingNumbers: matchingNumbers,
		}

		handleCard(card)
	}
}

func sliceToMap(slice []string) map[string]int {
	valueMap := map[string]int{}

	for _, value := range slice {
		number, err := strconv.Atoi(value)

		if err == nil {
			valueMap[value] = number
		}
	}

	return valueMap
}

func calculatePoints(numbers []int) int {
	count := len(numbers)

	if count < 2 {
		return count
	}

	return int(math.Pow(2, float64(count-1)))
}

func getCardLabel(line string) string {
	re := regexp.MustCompile(`^Card\s+\d+`)
	return re.FindString(line)
}
