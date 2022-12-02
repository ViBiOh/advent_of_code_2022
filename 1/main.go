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

	var topRank []int64
	var current int64

	for scanner.Scan() {
		line := scanner.Text()

		if len(line) == 0 {
			topRank = appendRank(topRank, current)
			current = 0

			continue
		}

		value, err := strconv.ParseInt(line, 10, 64)
		if err != nil {
			panic(fmt.Errorf("parse: %w", err))
		}

		current += value
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
