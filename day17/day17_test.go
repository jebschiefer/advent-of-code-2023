package day17

import (
	"aoc2023/utilities"
	"testing"
)

func TestExamplePart1(t *testing.T) {
	grid := utilities.ReadFileToGrid("./samples/example.txt")

	got := minDistance(gridStringsToInts(grid))
	want := 102

	utilities.CompareInts(t, got, want)
}
