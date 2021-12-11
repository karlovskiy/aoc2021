package aoc2021

import (
	"bufio"
	"bytes"
	"log"
	"math"
	"strconv"
	"strings"
)

func TreacheryOfWhales(input []byte, params ...interface{}) (int, error) {
	debug := params[0].(bool)
	expensive := params[1].(bool)
	nums := make([]int, 0, 32)
	maxPos := 0
	scanner := bufio.NewScanner(bytes.NewReader(input))
	for scanner.Scan() {
		line := scanner.Text()
		tokens := strings.Split(line, ",")
		for i := 0; i < len(tokens); i++ {
			val, err := strconv.Atoi(tokens[i])
			if err != nil {
				return 0, err
			}
			nums = append(nums, val)
			if val > maxPos {
				maxPos = val
			}
		}
	}
	if debug {
		log.Printf("nums=%v", nums)
	}
	numsLen := len(nums)
	min := math.MaxInt32
	for i := 1; i < maxPos; i++ {
		fuel := 0
		for j := 0; j < numsLen; j++ {
			if expensive {
				for k := 1; k < Abs(i-nums[j])+1; k++ {
					fuel += k
				}
			} else {
				fuel += Abs(i - nums[j])
			}
			if min < fuel {
				break
			}
		}
		if fuel < min {
			min = fuel
		}
		if debug {
			log.Printf("%d: fuel=%d", i, fuel)
		}
	}
	return min, nil
}
