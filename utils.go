package aoc2021

import (
	"os"
	"testing"
)

func LoadInput(inputFile string) []byte {
	data, err := os.ReadFile(inputFile)
	if err != nil {
		panic(err)
	}
	return data
}

func RunTest(t *testing.T, tf testFunc, inputFile string, expected int, params ...interface{}) {
	input := LoadInput(inputFile)
	actual, err := tf(input, params...)
	if err != nil {
		t.Fatalf("error testing: %v", err)
	}
	if actual != expected {
		t.Fatalf("wrong result, actual=%d, expected=%d", actual, expected)
	}
}

func RunBench(b *testing.B, tf testFunc, input []byte, params ...interface{}) {
	_, err := tf(input, params...)
	if err != nil {
		b.Fatalf("error benchmarking: %v", err)
	}
}

type testFunc func([]byte, ...interface{}) (int, error)
