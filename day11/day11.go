package day11

import (
	"aoc2023/utilities"
	"strings"
)

type Universe [][]string

type Galaxy struct {
	x int
	y int
}

func expandUniverse(universe Universe, scaleFactor int) []Galaxy {
	scaledGalaxies := getGalaxies(universe)

	scaledGalaxies = expandVertically(universe, scaledGalaxies, scaleFactor)
	scaledGalaxies = expandHorizontally(universe, scaledGalaxies, scaleFactor)

	return scaledGalaxies
}

func expandVertically(universe Universe, scaledGalaxies []Galaxy, scaleFactor int) []Galaxy {
	galaxies := getGalaxies(universe)
	emptyRow := strings.Repeat(".", len(universe[0]))

	for y, row := range universe {
		rowString := strings.Join(row, "")

		if rowString == emptyRow {
			for i, galaxy := range galaxies {
				if galaxy.y > y {
					scaledGalaxies[i].y += scaleFactor - 1
				}
			}
		}
	}

	return scaledGalaxies
}

func expandHorizontally(universe Universe, scaledGalaxies []Galaxy, scaleFactor int) []Galaxy {
	galaxies := getGalaxies(universe)
	firstRow := universe[0]
	yLength := len(universe)

	emptyColumn := strings.Repeat(".", yLength)

	for x := range firstRow {
		columnString := ""

		for y := 0; y < yLength; y++ {
			columnString += universe[y][x]
		}

		if columnString == emptyColumn {
			for i, galaxy := range galaxies {
				if galaxy.x > x {
					scaledGalaxies[i].x += scaleFactor - 1
				}
			}
		}
	}

	return scaledGalaxies
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

func sumOfShortestDistancesBetweenGalaxies(universe Universe, scaleFactor int) int {
	scaledGalaxies := expandUniverse(universe, scaleFactor)
	shortestDistances := getShortestPathsBetweenGalaxies(scaledGalaxies)

	return utilities.Sum(shortestDistances)
}
