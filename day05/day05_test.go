package day05

import (
	"aoc2023/utilities"
	"reflect"
	"testing"
)

func TestGetSeeds(t *testing.T) {
	input := utilities.ReadFile("./samples/example.txt")

	got := GetSeeds(input)
	want := []int{79, 14, 55, 13}

	utilities.DeepCompare(t, got, want)
}

func TestGetMapValue(t *testing.T) {
	mappings := []Mapping{
		{
			source:      98,
			destination: 50,
			length:      2,
		},
		{
			source:      50,
			destination: 52,
			length:      48,
		},
	}

	got := GetMapValue(mappings, 0)
	want := 0
	utilities.CompareInts(t, got, want)

	got = GetMapValue(mappings, 49)
	want = 49
	utilities.CompareInts(t, got, want)

	got = GetMapValue(mappings, 50)
	want = 52
	utilities.CompareInts(t, got, want)

	got = GetMapValue(mappings, 97)
	want = 99
	utilities.CompareInts(t, got, want)

	got = GetMapValue(mappings, 98)
	want = 50
	utilities.CompareInts(t, got, want)

	got = GetMapValue(mappings, 99)
	want = 51
	utilities.CompareInts(t, got, want)

	got = GetMapValue(mappings, 100)
	want = 100
	utilities.CompareInts(t, got, want)
}

func TestGetMaps(t *testing.T) {
	input := `seed-to-soil map:
50 98 2
52 50 48`

	got := GetMaps(input)
	want := [][]Mapping{
		{
			{
				source:      98,
				destination: 50,
				length:      2,
			},
			{
				source:      50,
				destination: 52,
				length:      48,
			},
		},
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestExpandSeedRange(t *testing.T) {
	input := "seeds: 10 5 30 3"

	seeds := GetSeeds(input)

	got := ExpandSeedRange(seeds)
	want := []int{10, 11, 12, 13, 14, 30, 31, 32}

	utilities.DeepCompare(t, got, want)
}

func TestExamplePart1(t *testing.T) {
	input := utilities.ReadFile("./samples/example.txt")

	locations := GetSeedLocations(input, false)

	got := Min(locations)
	want := 35

	utilities.CompareInts(t, got, want)
}

func TestPart1(t *testing.T) {
	input := utilities.ReadFile("./samples/input.txt")

	locations := GetSeedLocations(input, false)

	got := Min(locations)
	want := 278755257

	utilities.CompareInts(t, got, want)
}

func TestExamplePart2(t *testing.T) {
	input := utilities.ReadFile("./samples/example.txt")

	locations := GetSeedLocations(input, true)

	got := Min(locations)
	want := 46

	utilities.CompareInts(t, got, want)
}

// func TestPart2(t *testing.T) {
// 	input := utilities.ReadFile("./samples/input.txt")

// 	// Naive approach takes roughly 5 minutes
// 	locations := GetSeedLocations(input, true)

// 	got := Min(locations)
// 	want := 26829166

// 	utilities.CompareInts(t, got, want)
// }
