package day14

const ROUND = "O"
const SQUARE = "#"
const EMPTY = "."

const CLOCKWISE = "clockwise"
const COUNTER_CLOCKWISE = "counter clockwise"

type Grid [][]string

func countLoad(grid Grid) int {
	load := 0
	rowLoadFactor := len(grid)

	for _, row := range grid {
		rocks := 0

		for _, cell := range row {
			if cell == ROUND {
				rocks++
			}
		}

		load += rocks * rowLoadFactor
		rowLoadFactor--
	}

	return load
}

func tilt(grid Grid) Grid {
	for y, row := range grid {
		if y > 0 {
			for x, cell := range row {
				if cell == ROUND {
					for above := y - 1; above >= 0; above-- {
						if grid[above][x] == SQUARE {
							break
						}

						if grid[above][x] == EMPTY {
							grid[above][x] = ROUND
							grid[above+1][x] = EMPTY
						}
					}
				}
			}
		}
	}

	return grid
}

func rotate(grid Grid, direction string) Grid {
	rotated := Grid{}
	yLength := len(grid)

	if direction == CLOCKWISE {
		for x := range grid[0] {
			column := []string{}

			for y := 0; y < yLength; y++ {
				column = append(column, grid[y][x])
			}

			rotated = append(rotated, column)
		}
	}

	if direction == COUNTER_CLOCKWISE {
		start := len(grid) - 1

		for x := start; x >= 0; x-- {
			column := []string{}

			for y := 0; y < yLength; y++ {
				column = append(column, grid[y][x])
			}

			rotated = append(rotated, column)
		}
	}

	return rotated
}

func cycle(grid Grid, times int) Grid {
	for i := 0; i < times; i++ {
		for j := 0; j < 4; j++ {
			if j > 0 {
				for k := 0; k < j; k++ {
					grid = rotate(grid, CLOCKWISE)
				}
			}

			grid = tilt(grid)

			if j > 0 {
				for k := 0; k < j; k++ {
					grid = rotate(grid, COUNTER_CLOCKWISE)
				}
			}
		}
	}

	return grid
}

func stringify(grid Grid) string {
	content := ""

	for _, row := range grid {
		for _, cell := range row {
			content += cell
		}
	}

	return content
}

func findMinCycles(grid Grid, total int) int {
	seen := map[string]int{}

	cycleCount := 0
	cycleLength := 0
	cycleStart := 0
	cycleFound := false

	for !cycleFound {
		key := stringify(grid)

		if seenAtCount, ok := seen[key]; ok {
			cycleStart = seenAtCount
			cycleLength = cycleCount - seenAtCount
			cycleFound = true
		} else {
			seen[key] = cycleCount
			cycleCount++
			grid = cycle(grid, 1)
		}
	}

	remainingCycles := (total - cycleCount) % cycleLength

	return cycleStart + remainingCycles
}
