package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"sort"
)

func main() {
	var index, sum uint64
	var left, right []any

	scanner := bufio.NewScanner(os.Stdin)

	firstDivider := []any{[]any{float64(2)}}
	secondDivider := []any{[]any{float64(6)}}

	packets := []any{
		firstDivider,
		secondDivider,
	}

	for scanner.Scan() {
		line := scanner.Text()

		if len(line) == 0 {
			continue
		}

		if left == nil {
			_ = json.Unmarshal([]byte(line), &left)
		} else {
			_ = json.Unmarshal([]byte(line), &right)

			index++

			if compare(left, right) <= 0 {
				sum += index
			}

			packets = storePacket(packets, left)
			packets = storePacket(packets, right)

			right = nil
			left = nil
		}
	}

	fmt.Println("sum", sum)

	firstIndex := sort.Search(len(packets), func(i int) bool {
		return compare(packets[i], firstDivider) > 0
	})
	secondIndex := sort.Search(len(packets), func(i int) bool {
		return compare(packets[i], secondDivider) > 0
	})

	fmt.Println("divider", firstIndex*secondIndex)
}

func storePacket(packets []any, item any) []any {
	index := sort.Search(len(packets), func(i int) bool {
		return compare(packets[i], item) > 0
	})

	packets = append(packets, item)
	copy(packets[index+1:], packets[index:])
	packets[index] = item

	return packets
}

func compare(left, right any) int {
	switch leftValue := left.(type) {

	case float64:
		switch rightValue := right.(type) {

		case float64:
			return int(leftValue - rightValue)

		case []any:
			return compare([]any{leftValue}, rightValue)
		}

	case []any:
		switch rightValue := right.(type) {

		case float64:
			return compare(leftValue, []any{rightValue})

		case []any:
			rightValueLen := len(rightValue)

			for index, leftItem := range leftValue {
				if index == rightValueLen {
					return 1
				}

				diff := compare(leftItem, rightValue[index])
				if diff == 0 {
					continue
				}

				return diff
			}

			return len(leftValue) - rightValueLen
		}
	}

	panic("should not happen")
}
