package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type position struct {
	x, y, z int
}

func (p position) touch(other position) bool {
	if other.x == p.x {
		if other.y == p.y && abs(other.z-p.z) == 1 {
			return true
		}
		if other.z == p.z && abs(other.y-p.y) == 1 {
			return true
		}
	}

	if other.y == p.y && other.z == p.z && abs(other.x-p.x) == 1 {
		return true
	}

	return false
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	cubes := make(map[position]bool)

	var facesCovered int
	var bound position

	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), ",")

		var new position
		new.x, _ = strconv.Atoi(parts[0])
		new.y, _ = strconv.Atoi(parts[1])
		new.z, _ = strconv.Atoi(parts[2])

		for cube := range cubes {
			if new.touch(cube) {
				facesCovered++
			}
		}

		if new.x > bound.x {
			bound.x = new.x
		}
		if new.y > bound.y {
			bound.y = new.y
		}
		if new.z > bound.z {
			bound.z = new.z
		}

		cubes[new] = true
	}

	sum := len(cubes)*6 - facesCovered*2

	fmt.Println("part1", sum)
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
