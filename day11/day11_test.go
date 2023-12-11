package day11

import (
	"aoc2023/utilities"
	"testing"
)

func TestExample1(t *testing.T) {
	input := utilities.ReadFileToGrid("./samples/example.txt")

	got := sumOfShortestDistancesBetweenGalaxies(input, 2)
	want := 374

	utilities.CompareInts(t, got, want)
}

func TestPart1(t *testing.T) {
	input := utilities.ReadFileToGrid("./samples/input.txt")

	got := sumOfShortestDistancesBetweenGalaxies(input, 2)
	want := 9591768

	utilities.CompareInts(t, got, want)
}

func TestExample2Scale10(t *testing.T) {
	input := utilities.ReadFileToGrid("./samples/example.txt")

	got := sumOfShortestDistancesBetweenGalaxies(input, 10)
	want := 1030

	utilities.CompareInts(t, got, want)
}

func TestExample2Scale100(t *testing.T) {
	input := utilities.ReadFileToGrid("./samples/example.txt")

	got := sumOfShortestDistancesBetweenGalaxies(input, 100)
	want := 8410

	utilities.CompareInts(t, got, want)
}

func TestPart2(t *testing.T) {
	input := utilities.ReadFileToGrid("./samples/input.txt")

	got := sumOfShortestDistancesBetweenGalaxies(input, 1000000)
	want := 746962097860

	utilities.CompareInts(t, got, want)
}
