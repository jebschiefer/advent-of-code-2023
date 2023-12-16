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
