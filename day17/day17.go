package day17

import (
	"fmt"
	"math"
	"strconv"
)

type Grid [][]int

type Point struct {
	x     int
	y     int
	value int
}

func (point Point) isLast(grid Grid) bool {
	xMax := len(grid[0]) - 1
	yMax := len(grid) - 1

	return point.x == xMax && point.y == yMax
}

func minDistance(grid Grid) int {
	distance := 0

	currentPoint := Point{
		x:     0,
		y:     0,
		value: 0,
	}

	grid[0][0] = 0

	consecutive := []string{}
	direction := ""

	traveled := []Point{
		currentPoint,
	}

	for !currentPoint.isLast(grid) {
		currentPoint, direction = move(grid, currentPoint.x, currentPoint.y, direction, consecutive)

		streak := len(consecutive)

		if streak == 0 || consecutive[streak-1] == direction {
			consecutive = append(consecutive, direction)
		} else if consecutive[streak-1] != direction {
			consecutive = []string{}
		}

		distance += currentPoint.value
		traveled = append(traveled, currentPoint)
		grid[currentPoint.y][currentPoint.x] = 0
	}

	return distance
}

func gridStringsToInts(gridStrings [][]string) Grid {
	grid := Grid{}

	for _, row := range gridStrings {
		newRow := []int{}

		for _, cell := range row {
			value, _ := strconv.Atoi(cell)
			newRow = append(newRow, value)
		}

		grid = append(grid, newRow)
	}

	return grid
}

func move(grid Grid, x int, y int, lastDirection string, consecutive []string) (Point, string) {
	nextMove := Point{value: math.MaxInt}
	direction := ""

	xMin, xMax := 0, len(grid[0])-1
	yMin, yMax := 0, len(grid)-1

	possibleMoves := map[string]Point{
		"up":    {x: x, y: y - 1},
		"down":  {x: x, y: y + 1},
		"right": {x: x + 1, y: y},
		"left":  {x: x - 1, y: y},
	}

	for possibleDirection, next := range possibleMoves {
		// Can't go same direction more than 3 times in a row
		if possibleDirection == lastDirection && len(consecutive) >= 3 {
			continue
		}

		if possibleDirection == "up" && (next.x == xMax || len(consecutive) < 3) {
			continue
		}

		if possibleDirection == "left" && next.x != xMax {
			continue
		}

		if next.x >= xMin && next.x <= xMax && next.y >= yMin && next.y <= yMax {
			value := grid[next.y][next.x]

			if value < nextMove.value && value > 0 {
				direction = possibleDirection

				nextMove = Point{
					x:     next.x,
					y:     next.y,
					value: value,
				}
			}
		}
	}

	return nextMove, direction
}

func getKey(x int, y int) string {
	return fmt.Sprintf("%d,%d", x, y)
}

func getMin(dist map[string]int) (string, int) {
	min := math.MaxInt
	key := ""

	for distKey := range dist {
		value := dist[distKey]
		if value < min {
			min = value
			key = distKey
		}
	}

	return key, min
}
