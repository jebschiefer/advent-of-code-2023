package day14

const ROUND = "O"
const SQUARE = "#"
const EMPTY = "."

func countLoad(grid [][]string) int {
	grid = tilt(grid)

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

func tilt(grid [][]string) [][]string {
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
