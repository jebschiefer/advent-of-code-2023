package day09

import (
	"aoc2023/utilities"
	"testing"
)

func TestGetPredictions(t *testing.T) {
	lines := utilities.ReadFileToLines("./samples/example.txt")

	got := GetPredictions(lines, false)
	want := []int{18, 28, 68}

	utilities.DeepCompare(t, got, want)
}

func TestPart1(t *testing.T) {
	lines := utilities.ReadFileToLines("./samples/input.txt")
	predictions := GetPredictions(lines, false)

	got := utilities.Sum(predictions)
	want := 1708206096

	utilities.CompareInts(t, got, want)
}

func TestExamplePart2(t *testing.T) {
	lines := utilities.ReadFileToLines("./samples/example.txt")

	got := GetPredictions(lines, true)
	want := []int{-3, 0, 5}

	utilities.DeepCompare(t, got, want)
}

func TestPart2(t *testing.T) {
	lines := utilities.ReadFileToLines("./samples/input.txt")
	predictions := GetPredictions(lines, true)

	got := utilities.Sum(predictions)
	want := 1050

	utilities.CompareInts(t, got, want)
}
