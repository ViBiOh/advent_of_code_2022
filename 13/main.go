package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	var index, sum uint64
	var left, right []any

	scanner := bufio.NewScanner(os.Stdin)

	firstDivider := []any{[]any{float64(2)}}
	secondDivider := []any{[]any{float64(6)}}

	firstIndex := 1
	secondIndex := 2

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

			firstIndex = updateIndex(firstIndex, left, firstDivider)
			firstIndex = updateIndex(firstIndex, right, firstDivider)

			secondIndex = updateIndex(secondIndex, left, secondDivider)
			secondIndex = updateIndex(secondIndex, right, secondDivider)

			right = nil
			left = nil
		}
	}

	fmt.Println("sum", sum)

	fmt.Println("divider", firstIndex*secondIndex)
}

func updateIndex(index int, toCompare, divider []any) int {
	if compare(toCompare, divider) <= 0 {
		return index + 1
	}

	return index
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
