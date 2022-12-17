package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strconv"
)

var (
	re          = regexp.MustCompile(`(?mi).*?x=(.*?), y=(.*?):.*?x=(.*?), y=(.*?)$`)
	yPart1      = int64(2000000)
	yPart2      = int64(4000000)
	xMultiplier = 4000000
)

type position struct {
	x, y int64
}

func (p position) manhattan(other position) int64 {
	return abs(p.x-other.x) + abs(p.y-other.y)
}

type Range struct {
	start, end int64
}

type Ranges []Range

func (r Ranges) Insert(value Range) []Range {
	index := sort.Search(len(r), func(i int) bool {
		return r[i].start > value.start
	})

	r = append(r, value)
	copy(r[index+1:], r[index:])
	r[index] = value

	return r
}

func (r Ranges) Sum() int64 {
	if len(r) == 0 {
		return 0
	}

	latest := r[0].start

	var sum int64

	for _, item := range r {
		if item.start > latest {
			sum += item.end - item.start + 1
			latest = item.end
		} else if item.end > latest {
			sum += item.end - latest
			latest = item.end
		}
	}

	return sum
}

func (r Ranges) Gap(lower, upper int64) int64 {
	if len(r) == 0 {
		return 0
	}

	latest := max(r[0].start, lower)

	for _, item := range r {
		if item.start > latest {
			return latest + 1
		}

		if end := min(item.end, upper); end > latest {
			latest = end
		}
	}

	return -1
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	lines := make([]Ranges, yPart2)

	for scanner.Scan() {
		sensor, beacon := scanInput(scanner)

		for i := int64(0); i < yPart2; i++ {
			lines[i] = computeRange(lines[i], sensor, beacon, i)
		}
	}

	fmt.Println("not a beacon at", yPart1, ":", lines[yPart1].Sum())

	for i := int64(0); i < yPart2; i++ {
		if gap := lines[i].Gap(0, yPart2); gap != -1 {
			fmt.Println("lost beacon is at x=", gap, "y=", i, ", frequency is", gap*int64(xMultiplier)+i)
			break
		}
	}
}

func computeRange(ranges Ranges, sensor, beacon position, targetY int64) []Range {
	distance := sensor.manhattan(beacon)

	if yDelta := abs(targetY - sensor.y); yDelta <= distance {
		xRange := distance - yDelta
		ranges = ranges.Insert(Range{start: sensor.x - xRange, end: sensor.x + xRange})
	}

	return ranges
}

func abs(value int64) int64 {
	if value < 0 {
		return -value
	}
	return value
}

func min(a, b int64) int64 {
	if a < b {
		return a
	}
	return b
}

func max(a, b int64) int64 {
	if a < b {
		return b
	}
	return a
}

func scanInput(scanner *bufio.Scanner) (position, position) {
	matches := re.FindStringSubmatch(scanner.Text())

	var sensor position
	sensor.x, _ = strconv.ParseInt(matches[1], 10, 64)
	sensor.y, _ = strconv.ParseInt(matches[2], 10, 64)

	var beacon position
	beacon.x, _ = strconv.ParseInt(matches[3], 10, 64)
	beacon.y, _ = strconv.ParseInt(matches[4], 10, 64)

	return sensor, beacon
}
