package day16

import (
	"fmt"
	"strings"
)

type Grid [][]string
type Direction int

const (
	NONE Direction = iota
	UP
	DOWN
	LEFT
	RIGHT
)

func (grid Grid) stringify() string {
	content := ""

	for _, row := range grid {
		for _, cell := range row {
			content += cell
		}

		content += "\n"
	}

	return content
}

func (grid Grid) initEmpty() Grid {
	empty := Grid{}

	for _, row := range grid {
		emptyRow := strings.Split(strings.Repeat(".", len(row)), "")
		empty = append(empty, emptyRow)
	}

	return empty
}

func move(grid Grid, visitedGrid *Grid, visited map[string]int, x int, y int, direction Direction) {
	if isLoop(visited, x, y, direction) {
		return
	}

	nextX, nextY := x, y
	xMin, xMax := 0, len(grid[0])-1
	yMin, yMax := 0, len(grid)-1

	if x >= 0 {
		key := getKey(x, y, direction)
		visited[key] = 1
		(*visitedGrid)[y][x] = "#"
	}

	switch direction {
	case UP:
		nextY = y - 1

		if nextY >= yMin {
			nextTile := grid[nextY][nextX]

			if nextTile == "/" {
				move(grid, visitedGrid, visited, nextX, nextY, RIGHT)
			} else if nextTile == "\\" {
				move(grid, visitedGrid, visited, nextX, nextY, LEFT)
			} else if nextTile == "-" {
				move(grid, visitedGrid, visited, nextX, nextY, LEFT)
				move(grid, visitedGrid, visited, nextX, nextY, RIGHT)
			} else {
				move(grid, visitedGrid, visited, nextX, nextY, UP)
			}
		}
	case DOWN:
		nextY = y + 1

		if nextY <= yMax {
			nextTile := grid[nextY][nextX]

			if nextTile == "/" {
				move(grid, visitedGrid, visited, nextX, nextY, LEFT)
			} else if nextTile == "\\" {
				move(grid, visitedGrid, visited, nextX, nextY, RIGHT)
			} else if nextTile == "-" {
				move(grid, visitedGrid, visited, nextX, nextY, LEFT)
				move(grid, visitedGrid, visited, nextX, nextY, RIGHT)
			} else {
				move(grid, visitedGrid, visited, nextX, nextY, DOWN)
			}
		}
	case LEFT:
		nextX = x - 1

		if nextX >= xMin {
			nextTile := grid[nextY][nextX]

			if nextTile == "/" {
				move(grid, visitedGrid, visited, nextX, nextY, DOWN)
			} else if nextTile == "\\" {
				move(grid, visitedGrid, visited, nextX, nextY, UP)
			} else if nextTile == "|" {
				move(grid, visitedGrid, visited, nextX, nextY, UP)
				move(grid, visitedGrid, visited, nextX, nextY, DOWN)
			} else {
				move(grid, visitedGrid, visited, nextX, nextY, LEFT)
			}
		}
	case RIGHT:
		nextX = x + 1

		if nextX <= xMax {
			nextTile := grid[nextY][nextX]

			if nextTile == "/" {
				move(grid, visitedGrid, visited, nextX, nextY, UP)
			} else if nextTile == "\\" {
				move(grid, visitedGrid, visited, nextX, nextY, DOWN)
			} else if nextTile == "|" {
				move(grid, visitedGrid, visited, nextX, nextY, UP)
				move(grid, visitedGrid, visited, nextX, nextY, DOWN)
			} else {
				move(grid, visitedGrid, visited, nextX, nextY, RIGHT)
			}
		}
	}
}

func getKey(x int, y int, direction Direction) string {
	return fmt.Sprintf("%d-%d-%d", x, y, direction)
}

func isLoop(visited map[string]int, nextX int, nextY int, direction Direction) bool {
	key := getKey(nextX, nextY, direction)
	_, ok := visited[key]

	return ok
}

func countEnergized(grid Grid) int {
	count := 0

	for _, row := range grid {
		for _, tile := range row {
			if tile == "#" {
				count++
			}
		}
	}

	return count
}

func getMostEnergized(grid Grid) int {
	max := 0

	for i := 0; i < 4; i++ {
		if i > 0 {
			grid = rotate(grid)
		}

		for y := range grid {
			visited := grid.initEmpty()
			move(grid, &visited, map[string]int{}, -1, y, RIGHT)
			energized := countEnergized(visited)

			if energized > max {
				max = energized
			}
		}
	}

	return max
}

func rotate(grid Grid) Grid {
	rotated := Grid{}
	yLength := len(grid)

	for x := range grid[0] {
		column := []string{}

		for y := yLength - 1; y >= 0; y-- {
			symbol := grid[y][x]

			if symbol == "-" {
				symbol = "|"
			} else if symbol == "|" {
				symbol = "-"
			} else if symbol == "/" {
				symbol = "\\"
			} else if symbol == "\\" {
				symbol = "/"
			}

			column = append(column, symbol)
		}

		rotated = append(rotated, column)
	}

	return rotated
}
