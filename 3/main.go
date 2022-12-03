package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var (
	upperStart = int('A')
	lowerStart = int('a')
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	var first, second int

	previousLines := make([]string, 0, 3)

	for scanner.Scan() {
		line := scanner.Text()

		first += getItemPriority(line)

		if previousLines = append(previousLines, line); len(previousLines) == 3 {
			second += getBadgesPriority(previousLines)
			previousLines = previousLines[:0]
		}
	}

	fmt.Println("first", first)
	fmt.Println("second", second)
}

func getItemPriority(line string) int {
	middle := len(line) / 2

	for _, item := range line[:middle] {
		if strings.ContainsRune(line[middle:], item) {
			return getPriority(item)
		}
	}

	return 0
}

func getBadgesPriority(lines []string) int {
	for _, item := range lines[0] {
		if strings.ContainsRune(lines[1], item) && strings.ContainsRune(lines[2], item) {
			return getPriority(item)
		}
	}

	return 0
}

func getPriority(item rune) int {
	element := int(item)

	if element < lowerStart {
		return 27 + element - upperStart
	}

	return 1 + element - lowerStart
}
