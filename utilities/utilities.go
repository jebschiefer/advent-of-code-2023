package utilities

import "strings"

func GetLines(input string) []string {
	lines := strings.Split(input, "\n")

	for i, line := range lines {
		lines[i] = strings.TrimSpace(line)
	}

	return lines
}
