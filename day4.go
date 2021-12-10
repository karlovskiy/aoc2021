package aoc2021

import (
	"errors"
	"log"
	"strconv"
)

func GiantSquid(input []byte, params ...interface{}) (int, error) {
	debug := params[0].(bool)
	last := params[1].(bool)
	lastInputIndex := len(input) - 1
	nums := make([]int, 0, 32)
	numSymbols := make([]byte, 0, 16)
	boards := make(map[int]*[5][5]int)
	i, b, br, bc := 0, 0, 0, 0
	for j, c := range input {
		if i == 0 {
			if c == 0x2c || c == 0xa {
				number, err := strconv.Atoi(string(numSymbols))
				if err != nil {
					return 0, err
				}
				numSymbols = numSymbols[:0]
				nums = append(nums, number)
				if c == 0xa {
					i++
				}
				continue
			}
			numSymbols = append(numSymbols, c)
			continue
		}
		if c == 0xa && len(numSymbols) == 0 {
			if i != 1 {
				b++
			}
			boards[b] = &[5][5]int{}
			i++
			br, bc = 0, 0
			continue
		}
		if c == 0x20 || c == 0xa || j == lastInputIndex {
			if j == lastInputIndex {
				numSymbols = append(numSymbols, c)
			}
			if c == 0x20 && len(numSymbols) == 0 {
				continue
			}
			number, err := strconv.Atoi(string(numSymbols))
			if err != nil {
				return 0, err
			}
			numSymbols = numSymbols[:0]
			boards[b][br][bc] = number
			if c == 0x20 {
				bc++
			}
			if c == 0xa {
				i++
				br++
				bc = 0
			}
			continue
		}
		numSymbols = append(numSymbols, c)
	}
	if debug {
		log.Printf("nums: %v", nums)
		log.Printf("boards: %v", boards)
	}
	boardsLen := len(boards)
	var winner *[5][5]int
	var winNum int
	rowsMarked, colsMarked := make([][5]int, boardsLen), make([][5]int, boardsLen)
winnerFound:
	for _, num := range nums {
		for k, board := range boards {
		nextBoard:
			for i := 0; i < 5; i++ {
				for j := 0; j < 5; j++ {
					val := board[i][j]
					if num == val {
						val = -1
						boards[k][i][j] = val
						rowsMarked[k][i]++
						colsMarked[k][j]++
					}
					if i == 4 && rowsMarked[k][j] == 5 {
						if last && len(boards) > 1 {
							delete(boards, k)
							break nextBoard
						}
						winner = boards[k]
						winNum = num
						break winnerFound
					}
				}
				if colsMarked[k][i] == 5 {
					if last && len(boards) > 1 {
						delete(boards, k)
						break nextBoard
					}
					winner = boards[k]
					winNum = num
					break winnerFound
				}
			}
		}
	}
	if winner == nil {
		return 0, errors.New("winner board was not found")
	}
	if debug {
		log.Printf("winner: num=%d, board=%v", winNum, winner)
	}
	sum := 0
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			val := winner[i][j]
			if val > 0 {
				sum += val
			}
		}
	}
	result := sum * winNum
	if debug {
		log.Printf("result: %d", result)
	}
	return result, nil
}
