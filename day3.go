package aoc2021

import (
	"bytes"
	"errors"
	"log"
)

func BinaryDiagnostic(input []byte, debug bool) (int, error) {
	cols := bytes.IndexByte(input, 10)
	lastRow := (len(input)+1)/(cols+1) - 1
	ones := make([]int, cols)
	i, j, gamma, epsilon := 0, 0, 0, 0
	for _, c := range input {
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

func BinaryDiagnosticPowerConsumption(input []byte, debug bool) (int, error) {
	cols := bytes.IndexByte(input, 10)
	rows := (len(input) + 1) / (cols + 1)
	oxygenIndex, co2Index := -1, -1
	values := make([]int, rows)
	oxygenBitIndexes := make(map[int]byte, rows)
	co2BitIndexes := make(map[int]byte, rows)
	var mostOxygen, leastCO2 byte
	for j := 0; j < cols; j++ {
		oxygenZeros := 0
		co2Zeros := 0
		for i := 0; i < rows; i++ {
			index := j + i*(cols+1)
			val := input[index]
			if debug {
				log.Printf("%d:%d:%d:%v", j, i, index, val)
			}
			if val == 49 {
				values[i] |= 1 << (cols - j - 1)
			}
			if oxygenIndex == -1 {
				prevOxygen, ok := oxygenBitIndexes[i]
				if ok || j == 0 {
					if mostOxygen != 0 && mostOxygen != prevOxygen {
						if len(oxygenBitIndexes) == 1 {
							oxygenIndex = i
						} else {
							delete(oxygenBitIndexes, i)
						}
					} else {
						oxygenBitIndexes[i] = val
						if val == 49 {
							oxygenZeros++
						}
					}
				}
			}

			if co2Index == -1 {
				prevCO2, ok := co2BitIndexes[i]
				if ok || j == 0 {
					if leastCO2 != 0 && leastCO2 != prevCO2 {
						if len(co2BitIndexes) == 1 {
							co2Index = i
						} else {
							delete(co2BitIndexes, i)
						}
					} else {
						co2BitIndexes[i] = val
						if val == 49 {
							co2Zeros++
						}
					}
				}
			}
		}
		if oxygenZeros*2 >= len(oxygenBitIndexes) {
			mostOxygen = 49
		} else {
			mostOxygen = 48
		}
		if co2Zeros*2 < len(co2BitIndexes) {
			leastCO2 = 49
		} else {
			leastCO2 = 48
		}
		if j == cols-1 {
			if oxygenIndex == -1 {
				for index, v := range oxygenBitIndexes {
					if mostOxygen == v {
						oxygenIndex = index
					}
				}
			}
			if co2Index == -1 {
				for index, v := range co2BitIndexes {
					if leastCO2 == v {
						co2Index = index
					}
				}
			}
		}
		if debug {
			log.Printf("%d: oxygenBitIndexes:%v, oxygenZeros:%d, mostOxygen:%v", j, oxygenBitIndexes, oxygenZeros, mostOxygen)
			log.Printf("%d: co2BitIndexes:%v, co2Zeros:%d, leastCO2:%v", j, co2BitIndexes, co2Zeros, leastCO2)
		}
	}
	if debug {
		log.Printf("%v", values)
	}
	if oxygenIndex == -1 {
		return 0, errors.New("oxygen generator rating not found")
	}
	if co2Index == -1 {
		return 0, errors.New("CO2 scrubber rating not found")
	}
	oxygen := values[oxygenIndex]
	co2 := values[co2Index]
	if debug {
		log.Printf("oxygen: %d, co2: %d", oxygen, co2)
	}
	return oxygen * co2, nil
}
