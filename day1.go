package aoc2021

import (
	"log"
	"os"
	"strconv"
)

func SonarSweep(inputFile string, debug bool) (int, error) {
	data, err := os.ReadFile(inputFile)
	if err != nil {
		return 0, nil
	}
	lastRow := len(data) - 1
	num := make([]byte, 0, 16)
	i, prev, result := 0, -1, -1
	for j, c := range data {
		if c == 10 {
			i = 0
			cur, err := strconv.Atoi(string(num))
			if err != nil {
				return 0, err
			}
			if cur > prev {
				result++
			}
			prev = cur
			num = num[:0]
			if debug {
				log.Printf("\n")
			}
			continue
		}
		if debug {
			log.Printf("%d:%c", i, c)
		}
		num = append(num, c)
		i++
		if j == lastRow {
			cur, err := strconv.Atoi(string(num))
			if err != nil {
				return 0, err
			}
			if cur > prev {
				result++
			}
		}
	}
	return result, nil
}
