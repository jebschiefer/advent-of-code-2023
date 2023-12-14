package day14

import (
	"aoc2023/utilities"
	"testing"
)

func TestExamplePart1(t *testing.T) {
	input := utilities.ReadFileToGrid("./samples/example.txt")

	got := countLoad(input)
	want := 136

	utilities.CompareInts(t, got, want)
}

func TestPart1(t *testing.T) {
	input := utilities.ReadFileToGrid("./samples/input.txt")

	got := countLoad(input)
	want := 136

	utilities.CompareInts(t, got, want)
}
