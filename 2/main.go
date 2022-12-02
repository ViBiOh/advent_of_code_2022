package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var firstRound [][]int = [][]int{
	{3, 0, 6},
	{6, 3, 0},
	{0, 6, 3},
}

var secondRound [][]int = [][]int{
	{3, 1, 2},
	{1, 2, 3},
	{2, 3, 1},
}

func main() {
	file, err := os.OpenFile("./input.txt", os.O_RDONLY, 0o600)
	if err != nil {
		panic(fmt.Errorf("open: %w", err))
	}
	defer func() {
		if err := file.Close(); err != nil {
			panic(fmt.Errorf("close: %w", err))
		}
	}()

	scanner := bufio.NewScanner(file)

	var score int
	var expectation int

	for scanner.Scan() {
		line := scanner.Text()

		parts := strings.Split(line, " ")
		if len(parts) != 2 {
			panic(fmt.Errorf("invalid input `%s`", line))
		}

		meIndex := int(parts[1][0]) - int('X')
		opponentIndex := int(parts[0][0]) - int('A')

		score += 1 + meIndex + firstRound[meIndex][opponentIndex]
		expectation += secondRound[meIndex][opponentIndex] + meIndex*3
	}

	fmt.Println("score", score)
	fmt.Println("expectation", expectation)
}
