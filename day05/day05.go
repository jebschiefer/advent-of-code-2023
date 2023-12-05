package day05

import (
	"aoc2023/utilities"
	"regexp"
	"strings"
)

func GetSeeds(input string) []int {
	seeds := []int{}

	re := regexp.MustCompile(`seeds:((\s\d+)+)`)
	matches := re.FindAllStringSubmatch(input, -1)

	if matches[0] != nil && len(matches[0][1]) > 0 {
		numbersMatch := strings.TrimSpace(matches[0][1])
		seeds = utilities.StringToInts(numbersMatch)
	}

	return seeds
}

type Mapping struct {
	source      int
	destination int
	length      int
}

func GetMaps(input string) [][]Mapping {
	maps := [][]Mapping{}

	re := regexp.MustCompile(`map:\n((\d+\s?)+)`)
	matches := re.FindAllStringSubmatch(input, -1)

	for _, match := range matches {
		mappings := []Mapping{}

		if len(match[1]) > 0 {
			numberLines := strings.Split(match[1], "\n")

			for _, numberLine := range numberLines {
				numbers := utilities.StringToInts(numberLine)

				if len(numbers) == 3 {
					destination := numbers[0]
					source := numbers[1]
					length := numbers[2]

					mapping := Mapping{
						source:      source,
						destination: destination,
						length:      length,
					}

					mappings = append(mappings, mapping)
				}
			}
		}

		maps = append(maps, mappings)
	}

	return maps
}

func GetMapValue(mappings []Mapping, source int) int {
	value := source

	for _, mapping := range mappings {
		// source is inside the mapping range
		if source >= mapping.source && source < mapping.source+mapping.length {
			difference := source - mapping.source
			value = mapping.destination + difference
		}
	}

	return value
}

func GetSeedLocation(maps [][]Mapping, seed int) int {
	destination := seed

	for _, mapping := range maps {
		destination = GetMapValue(mapping, destination)
	}

	return destination
}

func GetSeedLocations(input string) []int {
	locations := []int{}

	seeds := GetSeeds(input)
	maps := GetMaps(input)

	for _, seed := range seeds {
		location := GetSeedLocation(maps, seed)
		locations = append(locations, location)
	}

	return locations
}

func Min(values []int) int {
	if len(values) == 0 {
		return 0
	}

	min := values[0]

	for _, value := range values {
		if value < min {
			min = value
		}
	}

	return min
}
