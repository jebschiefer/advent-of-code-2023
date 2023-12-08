package day08

import (
	"aoc2023/utilities"
	"testing"
)

func TestGetKey(t *testing.T) {
	got := GetKey("AAA = (BBB, BBB)")
	want := "AAA"

	utilities.CompareStrings(t, got, want)
}

func TestGetLeftRight(t *testing.T) {
	left, right := GetLeftRight("AAA = (BBB, CCC)")
	wantLeft := "BBB"
	wantRight := "CCC"

	utilities.CompareStrings(t, left, wantLeft)
	utilities.CompareStrings(t, right, wantRight)
}

func TestCountStepsExamplePart1(t *testing.T) {
	lines := utilities.ReadFileToLines("./samples/example.txt")

	got := CountSteps(lines)
	want := 6

	utilities.CompareInts(t, got, want)
}

func TestCountStepsPart1(t *testing.T) {
	lines := utilities.ReadFileToLines("./samples/input.txt")

	got := CountSteps(lines)
	want := 16409

	utilities.CompareInts(t, got, want)
}
