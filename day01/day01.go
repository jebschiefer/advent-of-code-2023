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

func GetCalibrationValues(lines string, convertWords bool) []int {
	values := []int{}
	scanner := bufio.NewScanner(strings.NewReader(lines))

	for scanner.Scan() {
		original := scanner.Text()
		line := scanner.Text()

		if convertWords && original != "" {
			line = ConvertWordsToInts(line)
		}

		calibrationValue := GetCalibrationNumber(line)
		values = append(values, calibrationValue)
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return values
}

func SumCalibrationValues(lines string, convertWords bool) int {
	values := GetCalibrationValues(lines, convertWords)

	sum := 0

	for _, value := range values {
		sum += value
	}

	return sum
}

func ConvertWordsToInts(line string) string {
	converted := ""

	for i, character := range line {
		isInt := character >= '0' && character <= '9'
		convertedNumber := getNumberSubstring(line, i)

		if isInt {
			converted += string(character)
		} else if convertedNumber != "" {
			converted += convertedNumber
		}
	}

	return converted
}

func getNumberSubstring(line string, offset int) string {
	number := ""

	wordMapping := map[string]string{
		"one":   "1",
		"two":   "2",
		"three": "3",
		"four":  "4",
		"five":  "5",
		"six":   "6",
		"seven": "7",
		"eight": "8",
		"nine":  "9",
	}

	for key, value := range wordMapping {
		end := offset + len(key)

		if end <= len(line) {
			substring := line[offset:end]

			if substring == key {
				number = value
			}
		}
	}

	return number
}
