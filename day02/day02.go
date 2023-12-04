package day02

import (
	"regexp"
	"strconv"
)

func Sum(ids []int) int {
	sum := 0

	for _, value := range ids {
		sum += value
	}

	return sum
}

func GetPossibleGameIds(lines []string, bag map[string]int) []int {
	possibleGames := []int{}

	for _, line := range lines {
		gameId := GetGameId(line)
		mostSeen := GetMostSeen(line)

		isPossible := true

		for key, seenValue := range mostSeen {
			bagValue, exists := bag[key]

			if !exists || seenValue > bagValue {
				isPossible = false
			}
		}

		if isPossible {
			possibleGames = append(possibleGames, gameId)
		}
	}

	return possibleGames
}

func GetPowers(lines []string) []int {
	powers := []int{}

	for _, line := range lines {
		power := GetPower(line)
		powers = append(powers, power)
	}

	return powers
}

func GetMostSeen(line string) map[string]int {
	re := regexp.MustCompile(`Game \d+: `)
	line = re.ReplaceAllString(line, "")

	seen := map[string]int{
		"red":   0,
		"blue":  0,
		"green": 0,
	}

	for key := range seen {
		re := regexp.MustCompile(`(\d+) (` + key + `)`)
		seen = CheckLine(line, re, seen)
	}

	return seen
}

func CheckLine(line string, re *regexp.Regexp, seen map[string]int) map[string]int {
	matches := re.FindAllStringSubmatch(line, -1)

	for _, match := range matches {
		color := match[2]
		value, err := strconv.Atoi(match[1])

		if err != nil {
			panic(err)
		}

		bagValue, exists := seen[color]

		if exists && value > bagValue {
			seen[color] = value
		}
	}

	return seen
}

func GetGameId(line string) int {
	re := regexp.MustCompile(`Game (\d+): `)
	match := re.FindStringSubmatch(line)

	id, err := strconv.Atoi(match[1])

	if err != nil {
		panic(err)
	}

	return id
}

func GetPower(line string) int {
	mostSeen := GetMostSeen(line)

	power := 1

	for _, value := range mostSeen {
		power *= value
	}

	return power
}
