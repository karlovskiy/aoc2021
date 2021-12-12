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
	nums := make([]int, 0, 4096)
	max := 0
	m := make(map[int]int)
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
			if val > max {
				max = val
			}
			if i == 0 {
				m[0] = 0
			} else {
				m[i] = m[i-1] + i
			}
		}
	}
	if debug {
		log.Printf("nums=%v, max=%d", nums, max)
	}
	numsLen := len(nums)
	min := math.MaxInt32
	for i := 1; i < max; i++ {
		fuel := 0
		skip := false
		for j := 0; j < numsLen; j++ {
			if expensive {
				moves := Abs(i - nums[j])
				val, ok := m[moves]
				if !ok {
					for k := 0; k < moves; k++ {
						val += k + 1
					}
					m[moves] = val
				}
				fuel += val
			} else {
				fuel += Abs(i - nums[j])
			}
			if min < fuel {
				skip = true
				break
			}
		}
		if fuel < min {
			min = fuel
		}
		if debug {
			log.Printf("%d: fuel=%d, skip=%t", i, fuel, skip)
		}
	}
	return min, nil
}
