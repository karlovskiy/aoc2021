package aoc2021

import (
	"bytes"
	"log"
	"os"
)

func BinaryDiagnostic(inputFile string, debug bool) (int, error) {
	data, err := os.ReadFile(inputFile)
	if err != nil {
		return 0, err
	}
	cols := bytes.IndexByte(data, 10)
	lastRow := (len(data)+1)/(cols+1) - 1
	ones := make([]int, cols)
	i, j, gamma, epsilon := 0, 0, 0, 0
	for _, c := range data {
		if c == 10 {
			i = 0
			j++
			if debug {
				log.Printf("\n")
			}
			continue
		}
		if debug {
			log.Printf("%d:%v", i, c)
		}
		if c == 49 {
			ones[i]++
		}
		if j == lastRow {
			bitPosition := cols - i - 1
			if ones[i]*2 > lastRow+1 {
				gamma |= 1 << bitPosition
			} else {
				epsilon |= 1 << bitPosition
			}
		}
		i++
	}
	if debug {
		log.Printf("ones:%v, gamma: %b, epsilon: %b", ones, gamma, epsilon)
	}
	result := gamma * epsilon
	return result, nil
}
