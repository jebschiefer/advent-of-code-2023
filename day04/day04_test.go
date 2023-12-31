package day04

import (
	"aoc2023/utilities"
	"reflect"
	"testing"
)

func TestGetPointsPerCardExample(t *testing.T) {
	lines := utilities.ReadFileToLines("./samples/example.txt")

	got := GetPointsPerCard(lines)
	want := []int{8, 2, 2, 1, 0, 0}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestExamplePart1(t *testing.T) {
	lines := utilities.ReadFileToLines("./samples/example.txt")

	points := GetPointsPerCard(lines)

	got := utilities.Sum(points)
	want := 13

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func TestPart1(t *testing.T) {
	lines := utilities.ReadFileToLines("./samples/input.txt")

	points := GetPointsPerCard(lines)

	got := utilities.Sum(points)
	want := 21558

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func TestExamplePart2(t *testing.T) {
	lines := utilities.ReadFileToLines("./samples/example.txt")

	got := CountInstancesOfCards(lines)
	want := 30

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func TestPart2(t *testing.T) {
	lines := utilities.ReadFileToLines("./samples/input.txt")

	got := CountInstancesOfCards(lines)
	want := 10425665

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}
