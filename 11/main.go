package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

const (
	monkeyPrefix    = "Monkey"
	itemsPrefix     = "Starting items: "
	operationPrefix = "Operation: "
	divisiblePrefix = "Test: divisible by "
	truePrefix      = "If true: throw to monkey "
	falsePrefix     = "If false: throw to monkey "

	roundCount = 10000
	worryLevel = 1
)

type operation func(uint64) uint64

func simpleAdd(value uint64) operation {
	return func(old uint64) uint64 {
		return old + value
	}
}

func selfAdd(old uint64) uint64 {
	return old + old
}

func simpleMultiply(value uint64) operation {
	return func(old uint64) uint64 {
		return old * value
	}
}

func selfMultiply(old uint64) uint64 {
	return old * old
}

func parseOperation(line string) operation {
	parts := strings.Split(strings.TrimPrefix(line, "new = old "), " ")
	switch parts[0] {
	case "+":
		if parts[1] == "old" {
			return selfAdd
		}
		increment, _ := strconv.ParseUint(parts[1], 10, 64)
		return simpleAdd(increment)
	case "*":
		if parts[1] == "old" {
			return selfMultiply
		}
		increment, _ := strconv.ParseUint(parts[1], 10, 64)
		return simpleMultiply(increment)
	default:
		panic("unknown operation " + parts[0])
	}
}

type monkey struct {
	items                  []uint64
	operation              operation
	divisibleBy            uint64
	monkeyWhenDivisible    uint64
	monkeyWhenNotDivisible uint64

	totalInspected uint64
}

func (m *monkey) addItem(item uint64) {
	m.items = append(m.items, item)
}

func (m *monkey) makeTurn(others []*monkey, prime uint64) {
	for _, item := range m.items {
		item = m.operation(item) / worryLevel
		m.totalInspected++

		item = item % prime

		if item%m.divisibleBy == 0 {
			others[m.monkeyWhenDivisible].addItem(item)
		} else {
			others[m.monkeyWhenNotDivisible].addItem(item)
		}
	}

	m.items = m.items[:0]
}

func main() {
	monkeys := parseMonkeys()
	prime := computePrime(monkeys)

	for i := 0; i < roundCount; i++ {
		makeRound(monkeys, prime)
	}

	sort.Sort(sort.Reverse(MonkeyByInspected(monkeys)))

	fmt.Println("total", monkeys[0].totalInspected*monkeys[1].totalInspected)
}

func makeRound(monkeys []*monkey, prime uint64) {
	for i := 0; i < len(monkeys); i++ {
		monkeys[i].makeTurn(monkeys, prime)
	}
}

func computePrime(monkeys []*monkey) uint64 {
	var prime uint64 = 1

	for _, monkey := range monkeys {
		prime *= monkey.divisibleBy
	}

	return prime
}

func parseMonkeys() []*monkey {
	scanner := bufio.NewScanner(os.Stdin)

	var monkeys []*monkey

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		switch {
		case len(line) == 0:
			continue

		case strings.HasPrefix(line, monkeyPrefix):
			monkeys = append(monkeys, &monkey{})

		case strings.HasPrefix(line, itemsPrefix):
			for _, part := range strings.Split(strings.TrimPrefix(line, itemsPrefix), ", ") {
				value, _ := strconv.ParseUint(part, 10, 64)
				monkeys[len(monkeys)-1].addItem(value)
			}

		case strings.HasPrefix(line, operationPrefix):
			monkeys[len(monkeys)-1].operation = parseOperation(strings.TrimPrefix(line, operationPrefix))

		case strings.HasPrefix(line, divisiblePrefix):
			monkeys[len(monkeys)-1].divisibleBy, _ = strconv.ParseUint(strings.TrimPrefix(line, divisiblePrefix), 10, 64)

		case strings.HasPrefix(line, truePrefix):
			monkeys[len(monkeys)-1].monkeyWhenDivisible, _ = strconv.ParseUint(strings.TrimPrefix(line, truePrefix), 10, 64)

		case strings.HasPrefix(line, falsePrefix):
			monkeys[len(monkeys)-1].monkeyWhenNotDivisible, _ = strconv.ParseUint(strings.TrimPrefix(line, falsePrefix), 10, 64)
		default:
			panic("unhandled line: " + line)
		}
	}

	return monkeys
}

type MonkeyByInspected []*monkey

func (a MonkeyByInspected) Len() int      { return len(a) }
func (a MonkeyByInspected) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a MonkeyByInspected) Less(i, j int) bool {
	return a[i].totalInspected < a[j].totalInspected
}
