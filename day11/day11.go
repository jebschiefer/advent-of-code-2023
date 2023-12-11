package day11

import (
	"aoc2023/utilities"
	"slices"
	"strings"
)

type Universe [][]string

type Galaxy struct {
	x int
	y int
}

func (universe Universe) stringify() string {
	universeString := ""

	for _, row := range universe {
		universeString += strings.Join(row, "") + "\n"
	}

	return universeString
}

func expandUniverse(universe Universe) Universe {
	universe = expandVertically(universe)
	universe = expandHorizontally(universe)
	return universe
}

func expandVertically(universe Universe) Universe {
	expanded := Universe{}

	emptyRow := strings.Repeat(".", len(universe[0]))

	for _, row := range universe {
		rowString := strings.Join(row, "")

		if rowString == emptyRow {
			expanded = append(expanded, row)
		}

		expanded = append(expanded, row)
	}

	return expanded
}

func expandHorizontally(universe Universe) Universe {
	expanded := Universe{}

	for _, row := range universe {
		rowCopy := make([]string, len(row))
		copy(rowCopy, row)
		expanded = append(expanded, rowCopy)
	}

	firstRow := universe[0]
	yLength := len(universe)

	emptyColumn := strings.Repeat(".", yLength)

	offset := 0

	for x := range firstRow {
		columnString := ""

		for y := 0; y < yLength; y++ {
			columnString += universe[y][x]
		}

		if columnString == emptyColumn {
			// Insert new column into each row at current index.
			for y := range expanded {
				expanded[y] = slices.Insert(expanded[y], x+offset, ".")
			}

			// Increment offset because a new column was added.
			offset++
		}
	}

	return expanded
}

func distanceBetweenGalaxies(a Galaxy, b Galaxy) int {
	return utilities.AbsoluteValue(b.x-a.x) + utilities.AbsoluteValue(b.y-a.y)
}

func getGalaxies(universe Universe) []Galaxy {
	galaxies := []Galaxy{}

	for y, row := range universe {
		for x := range row {
			if universe[y][x] == "#" {
				galaxies = append(galaxies, Galaxy{
					x: x,
					y: y,
				})
			}
		}
	}

	return galaxies
}

func getShortestPathsBetweenGalaxies(galaxies []Galaxy) []int {
	shortPathLengths := []int{}

	for i, galaxy := range galaxies {
		for j := i + 1; j < len(galaxies); j++ {
			shortestPathLength := distanceBetweenGalaxies(galaxy, galaxies[j])
			shortPathLengths = append(shortPathLengths, shortestPathLength)
		}
	}

	return shortPathLengths
}

func sumOfShortestDistancesBetweenGalaxies(universe Universe) int {
	universe = expandUniverse(universe)
	galaxies := getGalaxies(universe)
	shortestDistances := getShortestPathsBetweenGalaxies(galaxies)

	return utilities.Sum(shortestDistances)
}
