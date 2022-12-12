package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

type position struct {
	x int
	y int
	z int
}

func main() {
	maze, current, target, lowerPositions := loadMaze()

	fmt.Println("part1", targetStep(maze, current, target))

	min := math.MaxInt

	for _, lowerPosition := range lowerPositions {
		if result := targetStep(maze, lowerPosition, target); result != 0 && result < min {
			min = result
		}
	}

	fmt.Println("part2", min)
}

func loadMaze() (maze [][]int, current, target position, lowerPositions []position) {
	var currentLine int

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := scanner.Text()

		maze = append(maze, make([]int, len(line)))

		for y, value := range line {
			switch value {
			case 'S':
				current = position{
					x: currentLine,
					y: y,
					z: 'a',
				}
				maze[currentLine][y] = int('a')
			case 'E':
				target = position{
					x: currentLine,
					y: y,
					z: 'z',
				}
				maze[currentLine][y] = int('z')
			default:
				if value == 'a' {
					lowerPositions = append(lowerPositions, position{x: currentLine, y: y, z: 'a'})
				}
				maze[currentLine][y] = int(value)
			}
		}

		currentLine += 1
	}

	return
}

func targetStep(maze [][]int, start, target position) int {
	var current []position

	visited := map[position]bool{
		start: true,
	}
	next := []position{
		start,
	}

	var iteration int

	for len(next) != 0 {
		current = next
		next = make([]position, 0)

		for _, current := range current {
			if current.x == target.x && current.y == target.y {
				return iteration
			}

			if newX := current.x - 1; newX >= 0 && (maze[newX][current.y] <= current.z || maze[newX][current.y]-current.z <= 1) {
				newCurrent := position{x: newX, y: current.y, z: maze[newX][current.y]}
				if !visited[newCurrent] {
					visited[newCurrent] = true
					next = append(next, newCurrent)
				}
			}

			if newY := current.y + 1; newY < len(maze[current.x]) && (maze[current.x][newY] <= current.z || maze[current.x][newY]-current.z <= 1) {
				newCurrent := position{x: current.x, y: newY, z: maze[current.x][newY]}
				if !visited[newCurrent] {
					visited[newCurrent] = true
					next = append(next, newCurrent)
				}
			}

			if newX := current.x + 1; newX < len(maze) && (maze[newX][current.y] <= current.z || maze[newX][current.y]-current.z <= 1) {
				newCurrent := position{x: newX, y: current.y, z: maze[newX][current.y]}
				if !visited[newCurrent] {
					visited[newCurrent] = true
					next = append(next, newCurrent)
				}
			}

			if newY := current.y - 1; newY >= 0 && (maze[current.x][newY] <= current.z || maze[current.x][newY]-current.z <= 1) {
				newCurrent := position{x: current.x, y: newY, z: maze[current.x][newY]}
				if !visited[newCurrent] {
					visited[newCurrent] = true
					next = append(next, newCurrent)
				}
			}
		}

		iteration++
	}

	return 0
}
