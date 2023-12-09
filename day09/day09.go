package day09

import (
	"aoc2023/utilities"
	"slices"
)

func GetPredictions(lines []string, reverse bool) []int {
	predictions := []int{}

	for _, line := range lines {
		numbers := utilities.StringToInts(line)

		if reverse {
			slices.Reverse(numbers)
		}

		prediction := GetPrediction(numbers)
		predictions = append(predictions, prediction)
	}

	return predictions
}

func GetPrediction(numbers []int) int {
	differences := []int{}
	allZeroes := true

	for i := 0; i < len(numbers)-1; i++ {
		difference := numbers[i+1] - numbers[i]
		differences = append(differences, difference)

		if difference != 0 {
			allZeroes = false
		}
	}

	lastDifference := numbers[len(numbers)-1]

	if allZeroes {
		return lastDifference
	}

	return lastDifference + GetPrediction(differences)
}
