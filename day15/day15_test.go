package day15

import (
	"aoc2023/utilities"
	"testing"
)

func TestHash(t *testing.T) {
	got := hash("HASH")
	want := 52

	utilities.CompareInts(t, got, want)
}

func TestExamplePart1(t *testing.T) {
	input := utilities.ReadFile("./samples/example.txt")

	got := sumHash(input)
	want := 1320

	utilities.CompareInts(t, got, want)
}

func TestPart1(t *testing.T) {
	input := utilities.ReadFile("./samples/input.txt")

	got := sumHash(input)
	want := 517315

	utilities.CompareInts(t, got, want)
}
