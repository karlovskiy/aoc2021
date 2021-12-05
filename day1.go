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

func SonarSweepThreeMeasurement(inputFile string, debug bool) (int, error) {
	data, err := os.ReadFile(inputFile)
	if err != nil {
		return 0, nil
	}
	num := make([]byte, 0, 16)
	lastDataIndex := len(data) - 1
	windows := []*window{{
		index: 0,
		val:   0,
	}, {
		index: -1,
		val:   0,
	}, {
		index: -2,
		val:   0,
	}}
	prev, result := 0, -1 /*no previous sum*/
	for j, c := range data {
		if c == 10 || j == lastDataIndex {
			if j == lastDataIndex {
				num = append(num, c)
			}
			cur, err := strconv.Atoi(string(num))
			if err != nil {
				return 0, err
			}
			if debug {
				log.Printf("%d", cur)
			}
			for _, w := range windows {
				if w.index >= 0 {
					w.val += cur
				}
			}
			for i, w := range windows {
				if w.index == 2 {
					if debug {
						log.Printf("%d:%d", i+1, w.val)
					}
					if w.val > prev {
						result++
					}
					prev = w.val
					w.index = -1
					w.val = 0
				}
				w.index++
			}
			if debug {
				log.Printf("\n")
			}
			num = num[:0]
			continue
		}
		num = append(num, c)
	}
	return result, nil
}

type window struct {
	index int
	val   int
}
