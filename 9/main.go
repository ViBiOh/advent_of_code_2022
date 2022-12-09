package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type position struct {
	x int
	y int
}

func (p position) move(direction string) position {
	switch direction {
	case "R":
		p.x++
	case "U":
		p.y++
	case "L":
		p.x--
	case "D":
		p.y--
	default:
		panic(fmt.Sprintf("unknown direction `%s`", direction))
	}

	return p
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	knotsCount := 10
	knots := make([]position, knotsCount)

	tailIndex := knotsCount - 1
	tailPositions := map[position]bool{
		knots[tailIndex]: true,
	}

	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " ")

		direction := line[0]
		count, _ := strconv.ParseInt(line[1], 10, 64)

		for i := 0; i < int(count); i++ {
			knots[0] = knots[0].move(direction)

			for j := 1; j < knotsCount; j++ {
				knots[j] = computeTail(knots[j-1], knots[j])
			}

			if !tailPositions[knots[tailIndex]] {
				tailPositions[knots[tailIndex]] = true
			}
		}
	}

	fmt.Println("positions", len(tailPositions))
}

func computeTail(head, tail position) position {
	if abs(head.x-tail.x) > 1 {
		if head.x > tail.x {
			tail.x = head.x - 1
		} else {
			tail.x = head.x + 1
		}

		tail.y = head.y
	}

	if abs(head.y-tail.y) > 1 {
		if head.y > tail.y {
			tail.y = head.y - 1
		} else {
			tail.y = head.y + 1
		}

		tail.x = head.x
	}

	return tail
}

func abs(value int) int {
	if value < 0 {
		return -value
	}

	return value
}
