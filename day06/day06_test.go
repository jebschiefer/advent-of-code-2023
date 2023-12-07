package day06

import (
	"aoc2023/utilities"
	"testing"
)

func TestCalculateDistanceExamples(t *testing.T) {
	got := CalculateDistance(0, 7)
	want := 0
	utilities.CompareInts(t, got, want)

	got = CalculateDistance(1, 7)
	want = 6
	utilities.CompareInts(t, got, want)

	got = CalculateDistance(2, 7)
	want = 10
	utilities.CompareInts(t, got, want)

	got = CalculateDistance(3, 7)
	want = 12
	utilities.CompareInts(t, got, want)

	got = CalculateDistance(4, 7)
	want = 12
	utilities.CompareInts(t, got, want)

	got = CalculateDistance(5, 7)
	want = 10
	utilities.CompareInts(t, got, want)

	got = CalculateDistance(6, 7)
	want = 6
	utilities.CompareInts(t, got, want)

	got = CalculateDistance(7, 7)
	want = 0
	utilities.CompareInts(t, got, want)
}

func TestWinningHoldTimesExample(t *testing.T) {
	got := WinningHoldTimes(7, 9)
	want := []int{2, 3, 4, 5}
	utilities.DeepCompare(t, got, want)

	got = WinningHoldTimes(15, 40)
	want = []int{4, 5, 6, 7, 8, 9, 10, 11}
	utilities.DeepCompare(t, got, want)

	got = WinningHoldTimes(30, 200)
	want = []int{11, 12, 13, 14, 15, 16, 17, 18, 19}
	utilities.DeepCompare(t, got, want)
}

func TestNumberOfWaysToWin(t *testing.T) {
	winningHoldTimes := []int{2, 3, 4, 5}

	got := NumberOfWaysToWin(winningHoldTimes)
	want := 4

	utilities.CompareInts(t, got, want)
}

func TestPart1Example(t *testing.T) {
	lines := utilities.ReadFileToLines("./samples/example.txt")
	records := ParseRecords(lines)

	got := TotalWaysToWin(records)
	want := 288

	utilities.CompareInts(t, got, want)
}

func TestPart1(t *testing.T) {
	lines := utilities.ReadFileToLines("./samples/input.txt")
	records := ParseRecords(lines)

	got := TotalWaysToWin(records)
	want := 1731600

	utilities.CompareInts(t, got, want)
}
