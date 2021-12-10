package aoc2021

import (
	"log"
	"strconv"
)

func Lanternfish(input []byte, params ...interface{}) (int, error) {
	debug := params[0].(bool)
	days := params[1].(int)
	lastInputIndex := len(input) - 1
	numSymbols := make([]byte, 0)
	nums := make([]int, 9)
	for j, c := range input {
		if c == 0x2c || j == lastInputIndex {
			if j == lastInputIndex {
				numSymbols = append(numSymbols, c)
			}
			number, err := strconv.Atoi(string(numSymbols))
			if err != nil {
				return 0, err
			}
			numSymbols = numSymbols[:0]
			nums[number]++
			continue
		}
		numSymbols = append(numSymbols, c)
	}
	if debug {
		log.Printf("nums=%v", nums)
	}

	for k := 0; k < days; k++ {
		dayKNums := make([]int, 9)
		for i, num := range nums {
			if i == 0 {
				dayKNums[6] += num
				dayKNums[8] += num
			} else {
				dayKNums[i-1] += num
			}
		}
		nums = dayKNums
	}
	result := 0
	for _, num := range nums {
		result += num
	}
	if debug {
		log.Printf("result=%v", result)
	}
	return result, nil
}
