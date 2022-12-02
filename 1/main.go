package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

const topCount = 3

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	var topRank []int64
	var current int64

	for scanner.Scan() {
		line := scanner.Text()

		if len(line) == 0 {
			topRank = appendRank(topRank, current)
			current = 0
		} else {
			value, _ := strconv.ParseInt(line, 10, 64)
			current += value
		}
	}

	current = 0
	for _, top := range topRank {
		current += top
	}

	fmt.Println("top", topRank[0])
	fmt.Println("sum", current)
}

func appendRank(rank []int64, item int64) []int64 {
	if len(rank) == topCount && item < rank[topCount-1] {
		return rank
	}

	index := sort.Search(len(rank), func(i int) bool {
		return rank[i] <= item
	})

	rank = append(rank, item)
	copy(rank[index+1:], rank[index:])
	rank[index] = item

	if len(rank) > topCount {
		rank = rank[:topCount]
	}

	return rank
}
