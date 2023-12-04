package day03

import (
	"aoc2023/utilities"
	"reflect"
	"testing"
)

func TestGetPartNumbersExample(t *testing.T) {
	lines := utilities.ReadFile("./samples/example.txt")

	got := GetPartNumbers(lines)
	want := []int{467, 35, 633, 617, 592, 755, 664, 598}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestGetSumPartNumbersExample(t *testing.T) {
	lines := utilities.ReadFile("./samples/example.txt")

	partNumbers := GetPartNumbers(lines)

	got := utilities.Sum(partNumbers)
	want := 4361

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func TestGetSumPartNumbersExtraCases(t *testing.T) {
	lines := utilities.ReadFile("./samples/case-repeating-numbers.txt")

	partNumbers := GetPartNumbers(lines)

	got := utilities.Sum(partNumbers)
	want := 4

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func TestGetSumPartNumbersPart1(t *testing.T) {
	lines := utilities.ReadFile("./samples/input.txt")

	partNumbers := GetPartNumbers(lines)

	got := utilities.Sum(partNumbers)
	want := 514969

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func TestGetGearRatios(t *testing.T) {
	lines := utilities.ReadFile("./samples/example.txt")

	got := GetGearRatios(lines)
	want := []int{16345, 451490}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestGetGearRatiosSumExample(t *testing.T) {
	lines := utilities.ReadFile("./samples/example.txt")

	ratios := GetGearRatios(lines)

	got := utilities.Sum(ratios)
	want := 467835

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func TestGetGearRatiosSumPart2(t *testing.T) {
	lines := utilities.ReadFile("./samples/input.txt")

	ratios := GetGearRatios(lines)

	got := utilities.Sum(ratios)
	want := 78915902

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}
