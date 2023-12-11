package day11

import (
	"aoc2023/utilities"
	"testing"
)

func TestExpandUniverse(t *testing.T) {
	input := utilities.ReadFileToGrid("./samples/example.txt")

	got := expandUniverse(input)
	want := Universe(utilities.ReadFileToGrid("./samples/example-expanded.txt"))

	utilities.CompareStrings(t, "\n"+got.stringify(), "\n"+want.stringify())
}

func TestExample1(t *testing.T) {
	input := utilities.ReadFileToGrid("./samples/example.txt")

	got := sumOfShortestDistancesBetweenGalaxies(input)
	want := 374

	utilities.CompareInts(t, got, want)
}

func TestPart1(t *testing.T) {
	input := utilities.ReadFileToGrid("./samples/input.txt")

	got := sumOfShortestDistancesBetweenGalaxies(input)
	want := 9591768

	utilities.CompareInts(t, got, want)
}
