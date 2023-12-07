package day06

import (
	"aoc2023/utilities"
	"regexp"
	"strings"
)

type Record struct {
	time     int
	distance int
}

func CalculateDistance(holdTime int, raceLength int) int {
	timeRemaining := raceLength - holdTime
	acceleration := holdTime

	return timeRemaining * acceleration
}

func WinningHoldTimes(raceLength int, record int) []int {
	winningTimes := []int{}

	for i := 0; i < raceLength; i++ {
		distance := CalculateDistance(i, raceLength)

		if distance > record {
			winningTimes = append(winningTimes, i)
		}
	}

	return winningTimes
}

func NumberOfWaysToWin(winningHoldtimes []int) int {
	return len(winningHoldtimes)
}

func ParseRecords(lines []string, digitSeparator string) []Record {
	records := []Record{}

	timeLine := strings.Replace(lines[0], "Time:", "", -1)
	distanceLine := strings.Replace(lines[1], "Distance:", "", -1)

	numberPattern := regexp.MustCompile(`\d+`)

	timeParts := numberPattern.FindAllString(timeLine, -1)
	distanceParts := numberPattern.FindAllString(distanceLine, -1)

	times := utilities.StringToInts(strings.Join(timeParts, digitSeparator))
	distances := utilities.StringToInts(strings.Join(distanceParts, digitSeparator))

	for i, time := range times {
		record := Record{
			time:     time,
			distance: distances[i],
		}

		records = append(records, record)
	}

	return records
}

func TotalWaysToWin(records []Record) int {
	waysToWin := []int{}

	for _, record := range records {
		winningHoldTimes := WinningHoldTimes(record.time, record.distance)
		waysToWin = append(waysToWin, NumberOfWaysToWin(winningHoldTimes))
	}

	total := 1

	for _, number := range waysToWin {
		total *= number
	}

	return total
}
