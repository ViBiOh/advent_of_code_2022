package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	sand = 1
	rock = 2

	initialX = 500
)

type position struct {
	x, y int
}

func parsePosition(input string) position {
	var output position

	parts := strings.Split(input, ",")

	output.x, _ = strconv.Atoi(parts[0])
	output.y, _ = strconv.Atoi(parts[1])

	return output
}

func main() {
	cave, lowerBound := parseCave()

	var units int
	for ; flowSand(cave, lowerBound); units++ {
	}

	fmt.Println("part1", units)

	lowerBound += 2

	addRocks(cave, []position{
		{x: 0, y: lowerBound},
		{x: initialX * 2, y: lowerBound},
	}, lowerBound)

	for ; flowSand(cave, lowerBound); units++ {
	}

	fmt.Println("part2", units+1)
}

func flowSand(cave map[position]int, lowerBound int) bool {
	sandPosition := position{x: initialX}

	for {
		sandPosition.y += 1

		if sandPosition.y > lowerBound {
			return false
		}

		if cave[sandPosition] == 0 {
			continue
		}

		sandPosition.x--
		if cave[sandPosition] == 0 {
			continue
		}

		sandPosition.x += 2
		if cave[sandPosition] == 0 {
			continue
		}

		sandPosition.x--
		sandPosition.y--

		if sandPosition.x == initialX && sandPosition.y == 0 {
			return false
		}

		cave[sandPosition] = sand

		return true
	}
}

func parseCave() (map[position]int, int) {
	var coordinates []position

	cave := make(map[position]int)
	scanner := bufio.NewScanner(os.Stdin)

	lowerBound := 0

	for scanner.Scan() {
		for _, rawPath := range strings.Split(scanner.Text(), " -> ") {
			coordinates = append(coordinates, parsePosition(rawPath))
		}

		lowerBound = addRocks(cave, coordinates, lowerBound)

		coordinates = coordinates[:0]
	}

	return cave, lowerBound
}

func addRocks(cave map[position]int, coordinates []position, lowerBound int) int {
	for i := 1; i < len(coordinates); i++ {
		if coordinates[i-1].x != coordinates[i].x {
			start := coordinates[i-1].x
			end := coordinates[i].x

			if start > end {
				end = start
				start = coordinates[i].x
			}

			if coordinates[i].y > lowerBound {
				lowerBound = coordinates[i].y
			}

			for ; start <= end; start++ {
				cave[position{x: start, y: coordinates[i].y}] = rock
			}
		} else if coordinates[i-1].y != coordinates[i].y {
			start := coordinates[i-1].y
			end := coordinates[i].y

			if start > end {
				end = start
				start = coordinates[i].y
			}

			for ; start <= end; start++ {
				cave[position{x: coordinates[i].x, y: start}] = rock
			}

			if start-1 > lowerBound {
				lowerBound = start - 1
			}
		}
	}

	return lowerBound
}
