package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	if !scanner.Scan() {
		return
	}

	line := scanner.Text()

	var moving string
	var packetStart, messageStart int

	for index, value := range line {
		if position := strings.IndexRune(moving, value); position != -1 {
			moving = moving[position+1:]
		}

		moving += string(value)

		if packetStart == 0 && len(moving) == 4 {
			packetStart = index + 1
		} else if len(moving) == 14 {
			messageStart = index + 1
			break
		}
	}

	fmt.Println("packet", packetStart)
	fmt.Println("message", messageStart)
}
