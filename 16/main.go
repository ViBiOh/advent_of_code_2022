package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var re = regexp.MustCompile(`(?m)Valve (.*?) has flow rate=(.*?); tunnels? leads? to valves? (.*?)$`)

type valve struct {
	id          string
	connections []string
	rate        int
}

type distance struct {
	id    string
	value int
}

type graph struct {
	id       string
	distance []distance
	rate     int
	value    uint16
}

type path struct {
	current   string
	opened    uint16
	released  int
	remaining int
}

func main() {
	toOpen, allValue, distanceToFirstValve := valvesToGraph(parseValves())

	fmt.Println("part1", part1(openValves(toOpen, distanceToFirstValve, allValue, 30)))
	fmt.Println("part2", part2(openValves(toOpen, distanceToFirstValve, allValue, 26)))
}

func part1(results map[uint16]int) int {
	var output int

	for _, item := range results {
		if item > output {
			output = item
		}
	}

	return output
}

func part2(results map[uint16]int) int {
	var output int

	for current, currentReleased := range results {
		for elephant, elephantReleased := range results {
			if current&elephant == 0 {
				if total := currentReleased + elephantReleased; total > output {
					output = total
				}
			}
		}
	}

	return output
}

func openValves(toOpen map[string]*graph, distanceToFirstValve []distance, allValue uint16, remaining int) map[uint16]int {
	var items []path
	var next []path

	for _, current := range distanceToFirstValve {
		next = append(next, path{
			current:   current.id,
			remaining: remaining - current.value - 1,
		})
	}

	output := make(map[uint16]int)

	for len(next) > 0 {
		items = next
		next = nil

		fmt.Println(len(items))

		for _, item := range items {
			if value, ok := output[item.opened]; !ok || value < item.released {
				output[item.opened] = item.released
			}

			if item.remaining <= 0 {
				continue
			}

			itemValve := toOpen[item.current]
			item.released += item.remaining * itemValve.rate
			item.opened |= itemValve.value

			if item.opened == allValue {
				if value, ok := output[item.opened]; !ok || value < item.released {
					output[item.opened] = item.released
				}

				continue
			}

			for _, currentDistance := range itemValve.distance {
				if item.opened&toOpen[currentDistance.id].value != 0 {
					continue
				}

				next = append(next, path{
					current:   currentDistance.id,
					opened:    item.opened,
					released:  item.released,
					remaining: item.remaining - currentDistance.value - 1,
				})
			}
		}
	}

	return output
}

func parseValves() (map[string]valve, []valve) {
	output := make(map[string]valve)

	scanner := bufio.NewScanner(os.Stdin)

	var valvesToOpen []valve

	for scanner.Scan() {
		matches := re.FindStringSubmatch(scanner.Text())

		var item valve
		item.id = matches[1]
		item.rate, _ = strconv.Atoi(matches[2])
		item.connections = strings.Split(matches[3], ", ")

		output[item.id] = item

		if item.rate > 0 {
			valvesToOpen = append(valvesToOpen, item)
		}
	}

	return output, valvesToOpen
}

func valvesToGraph(valves map[string]valve, valvesToOpen []valve) (map[string]*graph, uint16, []distance) {
	valvesAsGraph := make(map[string]*graph)
	var distanceToFirstValve []distance

	var index int
	var allValue uint16

	firstValve := valves["AA"]

	for _, valve := range valvesToOpen {
		valvesAsGraph[valve.id] = &graph{
			id:    valve.id,
			rate:  valve.rate,
			value: 1 << index,
		}

		allValue |= 1 << index
		index++

		for _, other := range valvesToOpen {
			if valve.id == other.id {
				continue
			}

			valvesAsGraph[valve.id].distance = append(valvesAsGraph[valve.id].distance, distance{
				id:    other.id,
				value: computeDistance(valves, valve, other),
			})
		}

		distanceToFirstValve = append(distanceToFirstValve, distance{
			id:    valve.id,
			value: computeDistance(valves, firstValve, valve),
		})
	}

	return valvesAsGraph, allValue, distanceToFirstValve
}

func computeDistance(valves map[string]valve, source, destination valve) int {
	var items []valve
	next := []valve{
		source,
	}

	visited := make(map[string]bool)
	var count int

	for len(next) > 0 {
		items = next
		next = nil

		for _, item := range items {
			if item.id == destination.id {
				return count
			}

			visited[item.id] = true

			for _, connection := range item.connections {
				if visited[connection] {
					continue
				}

				next = append(next, valves[connection])
			}
		}

		count++
	}

	return -1
}
