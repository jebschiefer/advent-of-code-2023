package day07

import (
	"aoc2023/utilities"
	"testing"
)

func TestGetHandType(t *testing.T) {
	got := GetHandType("AAAAA", false)
	want := FiveOfAKind
	utilities.CompareInts(t, int(got), int(want))

	got = GetHandType("AA8AA", false)
	want = FourOfAKind
	utilities.CompareInts(t, int(got), int(want))

	got = GetHandType("23332", false)
	want = FullHouse
	utilities.CompareInts(t, int(got), int(want))

	got = GetHandType("TTT98", false)
	want = ThreeOfAKind
	utilities.CompareInts(t, int(got), int(want))

	got = GetHandType("23432", false)
	want = TwoPair
	utilities.CompareInts(t, int(got), int(want))

	got = GetHandType("A23A4", false)
	want = OnePair
	utilities.CompareInts(t, int(got), int(want))

	got = GetHandType("23456", false)
	want = HighCard
	utilities.CompareInts(t, int(got), int(want))
}

func TestGetHandTypeJoker(t *testing.T) {
	got := GetHandType("32T3K", true)
	want := OnePair
	utilities.CompareInts(t, int(got), int(want))

	got = GetHandType("KK677", true)
	want = TwoPair
	utilities.CompareInts(t, int(got), int(want))

	got = GetHandType("T55J5", true)
	want = FourOfAKind
	utilities.CompareInts(t, int(got), int(want))

	got = GetHandType("KTJJT", true)
	want = FourOfAKind
	utilities.CompareInts(t, int(got), int(want))

	got = GetHandType("QQQJA", true)
	want = FourOfAKind
	utilities.CompareInts(t, int(got), int(want))

	got = GetHandType("2JJJ5", true)
	want = FourOfAKind
	utilities.CompareInts(t, int(got), int(want))
}

func TestGetCardValue(t *testing.T) {
	got := GetCardValue('A', false)
	want := 14
	utilities.CompareInts(t, got, want)

	got = GetCardValue('K', false)
	want = 13
	utilities.CompareInts(t, got, want)

	got = GetCardValue('Q', false)
	want = 12
	utilities.CompareInts(t, got, want)

	got = GetCardValue('J', false)
	want = 11
	utilities.CompareInts(t, got, want)

	got = GetCardValue('J', true)
	want = 1
	utilities.CompareInts(t, got, want)

	got = GetCardValue('T', false)
	want = 10
	utilities.CompareInts(t, got, want)

	got = GetCardValue('9', false)
	want = 9
	utilities.CompareInts(t, got, want)

	got = GetCardValue('2', false)
	want = 2
	utilities.CompareInts(t, got, want)
}

func TestExamplePart1(t *testing.T) {
	lines := utilities.ReadFileToLines("./samples/example.txt")

	got := CalculateTotalWinnings(lines, false)
	want := 6440

	utilities.CompareInts(t, got, want)
}

func TestPart1(t *testing.T) {
	lines := utilities.ReadFileToLines("./samples/input.txt")

	got := CalculateTotalWinnings(lines, false)
	want := 248569531

	utilities.CompareInts(t, got, want)
}

func TestExamplePart2(t *testing.T) {
	lines := utilities.ReadFileToLines("./samples/example.txt")

	got := CalculateTotalWinnings(lines, true)
	want := 5905

	utilities.CompareInts(t, got, want)
}

func TestPart2(t *testing.T) {
	lines := utilities.ReadFileToLines("./samples/input.txt")

	got := CalculateTotalWinnings(lines, true)
	want := 248569531

	utilities.CompareInts(t, got, want)
}
