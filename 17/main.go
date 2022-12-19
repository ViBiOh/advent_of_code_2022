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
	cycle        = uint64(1000000000000)
	patternCycle = uint64(50000)
	rocks        = [][]int8{
		{1<<3 | 1<<2 | 1<<1 | 1<<0},
		{1 << 1, 1<<2 | 1<<1 | 1<<0, 1 << 1},
		{1 << 0, 1 << 0, 1<<2 | 1<<1 | 1<<0},
		{1, 1, 1, 1},
		{1<<1 | 1<<0, 1<<1 | 1<<0},
	}
)
var rocksLen = len(rocks)

type occurence struct {
	i      uint64
	height uint64
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	jet := scanner.Text()

	var jetIndex int
	var rockIndex int

	var lines []int8

	var patternKey string
	pattern := make(map[string][]occurence)

	for i := uint64(0); i < patternCycle; i++ {
		lines, jetIndex = fallRock(lines, jet, jetIndex, rockIndex)

		rockIndex++
		rockIndex %= rocksLen

		highest := highestRock(lines)
		patternKey = fmt.Sprintf("%d_%d_%d", jetIndex, rockIndex, lines[highest])
		if value, ok := pattern[patternKey]; ok {
			pattern[patternKey] = append(value, occurence{i: i, height: uint64(highest)})
		} else {
			pattern[patternKey] = []occurence{{i: i, height: uint64(highest)}}
		}
	}

	initialHeight := uint64(highestRock(lines))

	var projectedHeight uint64
	var restart int

	usedPattern := pattern[patternKey]
	gap := usedPattern[len(usedPattern)-1].height - usedPattern[len(usedPattern)-2].height
	size := usedPattern[len(usedPattern)-1].i - usedPattern[len(usedPattern)-2].i

	restart = int((cycle - usedPattern[0].i) % size)
	projectedHeight = usedPattern[0].height + (cycle-usedPattern[0].i)/size*gap

	for i := 0; i < restart; i++ {
		lines, jetIndex = fallRock(lines, jet, jetIndex, rockIndex)

		rockIndex++
		rockIndex %= rocksLen
	}

	projectedHeight += uint64(highestRock(lines)) - initialHeight

	fmt.Println("total", projectedHeight)
}

func fallRock(lines []int8, jet string, jetIndex, rockIndex int) ([]int8, int) {
	currentRock := rocks[rockIndex]

	x := 0

xPosition:
	for _, line := range currentRock {
		if leftGap&(line<<x) != 0 {
			goto lines
		}
	}
	x += 1
	goto xPosition

lines:
	y := highestRock(lines) + bottomGap + 1

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
	jetIndex %= len(jet)

	if canMove(lines, currentRock, x, y-1) {
		y -= 1

		goto jet
	}

	convertRock(lines, currentRock, x, y)

	return lines, jetIndex
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

func highestRock(lines []int8) int {
	for i := len(lines) - 1; i >= 0; i-- {
		if lines[i] != 0 {
			return i
		}
	}

	return -1
}
