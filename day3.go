package aoc2021

import (
	"bytes"
	"log"
	"os"
	"strconv"
)

func BinaryDiagnostic(inputFile string, debug bool) (int64, error) {
	data, err := os.ReadFile(inputFile)
	if err != nil {
		return 0, err
	}
	size := bytes.IndexByte(data, 10)
	zeros := make([]int, size, size)
	ones := make([]int, size, size)
	i := 0
	for _, c := range data {
		if c == 10 {
			i = 0
			if debug {
				log.Printf("\n")
			}
			continue
		}
		if debug {
			log.Printf("%d:%v", i, c)
		}
		if c == 48 {
			zeros[i]++
		} else if c == 49 {
			ones[i]++
		}
		i++
	}
	if debug {
		log.Printf("zeros: %v", zeros)
		log.Printf("ones:  %v", ones)
	}
	g := make([]byte, size, size)
	e := make([]byte, size, size)
	for i := 0; i < size; i++ {
		if zeros[i] > ones[i] {
			g[i] = 48
			e[i] = 49
		} else if ones[i] > zeros[i] {
			g[i] = 49
			e[i] = 48
		}
	}
	if debug {
		log.Printf("g: %s", g)
		log.Printf("e: %s", e)
	}
	gammaRate, err := strconv.ParseInt(string(g), 2, 64)
	if err != nil {
		return 0, err
	}
	epsilonRate, err := strconv.ParseInt(string(e), 2, 64)
	if err != nil {
		return 0, err
	}
	result := gammaRate * epsilonRate
	return result, nil
}
