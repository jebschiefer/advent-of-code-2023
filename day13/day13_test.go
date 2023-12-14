package day13

import (
	"aoc2023/utilities"
	"testing"
)

func TestSequential(t *testing.T) {
	numbers := []int{1, 2, 3, 4}
	utilities.CompareBools(t, isSequentialList(numbers), true)

	numbers = []int{5, 6, 7, 8}
	utilities.CompareBools(t, isSequentialList(numbers), true)

	numbers = []int{5, 7, 10, 12}
	utilities.CompareBools(t, isSequentialList(numbers), false)

	// sublist 10, 11
	numbers = []int{5, 7, 10, 11}
	utilities.CompareBools(t, isSequentialList(numbers), true)

	// sublist 5, 6
	numbers = []int{5, 6, 10, 15}
	utilities.CompareBools(t, isSequentialList(numbers), true)

	// sublist 6, 7
	numbers = []int{4, 6, 7, 15}
	utilities.CompareBools(t, isSequentialList(numbers), true)
}

func TestSummarizeExample(t *testing.T) {
	input := utilities.ReadFile("./samples/example.txt")

	got := summarize(input)
	want := 405

	utilities.CompareInts(t, got, want)
}

func TestSummarizeExample2(t *testing.T) {
	input := utilities.ReadFile("./samples/example2.txt")

	got := summarize(input)
	want := 709

	utilities.CompareInts(t, got, want)
}

func TestPart1(t *testing.T) {
	input := utilities.ReadFile("./samples/input.txt")

	got := summarize(input)
	want := 405

	utilities.CompareInts(t, got, want)
}
