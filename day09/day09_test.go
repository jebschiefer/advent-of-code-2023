package day09

import (
	"aoc2023/utilities"
	"testing"
)

func TestGetPredictions(t *testing.T) {
	lines := utilities.ReadFileToLines("./samples/example.txt")

	got := GetPredictions(lines)
	want := []int{18, 28, 68}

	utilities.DeepCompare(t, got, want)
}

func TestExample1(t *testing.T) {
	lines := utilities.ReadFileToLines("./samples/input.txt")
	predictions := GetPredictions(lines)

	got := utilities.Sum(predictions)
	want := 114

	utilities.CompareInts(t, got, want)
}
