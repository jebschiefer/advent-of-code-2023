package day14

import (
	"aoc2023/utilities"
	"testing"
)

func TestExamplePart1(t *testing.T) {
	input := utilities.ReadFileToGrid("./samples/example.txt")
	grid := tilt(input)

	got := countLoad(grid)
	want := 136

	utilities.CompareInts(t, got, want)
}

func TestExamplePart2(t *testing.T) {
	input := utilities.ReadFileToGrid("./samples/example.txt")

	minCycles := findMinCycles(input, 1000000000)

	grid := cycle(input, minCycles)

	got := countLoad(grid)
	want := 64

	utilities.CompareInts(t, got, want)
}

func TestPart1(t *testing.T) {
	input := utilities.ReadFileToGrid("./samples/input.txt")
	grid := tilt(input)

	got := countLoad(grid)
	want := 109661

	utilities.CompareInts(t, got, want)
}

func TestPart2(t *testing.T) {
	input := utilities.ReadFileToGrid("./samples/input.txt")

	minCycles := findMinCycles(input, 1000000000)

	grid := cycle(input, minCycles)

	got := countLoad(grid)
	want := 90176

	utilities.CompareInts(t, got, want)
}
