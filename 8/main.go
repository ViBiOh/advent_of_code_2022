package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var hidden int8 = 15

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	var grid [][]int

	for scanner.Scan() {
		line := scanner.Text()

		if len(grid) == 0 {
			grid = make([][]int, len(line))
		}

		for i, tree := range line {
			treeValue, _ := strconv.ParseInt(string(tree), 10, 64)

			grid[i] = append(grid[i], int(treeValue))
		}
	}

	visible := len(grid)*2 + (len(grid[0])-2)*2
	var maxScore int

	for i := 1; i < len(grid)-1; i++ {
		for j := 1; j < len(grid[i])-1; j++ {
			if isVisible(i, j, grid) {
				visible += 1
			}
			if score := scenicScore(i, j, grid); score > maxScore {
				maxScore = score
			}
		}
	}

	fmt.Println("visible", visible)
	fmt.Println("score", maxScore)
}

func isVisible(x, y int, grid [][]int) bool {
	value := grid[x][y]

	var visibility int8

	for i := 0; i < x; i++ {
		if grid[i][y] >= value {
			visibility += 1 << 0
			break
		}
	}

	for i := x + 1; i < len(grid); i++ {
		if grid[i][y] >= value {
			visibility += 1 << 1
			break
		}
	}

	for i := 0; i < y; i++ {
		if grid[x][i] >= value {
			visibility += 1 << 2
			break
		}
	}

	for i := y + 1; i < len(grid[x]); i++ {
		if grid[x][i] >= value {
			visibility += 1 << 3
			break
		}
	}

	return visibility != hidden
}

func scenicScore(x, y int, grid [][]int) int {
	value := grid[x][y]

	var left, right, top, bottom int

	for i := x - 1; i >= 0; i-- {
		left++
		if grid[i][y] >= value {
			break
		}
	}

	for i := x + 1; i < len(grid); i++ {
		right++
		if grid[i][y] >= value {
			break
		}
	}

	for i := y - 1; i >= 0; i-- {
		top++
		if grid[x][i] >= value {
			break
		}
	}

	for i := y + 1; i < len(grid[x]); i++ {
		bottom++
		if grid[x][i] >= value {
			break
		}
	}

	return left * right * top * bottom
}
