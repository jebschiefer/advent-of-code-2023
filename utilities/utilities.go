package utilities

import "strings"

func GetLines(input string) []string {
	lines := strings.Split(input, "\n")

	for i, line := range lines {
		lines[i] = strings.TrimSpace(line)
	}

	return lines
}

func Sum(values []int) int {
	sum := 0

	for _, value := range values {
		sum += value
	}

	return sum
}
