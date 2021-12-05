package aoc2021

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func Dive(inputFile string) (int, error) {
	f, err := os.Open(inputFile)
	if err != nil {
		return 0, nil
	}
	hPos, depth := 0, 0
	scanner := bufio.NewScanner(f)
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
		case "down":
			depth += val
		case "up":
			depth -= val
		}
	}
	result := hPos * depth
	return result, nil
}
