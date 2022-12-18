package main

import (
	"bufio"
	"fmt"
	"os"
)

const (
	leftGap   = 1 << 4
	bottomGap = 3
)

var (
	cycle = uint64(2022)
	rocks = [][]int8{
		{1<<3 | 1<<2 | 1<<1 | 1<<0},
		{1 << 1, 1<<2 | 1<<1 | 1<<0, 1 << 1},
		{1 << 0, 1 << 0, 1<<2 | 1<<1 | 1<<0},
		{1, 1, 1, 1},
		{1<<1 | 1<<0, 1<<1 | 1<<0},
	}
)
var rocksLen = len(rocks)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	jet := scanner.Text()
	jetLen := len(jet)
	var jetIndex int

	var currentRock []int8
	var currentIndex int

	var lines []int8

	var x, y int

	for i := uint64(0); i < cycle; i++ {
		currentRock = rocks[currentIndex]

		x = 0

	xPosition:
		for _, line := range currentRock {
			if leftGap&(line<<x) != 0 {
				goto lines
			}
		}
		x += 1
		goto xPosition

	lines:
		y = highestRock(lines) + bottomGap + 1

		for i := len(lines) - len(currentRock); i < y; i++ {
			lines = append(lines, 0)
		}

	jet:
		switch jet[jetIndex] {
		case '<':
			if canMove(lines, currentRock, x+1, y) {
				x += 1
			}
		case '>':
			if canMove(lines, currentRock, x-1, y) {
				x -= 1
			}
		}

		jetIndex += 1
		jetIndex %= jetLen

		if canMove(lines, currentRock, x, y-1) {
			y -= 1

			goto jet
		}

		convertRock(lines, currentRock, x, y)
		// if pattern := hasPattern(lines); pattern != -1 {
		// 	fmt.Printf("pattern detected, patternSize=%d, i=%d, high=%d\n", pattern, i, highestRock(lines))
		// 	break
		// }

		currentIndex++
		currentIndex %= rocksLen
	}

	fmt.Println("total", highestRock(lines)+1)
}

func hasPattern(lines []int8) int {
	firstLine := lines[0]

	for i := 1; i < len(lines); i++ {
		if firstLine != lines[i] || i*2 > len(lines) {
			continue
		}

		var diff bool

		for j := 1; j < i; j++ {
			if lines[j] != lines[j+i] {
				continue
			}

			diff = true
		}

		if !diff {
			return i - 1
		}
	}

	return -1
}

func canMove(lines []int8, rock []int8, x, y int) bool {
	if x == -1 {
		return false
	}

	offset := y + len(rock) - 1

	for i := 0; i < len(rock); i++ {
		moved := rock[i] << x
		if moved < 0 || offset-i < 0 {
			return false
		}

		if lines[offset-i]&moved != 0 {
			return false
		}
	}

	return true
}

func convertRock(lines []int8, rock []int8, x, y int) {
	offset := y + len(rock) - 1

	for i := len(rock) - 1; i >= 0; i-- {
		lines[offset-i] |= rock[i] << x
	}
}

func render(lines []int8) {
	fmt.Println("\nRendering")

	for i := len(lines) - 1; i >= 0; i-- {
		for x := 6; x >= 0; x-- {
			if lines[i]&(1<<x) != 0 {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}

		fmt.Print("\n")
	}
}

func highestRock(lines []int8) int {
	for i := len(lines) - 1; i >= 0; i-- {
		if lines[i] != 0 {
			return i
		}
	}

	return -1
}
