package aoc2021

import "log"

func SmokeBasin(input []byte, params ...interface{}) (int, error) {
	debug := params[0].(bool)
	lastIndex := len(input) - 1
	nums := make([][]int, 0, 1024)
	row := make([]int, 0, 128)
	cols := 0
	for k, c := range input {
		if c == 0xa || k == lastIndex {
			if k == lastIndex {
				row = append(row, int(c-0x30))
				cols = len(row)
			}
			nums = append(nums, row)
			row = make([]int, 0, 128)
			continue
		}
		row = append(row, int(c-0x30))
	}
	rows := len(nums)
	if debug {
		log.Printf("nums=%v, rows=%d, cols=%d", nums, rows, cols)
	}
	riskLevel := 0
	for i := 0; i < cols; i++ {
		for j := 0; j < rows; j++ {
			li, ri, uj, dj := i-1, i+1, j-1, j+1
			val := nums[j][i]
			if (li < 0 || val < nums[j][li]) && (ri >= cols || val < nums[j][ri]) &&
				(uj < 0 || val < nums[uj][i]) && (dj >= rows || val < nums[dj][i]) {
				if debug {
					log.Printf("%d:%d %d", j, i, val)
				}
				riskLevel += val + 1
			}
		}
	}
	return riskLevel, nil
}
