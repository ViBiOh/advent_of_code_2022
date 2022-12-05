package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var moveRegex = regexp.MustCompile(`(?m)move ([0-9]+) from ([0-9]+) to ([0-9]+)`)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	var stacks [][]rune

	for scanner.Scan() {
		line := scanner.Text()

		if strings.Contains(line, "[") {
			stacks = loadStacks(stacks, line)
		}

		if strings.HasPrefix(line, "move") {
			count, from, to := parseMove(line)

			stacks = moveImproved(stacks, from-1, to-1, count)
		}
	}

	for _, stack := range stacks {
		fmt.Print(string(stack[len(stack)-1]))
	}
}

func loadStacks(stacks [][]rune, line string) [][]rune {
	if len(stacks) == 0 {
		stacks = make([][]rune, (len(line)+1)/4)
	}

	var stackIndex int

	for i := 1; i < len(line); i += 4 {
		if line[i] != ' ' {
			stacks[stackIndex] = append([]rune{rune(line[i])}, stacks[stackIndex]...)
		}

		stackIndex++
	}

	return stacks
}

func parseMove(line string) (int, int, int) {
	matches := moveRegex.FindAllStringSubmatch(line, -1)

	count, _ := strconv.ParseInt(matches[0][1], 10, 64)
	from, _ := strconv.ParseInt(matches[0][2], 10, 64)
	to, _ := strconv.ParseInt(matches[0][3], 10, 64)

	return int(count), int(from), int(to)
}

func moveOneByOne(stacks [][]rune, from, to, count int) [][]rune {
	fromLen := len(stacks[from])

	for i := 0; i < count; i++ {
		stacks[to] = append(stacks[to], stacks[from][fromLen-1-i])
	}

	stacks[from] = stacks[from][:fromLen-count]

	return stacks
}

func moveImproved(stacks [][]rune, from, to, count int) [][]rune {
	toRemove := len(stacks[from]) - count

	stacks[to] = append(stacks[to], stacks[from][toRemove:]...)
	stacks[from] = stacks[from][:toRemove]

	return stacks
}
