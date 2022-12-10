package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const crtWidth = 40

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	var sum int64
	var nextComputeCycle int64 = 20

	fmt.Print("#")

	var x int64 = 1
	var cycle int64 = 1

	for scanner.Scan() {
		switch parts := strings.Split(scanner.Text(), " "); parts[0] {
		case "addx":
			cycle, sum, nextComputeCycle = incrementCycle(cycle, nextComputeCycle, sum, x)
			handleCrt(cycle, x)

			value, _ := strconv.ParseInt(parts[1], 10, 64)
			x += value

			cycle, sum, nextComputeCycle = incrementCycle(cycle, nextComputeCycle, sum, x)
			handleCrt(cycle, x)

		case "noop":
			cycle, sum, nextComputeCycle = incrementCycle(cycle, nextComputeCycle, sum, x)
			handleCrt(cycle, x)
		}
	}

	fmt.Println("sum", sum)
}

func incrementCycle(cycle, nextComputeCycle, sum, x int64) (int64, int64, int64) {
	cycle++

	if nextComputeCycle == cycle {
		sum = sum + cycle*x
		nextComputeCycle += 40
	}

	return cycle, sum, nextComputeCycle
}

func handleCrt(cycle, x int64) {
	if abs((cycle-1)%crtWidth-x) <= 1 {
		fmt.Print("#")
	} else {
		fmt.Print(".")
	}

	if cycle%crtWidth == 0 {
		fmt.Print("\n")
	}
}

func abs(a int64) int64 {
	if a < 0 {
		return -a
	}

	return a
}
