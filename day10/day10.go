package day10

import "fmt"

type Direction struct {
	value             string
	nextPossibleTiles []string
}

var UP = Direction{
	value:             "UP",
	nextPossibleTiles: []string{"|", "F", "7"},
}

var DOWN = Direction{
	value:             "DOWN",
	nextPossibleTiles: []string{"|", "L", "J"},
}

var RIGHT = Direction{
	value:             "RIGHT",
	nextPossibleTiles: []string{"-", "7", "J"},
}

var LEFT = Direction{
	value:             "LEFT",
	nextPossibleTiles: []string{"-", "L", "F"},
}

func CountSteps(startSymbol string, grid [][]string) int {
	steps := 0

	x, y := FindStart(startSymbol, grid)
	visted := map[string]string{}

	yLength := len(grid)
	xLength := len(grid[0])

	end := false

	for !end {
		currentTile := grid[y][x]
		key := fmt.Sprintf("%d,%d", x, y)
		visted[key] = currentTile

		if y > 0 && isAllowed(UP, x, y-1, grid, visted, currentTile, startSymbol) {
			// Move up
			y--
		} else if x < xLength-1 && isAllowed(RIGHT, x+1, y, grid, visted, currentTile, startSymbol) {
			// Move right
			x++
		} else if y < yLength-1 && isAllowed(DOWN, x, y+1, grid, visted, currentTile, startSymbol) {
			// Move down
			y++
		} else if x > 0 && isAllowed(LEFT, x-1, y, grid, visted, currentTile, startSymbol) {
			// Move left
			x--
		}

		steps++
		end = grid[y][x] == startSymbol
	}

	return steps
}

func FindStart(startSymbol string, grid [][]string) (x int, y int) {
	for y, row := range grid {
		for x, symbol := range row {
			if symbol == startSymbol {
				return x, y
			}
		}
	}

	return 0, 0
}

func isAllowed(direction Direction, x int, y int, grid [][]string, visited map[string]string, currentSymbol string, startSymbol string) bool {
	key := fmt.Sprintf("%d,%d", x, y)

	if value, ok := visited[key]; ok && value != startSymbol {
		// Already visited
		return false
	}

	if !IsPossibleDirection(currentSymbol, direction) {
		return false
	}

	if grid[y][x] == startSymbol {
		// Reached the end
		return true
	}

	allowed := false

	for _, possible := range direction.nextPossibleTiles {
		if grid[y][x] == possible {
			allowed = true
		}
	}

	return allowed
}

func IsPossibleDirection(symbol string, direction Direction) bool {
	moves := map[string][]string{
		"F": {"DOWN", "RIGHT"},
		"-": {"LEFT", "RIGHT"},
		"7": {"LEFT", "DOWN"},
		"|": {"UP", "DOWN"},
		"J": {"UP", "LEFT"},
		"L": {"UP", "RIGHT"},
		"S": {"UP", "DOWN", "LEFT", "RIGHT"},
	}

	possibleMoves, hasPossibleMove := moves[symbol]

	if !hasPossibleMove {
		return false
	}

	couldMove := false

	for _, moveOption := range possibleMoves {
		if moveOption == direction.value {
			couldMove = true
		}
	}

	return couldMove
}
