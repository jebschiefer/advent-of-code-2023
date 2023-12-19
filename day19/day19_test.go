package day19

import (
	"aoc2023/utilities"
	"testing"
)

func TestExamplePart1(t *testing.T) {
	input := utilities.ReadFile("./samples/example.txt")
	acceptedParts := getAccepted(input)

	got := sumParts(acceptedParts)
	want := 19114

	utilities.CompareInts(t, got, want)
}

func TestPart1(t *testing.T) {
	input := utilities.ReadFile("./samples/input.txt")
	acceptedParts := getAccepted(input)

	got := sumParts(acceptedParts)
	want := 374873

	utilities.CompareInts(t, got, want)
}
