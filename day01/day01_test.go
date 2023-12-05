package day01

import (
	"aoc2023/utilities"
	"testing"
)

func TestGetCalibrationNumberEnds(t *testing.T) {
	line := "1abc2"
	got := GetCalibrationNumber(line)
	want := 12

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func TestGetCalibrationNumberInside(t *testing.T) {
	line := "pqr3stu8vwx"
	got := GetCalibrationNumber(line)
	want := 38

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func TestGetCalibrationNumberMultiple(t *testing.T) {
	line := "a1b2c3d4e5f"
	got := GetCalibrationNumber(line)
	want := 15

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func TestGetCalibrationNumberSingle(t *testing.T) {
	line := "treb7uchet"
	got := GetCalibrationNumber(line)
	want := 77

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func TestConvertsWordsToInts(t *testing.T) {
	line := "two1nine"
	got := ConvertWordsToInts(line)
	want := "219"

	if got != want {
		t.Errorf("got %s, wanted %s", got, want)
	}
}

func TestConvertsWordsToIntsOverlapping(t *testing.T) {
	line := "eightwothree"
	got := ConvertWordsToInts(line)
	want := "823"

	if got != want {
		t.Errorf("got %s, wanted %s", got, want)
	}
}

func TestGetCalibrationNumberConvertWordsEnds(t *testing.T) {
	line := "two1nine"
	line = ConvertWordsToInts(line)

	got := GetCalibrationNumber(line)
	want := 29

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func TestGetCalibrationNumberConvertWordsOverlapping(t *testing.T) {
	line := "eightwothree"
	line = ConvertWordsToInts(line)

	got := GetCalibrationNumber(line)
	want := 83

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func TestGetCalibrationNumberConvertWordsContained(t *testing.T) {
	line := "abcone2threexyz"
	line = ConvertWordsToInts(line)

	got := GetCalibrationNumber(line)
	want := 13

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func TestGetCalibrationNumberConvertWordsContainedOverlapping(t *testing.T) {
	line := "xtwone3four"
	line = ConvertWordsToInts(line)

	got := GetCalibrationNumber(line)
	want := 24

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func TestGetCalibrationNumberConvertWordsEndsNumbers(t *testing.T) {
	line := "4nineeightseven2"
	line = ConvertWordsToInts(line)

	got := GetCalibrationNumber(line)
	want := 42

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func TestGetCalibrationNumberConvertWordsWordDigit(t *testing.T) {
	line := "zoneight234"
	line = ConvertWordsToInts(line)

	got := GetCalibrationNumber(line)
	want := 14

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func TestGetCalibrationNumberConvertWordsDigitWord(t *testing.T) {
	line := "7pqrstsixteen"
	line = ConvertWordsToInts(line)

	got := GetCalibrationNumber(line)
	want := 76

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func TestSumCalibrationValuesPart1Example(t *testing.T) {
	lines := utilities.ReadFileToLines("./samples/example.txt")

	got := SumCalibrationValues(lines, false)
	want := 142

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func TestSumCalibrationValuesPart1(t *testing.T) {
	lines := utilities.ReadFileToLines("./samples/input.txt")

	got := SumCalibrationValues(lines, false)
	want := 55834

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func TestSumCalibrationValuesPart2Example(t *testing.T) {
	lines := utilities.ReadFileToLines("./samples/example2.txt")

	got := SumCalibrationValues(lines, true)
	want := 281

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func TestSumCalibrationValuesPart2(t *testing.T) {
	lines := utilities.ReadFileToLines("./samples/input.txt")

	got := SumCalibrationValues(lines, true)
	want := 53221

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}
