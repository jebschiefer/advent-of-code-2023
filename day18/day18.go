package day18

import (
	"math"
	"regexp"
	"strconv"
	"strings"
)

type Grid []Row
type Row []string

type Instruction struct {
	direction string
	steps     int
	color     string
}

type Point struct {
	x int
	y int
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

func getInstructionsHex(lines []string) []Instruction {
	instructions := []Instruction{}

	re := regexp.MustCompile(`\((.+)\)`)

	for _, line := range lines {
		matches := re.FindStringSubmatch(line)
		color := matches[1]
		direction := color[len(color)-1:]
		steps, _ := strconv.ParseInt(color[1:len(color)-1], 16, 64)

		switch direction {
		case "0":
			direction = "R"
		case "1":
			direction = "D"
		case "2":
			direction = "L"
		case "3":
			direction = "U"
		}

		instruction := Instruction{
			direction: direction,
			steps:     int(steps),
		}

		instructions = append(instructions, instruction)
	}

	return instructions
}

func createGrid(lines []string, instructions []Instruction) Grid {
	grid := initGrid(500, 500)

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

func getPoints(instructions []Instruction) []Point {
	points := []Point{}
	x, y := 0, 0

	for _, instruction := range instructions {
		switch instruction.direction {
		case "L":
			stop := x - instruction.steps
			for ; x > stop; x-- {
				points = append(points, Point{
					x: x,
					y: y,
				})
			}
		case "R":
			stop := x + instruction.steps
			for ; x < stop; x++ {
				points = append(points, Point{
					x: x,
					y: y,
				})
			}
		case "U":
			stop := y - instruction.steps
			for ; y > stop; y-- {
				points = append(points, Point{
					x: x,
					y: y,
				})
			}
		case "D":
			stop := y + instruction.steps
			for ; y < stop; y++ {
				points = append(points, Point{
					x: x,
					y: y,
				})
			}
		}
	}

	return points
}

func getShoelaceArea(points []Point) int {
	n := len(points) - 1
	area := 0

	for i := 0; i < n; i++ {
		x1 := points[i].x
		y1 := points[i].y

		x2 := points[0].x
		y2 := points[0].y

		if i != n {
			x2 = points[i+1].x
			y2 = points[i+1].y
		}

		area += ((x1 * y2) - (x2 * y1))
	}

	return int(math.Abs(float64(area / 2)))
}

func getFillArea(points []Point) int {
	a := getShoelaceArea(points)
	b := len(points)

	// Pick's theorem, need i
	// a = i + (b / 2) - 1
	// a - (b / 2) = i - 1
	// a - (b / 2) + 1 = i
	i := a - (b / 2) + 1

	// Internal + boundary points
	return i + b
}
