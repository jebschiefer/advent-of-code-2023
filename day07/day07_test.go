package day07

import (
	"aoc2023/utilities"
	"testing"
)

func TestGetHandType(t *testing.T) {
	got := GetHandType("AAAAA")
	want := FiveOfAKind
	utilities.CompareInts(t, int(got), int(want))

	got = GetHandType("AA8AA")
	want = FourOfAKind
	utilities.CompareInts(t, int(got), int(want))

	got = GetHandType("23332")
	want = FullHouse
	utilities.CompareInts(t, int(got), int(want))

	got = GetHandType("TTT98")
	want = ThreeOfAKind
	utilities.CompareInts(t, int(got), int(want))

	got = GetHandType("23432")
	want = TwoPair
	utilities.CompareInts(t, int(got), int(want))

	got = GetHandType("A23A4")
	want = OnePair
	utilities.CompareInts(t, int(got), int(want))

	got = GetHandType("23456")
	want = HighCard
	utilities.CompareInts(t, int(got), int(want))
}

func TestGetCardValue(t *testing.T) {
	got := GetCardValue('A')
	want := 14
	utilities.CompareInts(t, got, want)

	got = GetCardValue('K')
	want = 13
	utilities.CompareInts(t, got, want)

	got = GetCardValue('Q')
	want = 12
	utilities.CompareInts(t, got, want)

	got = GetCardValue('J')
	want = 11
	utilities.CompareInts(t, got, want)

	got = GetCardValue('T')
	want = 10
	utilities.CompareInts(t, got, want)

	got = GetCardValue('9')
	want = 9
	utilities.CompareInts(t, got, want)

	got = GetCardValue('2')
	want = 2
	utilities.CompareInts(t, got, want)
}

func TestExamplePart1(t *testing.T) {
	lines := utilities.ReadFileToLines("./samples/example.txt")

	got := CalculateTotalWinnings(lines)
	want := 6440

	utilities.CompareInts(t, got, want)
}

func TestPart1(t *testing.T) {
	lines := utilities.ReadFileToLines("./samples/input.txt")

	got := CalculateTotalWinnings(lines)
	want := 248569531

	utilities.CompareInts(t, got, want)
}
