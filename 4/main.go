package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type elve struct {
	start int64
	end   int64
}

func (e elve) fullOverlap(other elve) bool {
	return e.start <= other.start && e.end >= other.end
}

func (e elve) partialOverlap(other elve) bool {
	return e.start <= other.end && e.end >= other.start
}

func parseElve(raw string) elve {
	var output elve

	sections := strings.SplitN(raw, "-", 2)

	output.start, _ = strconv.ParseInt(sections[0], 10, 64)
	output.end, _ = strconv.ParseInt(sections[1], 10, 64)

	return output
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	var fullOverlap, partialOverlap int

	for scanner.Scan() {
		line := scanner.Text()

		elves := strings.SplitN(line, ",", 2)
		first := parseElve(elves[0])
		second := parseElve(elves[1])

		if first.fullOverlap(second) || second.fullOverlap(first) {
			fullOverlap++
		}

		if first.partialOverlap(second) {
			partialOverlap++
		}
	}

	fmt.Println("full", fullOverlap)
	fmt.Println("partial", partialOverlap)
}
