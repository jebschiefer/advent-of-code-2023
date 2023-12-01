package day01

import (
	"bufio"
	"regexp"
	"strconv"
	"strings"
)

func GetCalibrationNumber(line string) int {
	re := regexp.MustCompile(`[^\d]`)
	numString := re.ReplaceAllString(line, "")

	if len(numString) > 0 {
		first := numString[0:1]
		last := numString[len(numString)-1:]
		numString := first + last

		num, err := strconv.Atoi(numString)

		if err != nil {
			panic(err)
		}

		return num
	}

	return 0
}

func SumCalibrationValues(lines string) int {
	scanner := bufio.NewScanner(strings.NewReader(lines))

	sum := 0

	for scanner.Scan() {
		calibrationValue := GetCalibrationNumber(scanner.Text())
		sum += calibrationValue
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return sum
}
