package day04

import (
	"math"
	"regexp"
	"strconv"
	"strings"
)

func GetPointsPerCard(lines []string) []int {
	points := []int{}

	for _, line := range lines {
		ourWinningNumbers := []int{}

		cardNumberPattern := regexp.MustCompile(`^Card\s+\d+:`)
		line = cardNumberPattern.ReplaceAllString(line, "")

		cardParts := strings.Split(line, "|")

		re := regexp.MustCompile(`\d+`)

		winningNumbersList := re.FindAllString(cardParts[0], -1)
		ourNumbersList := re.FindAllString(cardParts[1], -1)

		winningNumbers := sliceToMap(winningNumbersList)

		for _, ourNumber := range ourNumbersList {
			if winningNumber, ok := winningNumbers[ourNumber]; ok {
				ourWinningNumbers = append(ourWinningNumbers, winningNumber)
			}
		}

		cardPoints := calculatePoints(ourWinningNumbers)

		points = append(points, cardPoints)
	}

	return points
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
