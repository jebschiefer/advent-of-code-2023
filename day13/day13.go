package day13

import (
	"aoc2023/utilities"
	"regexp"
	"slices"
)

func summarize(input string) int {
	re := regexp.MustCompile(`\n\n`)
	groups := re.Split(input, -1)

	sum := 0

	for _, group := range groups {
		lines := utilities.GetLines(group)
		grid := utilities.LinesToGrid(lines)

		horizontalMirror := findHorizonalMirror(lines)
		verticalMirror := findVerticalMirror(grid)

		if horizontalMirror > 0 {
			horizontalMirror *= 100
			sum += horizontalMirror
		}

		if verticalMirror > 0 {
			sum += verticalMirror
		}
	}

	return sum
}

func findHorizonalMirror(lines []string) int {
	mirrorStart := -1
	set := map[int]struct{}{}

	for i, line := range lines {
		for j := i + 1; j < len(lines); j++ {
			if line == lines[j] {
				set[i] = struct{}{}
				set[j] = struct{}{}

				if j-i == 1 {
					mirrorStart = i
				}
			}
		}
	}

	matches := []int{}

	for match := range set {
		matches = append(matches, match)
	}

	slices.Sort(matches)
	totalMatches := len(matches)
	sequential := isSequentialList(matches)

	if mirrorStart >= 0 && totalMatches > 2 && sequential {
		fromStart := matches[0] == 0
		toEnd := matches[totalMatches-1] == len(lines)-1

		if fromStart || toEnd {
			return mirrorStart + 1
		}
	}

	return -1
}

// func findHorizonalMirror(lines []string) int {
// 	reverse := []string{}
// 	reverse = append(reverse, lines...)
// 	slices.Reverse(reverse)

// 	for i, line := range lines {
// 		for j := i + 1; j < len(lines); j++ {
// 			if line == reverse[j] {
// 			}
// 		}
// 	}

// 	return -1
// }

func findVerticalMirror(grid [][]string) int {
	columns := []string{}
	yLength := len(grid)

	for x := range grid[0] {
		column := ""

		for y := 0; y < yLength; y++ {
			column += grid[y][x]
		}

		columns = append(columns, column)
	}

	return findHorizonalMirror(columns)
}

func isSequentialList(numbers []int) bool {
	if len(numbers) < 2 {
		return false
	}

	for i := 0; i < len(numbers)-1; i++ {
		if numbers[i] != numbers[i+1]-1 {
			return isSequentialList(numbers[:i+1]) || isSequentialList(numbers[i+1:])
		}
	}

	return true
}
