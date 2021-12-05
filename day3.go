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
	cols := bytes.IndexByte(data, 10)
	lastRow := (len(data)+1)/(cols+1) - 1
	zeros, ones := make([]int, cols, cols), make([]int, cols, cols)
	gammaBits, epsilonBits := make([]byte, cols, cols), make([]byte, cols, cols)
	i, j := 0, 0
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
		if c == 48 {
			zeros[i]++
		} else if c == 49 {
			ones[i]++
		}
		if j == lastRow {
			if zeros[i] > ones[i] {
				gammaBits[i] = 48
				epsilonBits[i] = 49
			} else if ones[i] > zeros[i] {
				gammaBits[i] = 49
				epsilonBits[i] = 48
			}
		}
		i++
	}
	if debug {
		log.Printf("zeros: %v, ones: %v", zeros, ones)
		log.Printf("g: %s, e: %s", gammaBits, epsilonBits)
	}
	gammaRate, err := strconv.ParseInt(string(gammaBits), 2, 64)
	if err != nil {
		return 0, err
	}
	epsilonRate, err := strconv.ParseInt(string(epsilonBits), 2, 64)
	if err != nil {
		return 0, err
	}
	result := gammaRate * epsilonRate
	return result, nil
}
