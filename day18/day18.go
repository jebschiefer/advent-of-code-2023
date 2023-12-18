package day18

import (
	"regexp"
	"strconv"
	"strings"
)

// plot
// surround with box of another shape
// remove every dot between inner and out perimeters

type Grid []Row
type Row []string

type Instruction struct {
	direction string
	steps     int
	color     string
}

func (grid Grid) stringify() string {
	content := ""

	for _, row := range grid {
		line := strings.Join(row, "")
		content += line + "\n"
	}

	return content
}

func initGrid(width int, height int) Grid {
	grid := Grid{}

	for y := 0; y < height; y++ {
		row := Row{}

		for x := 0; x < width; x++ {
			row = append(row, ".")
		}

		grid = append(grid, row)
	}

	return grid
}

func getInstructions(lines []string) []Instruction {
	instructions := []Instruction{}

	re := regexp.MustCompile(`([R,L,U,D]) (\d+) \((.+)\)`)

	for _, line := range lines {
		matches := re.FindStringSubmatch(line)
		steps, _ := strconv.Atoi(matches[2])

		instruction := Instruction{
			direction: matches[1],
			steps:     steps,
			color:     matches[3],
		}

		instructions = append(instructions, instruction)
	}

	return instructions
}

func createGrid(lines []string) Grid {
	grid := initGrid(500, 500)
	instructions := getInstructions(lines)

	x, y := 100, 200

	for _, instruction := range instructions {
		switch instruction.direction {
		case "L":
			stop := x - instruction.steps
			for ; x > stop; x-- {
				grid[y][x] = "#"
			}
		case "R":
			stop := x + instruction.steps
			for ; x < stop; x++ {
				grid[y][x] = "#"
			}
		case "U":
			stop := y - instruction.steps
			for ; y > stop; y-- {
				grid[y][x] = "#"
			}
		case "D":
			stop := y + instruction.steps
			for ; y < stop; y++ {
				grid[y][x] = "#"
			}
		}
	}

	return grid
}

func pointInPolygon(grid Grid, x int, y int) bool {
	rowAbove := Row{}

	if y == 0 {
		rowLength := len(grid[0])
		rowAbove = strings.Split(strings.Repeat(".", rowLength), "")
	} else {
		rowAbove = grid[y-1]
	}

	count := 0

	if grid[y][x] == "#" {
		// Point is a boundary
		return false
	}

	for i := x; i < len(grid[y]); i++ {
		if grid[y][i] == "#" && rowAbove[i] == "#" {
			count++
		}
	}

	// Odd numbers are in polygon
	return count%2 == 1
}

func area(grid Grid) int {
	area := 0

	for y, row := range grid {
		for x := range row {
			if grid[y][x] == "#" || pointInPolygon(grid, x, y) {
				area++
			}
		}
	}

	return area
}
