package day02

import (
	"aoc2023/utilities"
	"reflect"
	"testing"
)

func TestSumOfPossibleGameIdsExample(t *testing.T) {
	lines := utilities.ReadFileToLines("./samples/example.txt")

	bag := map[string]int{
		"red":   12,
		"blue":  14,
		"green": 13,
	}

	ids := GetPossibleGameIds(lines, bag)

	got := Sum(ids)
	want := 8

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func TestGetPossibleGameIdsExample(t *testing.T) {
	lines := utilities.ReadFileToLines("./samples/example.txt")

	bag := map[string]int{
		"red":   12,
		"blue":  14,
		"green": 13,
	}

	got := GetPossibleGameIds(lines, bag)
	want := []int{1, 2, 5}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestGetMostSeen(t *testing.T) {
	line := "Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green"

	got := GetMostSeen(line)
	want := map[string]int{
		"red":   4,
		"blue":  6,
		"green": 2,
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestGetMostSeenExampleGame3(t *testing.T) {
	line := "Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red"

	got := GetMostSeen(line)
	want := map[string]int{
		"red":   20,
		"blue":  6,
		"green": 13,
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestGetGameId(t *testing.T) {
	got := GetGameId("Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green")
	want := 1

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func TestSumOfPossibleGameIdsPart1(t *testing.T) {
	lines := utilities.ReadFileToLines("./samples/input.txt")

	bag := map[string]int{
		"red":   12,
		"blue":  14,
		"green": 13,
	}

	ids := GetPossibleGameIds(lines, bag)

	got := Sum(ids)
	want := 2176

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func TestGetPowerExampleGame1(t *testing.T) {
	input := `Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green`

	got := GetPower(input)
	want := 48

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func TestGetPowersExample(t *testing.T) {
	lines := utilities.ReadFileToLines("./samples/example.txt")

	got := GetPowers(lines)
	want := []int{48, 12, 1560, 630, 36}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestGetSumOfPowers(t *testing.T) {
	lines := utilities.ReadFileToLines("./samples/example.txt")

	powers := GetPowers(lines)

	got := Sum(powers)
	want := 2286

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func TestSumOfPowersPart2(t *testing.T) {
	lines := utilities.ReadFileToLines("./samples/input.txt")

	powers := GetPowers(lines)

	got := Sum(powers)
	want := 63700

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}
