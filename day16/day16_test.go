package day16

import (
	"aoc2023/utilities"
	"testing"
)

func TestExamplePart1(t *testing.T) {
	grid := Grid(utilities.ReadFileToGrid("./samples/example.txt"))
	visited := grid.initEmpty()

	move(grid, &visited, map[string]int{}, -1, 0, RIGHT)

	got := countEnergized(visited)
	want := 46

	utilities.CompareInts(t, got, want)
}

func TestPart1(t *testing.T) {
	grid := Grid(utilities.ReadFileToGrid("./samples/input.txt"))
	visited := grid.initEmpty()

	move(grid, &visited, map[string]int{}, -1, 0, RIGHT)

	got := countEnergized(visited)
	want := 7242

	utilities.CompareInts(t, got, want)
}

func TestExamplePart2(t *testing.T) {
	grid := Grid(utilities.ReadFileToGrid("./samples/example.txt"))

	got := getMostEnergized(grid)
	want := 51

	utilities.CompareInts(t, got, want)
}

func TestPart2(t *testing.T) {
	grid := Grid(utilities.ReadFileToGrid("./samples/input.txt"))

	got := getMostEnergized(grid)
	want := 7572

	utilities.CompareInts(t, got, want)
}
