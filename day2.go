package aoc2021

import (
	"bufio"
	"bytes"
	"strconv"
	"strings"
)

func Dive(input []byte, withAim bool) (int, error) {
	hPos, depth, aim := 0, 0, 0
	scanner := bufio.NewScanner(bytes.NewReader(input))
	for scanner.Scan() {
		line := scanner.Text()
		tokens := strings.Split(line, " ")
		command := tokens[0]
		val, err := strconv.Atoi(tokens[1])
		if err != nil {
			return 0, err
		}
		switch command {
		case "forward":
			hPos += val
			if withAim && aim > 0 {
				depth += aim * val
			}
		case "down":
			if withAim {
				aim += val
			} else {
				depth += val
			}
		case "up":
			if withAim {
				aim -= val
			} else {
				depth -= val
			}
		}
	}
	result := hPos * depth
	return result, nil
}
