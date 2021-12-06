package aoc2021

import "os"

func LoadInput(inputFile string) []byte {
	data, err := os.ReadFile(inputFile)
	if err != nil {
		panic(err)
	}
	return data
}
