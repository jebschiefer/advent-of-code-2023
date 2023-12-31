package day03

import (
	"fmt"
	"regexp"
	"sort"
	"strconv"
)

func GetPartNumbers(lines []string) []int {
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

func GetGearRatios(lines []string) []int {
	ratioParts := map[string][]int{}

	for i, line := range lines {
		re := regexp.MustCompile(`\d+`)

		matches := re.FindAllStringIndex(line, -1)

		for _, match := range matches {
			start := match[0]
			end := match[1]

			number := line[start:end]

			// check left
			if start > 0 && string(line[start-1]) == "*" {
				key := createKey(start-1, i)
				ratioParts = appendRatioPart(ratioParts, key, number)
			}

			// check right
			if end < len(line)-1 && string(line[end]) == "*" {
				key := createKey(end, i)
				ratioParts = appendRatioPart(ratioParts, key, number)
			}

			leftStart := max(0, start-1)
			rightEnd := min(end, len(line)-1)

			// check above
			if i > 0 {
				for k := leftStart; k <= rightEnd; k++ {
					aboveChar := string(lines[i-1][k])

					if aboveChar == "*" {
						key := createKey(k, i-1)
						ratioParts = appendRatioPart(ratioParts, key, number)
					}
				}
			}

			// check below
			if i < len(lines)-1 {
				for k := leftStart; k <= rightEnd; k++ {
					belowChar := string(lines[i+1][k])

					if belowChar == "*" {
						key := createKey(k, i+1)
						ratioParts = appendRatioPart(ratioParts, key, number)
					}
				}
			}
		}
	}

	return createRatios(ratioParts)
}

func isSymbol(char string) bool {
	re := regexp.MustCompile(`[0-9.]`)
	return !re.MatchString(char)
}

func createKey(x int, y int) string {
	return fmt.Sprintf("%d,%d", x, y)
}

func appendRatioPart(ratioParts map[string][]int, key string, value string) map[string][]int {
	values := ratioParts[key]

	if values == nil {
		values = []int{}
	}

	number, err := strconv.Atoi(value)

	if err == nil {
		ratioParts[key] = append(values, number)
	}

	return ratioParts
}

func createRatios(ratioParts map[string][]int) []int {
	ratios := []int{}

	for _, values := range ratioParts {
		if len(values) == 2 {
			product := values[0] * values[1]
			ratios = append(ratios, product)
		}
	}

	sort.Ints(ratios)

	return ratios
}
