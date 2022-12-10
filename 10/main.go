package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const crtWidth = 40

type cpu struct {
	cycle     int64
	register  int64
	listeners []func(*cpu)
}

func newCpu(listeners []func(*cpu)) *cpu {
	return &cpu{
		register:  1,
		listeners: listeners,
	}
}

func (c *cpu) do(cycleCount int, value int64) {
	for i := 0; i < cycleCount; i++ {
		c.cycle += 1

		for _, listener := range c.listeners {
			listener(c)
		}
	}

	c.register += value
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	var sum int64
	var nextComputeCycle int64 = 20

	cpu := newCpu([]func(*cpu){
		func(current *cpu) {
			if nextComputeCycle == current.cycle {
				sum = sum + current.cycle*current.register
				nextComputeCycle += 40
			}
		},
		func(current *cpu) {
			if abs((current.cycle-1)%crtWidth-current.register) <= 1 {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}

			if current.cycle%crtWidth == 0 {
				fmt.Print("\n")
			}
		},
	})

	for scanner.Scan() {
		switch parts := strings.Split(scanner.Text(), " "); parts[0] {
		case "addx":
			value, _ := strconv.ParseInt(parts[1], 10, 64)
			cpu.do(2, value)

		case "noop":
			cpu.do(1, 0)
		}
	}

	fmt.Println("sum", sum)
}

func abs(a int64) int64 {
	if a < 0 {
		return -a
	}

	return a
}
