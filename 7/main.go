package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

const (
	atMost    = 100000
	totalSize = 70000000
	freeDisk  = 30000000
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	directories := make(map[string]int64)
	pwd := []string{""}

	for scanner.Scan() {
		line := scanner.Text()

		switch line[0:4] {
		case "$ cd":
			cd := strings.TrimPrefix(line, "$ cd ")
			if cd == ".." {
				pwd = pwd[:len(pwd)-1]
			} else if cd == "/" {
				pwd = pwd[:1]
			} else {
				pwd = append(pwd, cd)
			}

		case "$ ls":
		default:
			parts := strings.Split(line, " ")
			size, _ := strconv.ParseInt(parts[0], 10, 64)

			for i := 0; i < len(pwd); i++ {
				directories[strings.Join(pwd[:i+1], "/")] += size
			}
		}
	}

	sumAtMost(directories)
	directoryToDelete(directories)
}

func sumAtMost(directories map[string]int64) {
	var sumAtMost int64

	for _, size := range directories {
		if size <= atMost {
			sumAtMost += size
		}
	}

	fmt.Println(sumAtMost)
}

func directoryToDelete(directories map[string]int64) {
	remainingSpace := totalSize - directories[""]
	neededSpace := freeDisk - remainingSpace

	var nearestSize int64 = math.MaxInt64

	for _, size := range directories {
		if size > neededSpace && size < nearestSize {
			nearestSize = size
		}
	}

	fmt.Println(nearestSize)
}
