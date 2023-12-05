package utilities

import (
	"os"
	"strings"
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
