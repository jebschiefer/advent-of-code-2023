package day18

import (
	"aoc2023/utilities"
	"testing"
)

func TestPointInPolygon(t *testing.T) {
	grid := Grid{
		{"#", ".", "#"},
		{"#", ".", "#"},
	}
	got := pointInPolygon(grid, 1, 1)
	want := true
	utilities.CompareBools(t, got, want)

	grid = Grid{
		{"#", ".", "#"},
		{"#", ".", "#"},
	}
	got = pointInPolygon(grid, 0, 1)
	want = false
	utilities.CompareBools(t, got, want)

	grid = Grid{
		{"#", ".", "#"},
		{"#", ".", "#"},
	}
	got = pointInPolygon(grid, 2, 1)
	want = false
	utilities.CompareBools(t, got, want)

	grid = Grid{
		{".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", "#", ".", ".", ".", ".", ".", ".", ".", "#"},
		{"#", "#", "#", "#", "#", ".", ".", ".", ".", ".", ".", "#", ".", ".", ".", ".", ".", ".", ".", "#"},
	}
	got = pointInPolygon(grid, 2, 1)
	want = false
	utilities.CompareBools(t, got, want)

	grid = Grid{
		{".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", "#", ".", ".", ".", ".", ".", ".", ".", "#"},
		{"#", "#", "#", "#", "#", ".", ".", ".", ".", ".", ".", "#", ".", ".", ".", ".", ".", ".", ".", "#"},
	}
	got = pointInPolygon(grid, 6, 1)
	want = false
	utilities.CompareBools(t, got, want)

	grid = Grid{
		{".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", "#", ".", ".", ".", ".", ".", ".", ".", "#"},
		{"#", "#", "#", "#", "#", ".", ".", ".", ".", ".", ".", "#", ".", ".", ".", ".", ".", ".", ".", "#"},
	}
	got = pointInPolygon(grid, 12, 1)
	want = true
	utilities.CompareBools(t, got, want)

	// #...########.......####
	grid = Grid{
		{"#", ".", ".", ".", "#", ".", ".", ".", ".", ".", ".", "#", ".", ".", ".", ".", ".", ".", ".", "#", ".", ".", "."},
		{"#", ".", ".", ".", "#", "#", "#", "#", "#", "#", "#", "#", ".", ".", ".", ".", ".", ".", ".", "#", "#", "#", "#"},
	}
	got = pointInPolygon(grid, 2, 1)
	want = true
	utilities.CompareBools(t, got, want)

	// ...#...#.
	// .###.###.
	grid = Grid{
		{".", ".", ".", "#", ".", ".", ".", "#", "."},
		{".", "#", "#", "#", ".", "#", "#", "#", "."},
	}
	got = pointInPolygon(grid, 0, 1)
	want = false
	utilities.CompareBools(t, got, want)
}

func TestExamplePart1(t *testing.T) {
	lines := utilities.ReadFileToLines("./samples/example.txt")
	instructions := getInstructions(lines)
	grid := createGrid(lines, instructions)

	got := area(grid)
	want := 62

	utilities.CompareInts(t, got, want)
}

func TestPart1(t *testing.T) {
	lines := utilities.ReadFileToLines("./samples/input.txt")
	instructions := getInstructions(lines)
	grid := createGrid(lines, instructions)

	got := area(grid)
	want := 35991

	utilities.CompareInts(t, got, want)
}

func TestExamplePart1Shoelace(t *testing.T) {
	lines := utilities.ReadFileToLines("./samples/example.txt")
	intstructions := getInstructions(lines)
	points := getPoints(intstructions)

	got := getFillArea(points)
	want := 62

	utilities.CompareInts(t, got, want)
}

func TestPart1Shoelace(t *testing.T) {
	lines := utilities.ReadFileToLines("./samples/input.txt")
	intstructions := getInstructions(lines)
	points := getPoints(intstructions)

	got := getFillArea(points)
	want := 35991

	utilities.CompareInts(t, got, want)
}

func TestExamplePart2(t *testing.T) {
	lines := utilities.ReadFileToLines("./samples/example.txt")
	intstructions := getInstructionsHex(lines)
	points := getPoints(intstructions)

	got := getFillArea(points)
	want := 952408144115

	utilities.CompareInts(t, got, want)
}

func TestPart2(t *testing.T) {
	lines := utilities.ReadFileToLines("./samples/input.txt")
	intstructions := getInstructionsHex(lines)
	points := getPoints(intstructions)

	got := getFillArea(points)
	want := 54058824661845

	utilities.CompareInts(t, got, want)
}
