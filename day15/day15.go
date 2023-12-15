package day15

import "strings"

func hash(input string) int {
	currentValue := 0

	for _, char := range input {
		currentValue += int(char)
		currentValue *= 17
		currentValue %= 256
	}

	return currentValue
}

func sumHash(input string) int {
	input = strings.ReplaceAll(input, "\n", "")
	sequence := strings.Split(input, ",")

	sum := 0

	for _, part := range sequence {
		sum += hash(part)
	}

	return sum
}
