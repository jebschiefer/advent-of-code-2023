package day10

import (
	"aoc2023/utilities"
	"testing"
)

func TestCountSteps(t *testing.T) {
	input := utilities.ReadFileToGrid("./samples/example.txt")

	got := CountSteps("S", input)
	want := 8

	utilities.CompareInts(t, got, want)
}

func TestExamplePart1(t *testing.T) {
	input := utilities.ReadFileToGrid("./samples/example.txt")
	steps := CountSteps("S", input)

	got := steps / 2
	want := 4
	utilities.CompareInts(t, got, want)

	input = utilities.ReadFileToGrid("./samples/example2.txt")
	steps = CountSteps("S", input)

	got = steps / 2
	want = 8
	utilities.CompareInts(t, got, want)
}

func TestPart1(t *testing.T) {
	input := utilities.ReadFileToGrid("./samples/input.txt")
	steps := CountSteps("S", input)

	got := steps / 2
	want := 6682
	utilities.CompareInts(t, got, want)
}
