package day03

import (
	"aoc2023/utilities"
	"regexp"
	"strconv"
)

func GetPartNumbers(input string) []int {
	lines := utilities.GetLines(input)

	partNumbers := []int{}

	for i, line := range lines {
		re := regexp.MustCompile(`\d+`)

		matches := re.FindAllStringIndex(line, -1)

		for _, match := range matches {
			isPartNumber := false

			start := match[0]
			end := match[1]

			// check left
			if start > 0 && isSymbol(string(line[start-1])) {
				isPartNumber = true
			}

			// check right
			if end < len(line)-1 && isSymbol(string(line[end])) {
				isPartNumber = true
			}

			leftStart := max(0, start-1)
			rightEnd := min(end, len(line)-1)

			// check above
			if i > 0 {
				for k := leftStart; k <= rightEnd; k++ {
					aboveChar := string(lines[i-1][k])

					if isSymbol(aboveChar) {
						isPartNumber = true
					}
				}
			}

			// check below
			if i < len(lines)-1 {
				for k := leftStart; k <= rightEnd; k++ {
					belowChar := string(lines[i+1][k])

					if isSymbol(belowChar) {
						isPartNumber = true
					}
				}
			}

			if isPartNumber {
				number := line[start:end]
				partNumber, err := strconv.Atoi(number)

				if err == nil {
					partNumbers = append(partNumbers, partNumber)
				}
			}
		}
	}

	return partNumbers
}

func isSymbol(char string) bool {
	re := regexp.MustCompile(`[0-9.]`)
	return !re.MatchString(char)
}
