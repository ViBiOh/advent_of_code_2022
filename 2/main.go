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
	scanner := bufio.NewScanner(os.Stdin)

	var score int
	var expectation int

	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), " ")

		meIndex := int(parts[1][0]) - int('X')
		opponentIndex := int(parts[0][0]) - int('A')

		score += 1 + meIndex + firstRound[meIndex][opponentIndex]
		expectation += secondRound[meIndex][opponentIndex] + meIndex*3
	}

	fmt.Println("score", score)
	fmt.Println("expectation", expectation)
}
