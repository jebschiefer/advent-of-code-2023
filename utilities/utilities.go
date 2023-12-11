package utilities

import (
	"os"
	"reflect"
	"strconv"
	"strings"
	"testing"
)

func GetLines(input string) []string {
	lines := strings.Split(input, "\n")

	for i, line := range lines {
		lines[i] = strings.TrimSpace(line)
	}

	return lines
}

func Sum(values []int) int {
	sum := 0

	for _, value := range values {
		sum += value
	}

	return sum
}

func ReadFile(path string) string {
	content, err := os.ReadFile(path)

	if err != nil {
		panic(err)
	}

	return string(content)
}

func ReadFileToLines(path string) []string {
	content := ReadFile(path)
	return GetLines(content)
}

func ReadFileToGrid(path string) [][]string {
	grid := [][]string{}
	lines := ReadFileToLines(path)

	for _, line := range lines {
		row := []string{}

		for _, char := range line {
			row = append(row, string(char))
		}

		grid = append(grid, row)
	}

	return grid
}

func StringToInts(value string) []int {
	numbers := []int{}
	values := strings.Split(value, " ")

	for _, num := range values {
		number, err := strconv.Atoi(num)

		if err == nil {
			numbers = append(numbers, number)
		}
	}

	return numbers
}

func CompareInts(t *testing.T, got int, want int) {
	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func CompareBools(t *testing.T, got bool, want bool) {
	if got != want {
		t.Errorf("got %t, want %t", got, want)
	}
}

func CompareStrings(t *testing.T, got string, want string) {
	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}

func DeepCompare(t *testing.T, got []int, want []int) {
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func AbsoluteValue(x int) int {
	if x < 0 {
		return 0 - x
	}

	return x
}
