package aoc2021

import (
	"testing"
)

func TestBinaryDiagnosticExample(t *testing.T) {
	input := LoadInput("testdata/day3_example")
	result, err := BinaryDiagnostic(input, false)
	if err != nil {
		t.Fatalf("error testing example input file: %v", err)
	}
	if result != 198 {
		t.Fatalf("wrong result for example input file, actual=%d, expected=%d", result, 198)
	}
}

func TestBinaryDiagnostic(t *testing.T) {
	input := LoadInput("testdata/day3_input")
	result, err := BinaryDiagnostic(input, false)
	if err != nil {
		t.Fatalf("error testing input file: %v", err)
	}
	if result != 4001724 {
		t.Fatalf("wrong result for input file, actual=%d, expected=%d", result, 198)
	}
}

func TestBinaryDiagnosticPowerConsumptionExample(t *testing.T) {
	input := LoadInput("testdata/day3_example")
	result, err := BinaryDiagnosticPowerConsumption(input, false)
	if err != nil {
		t.Fatalf("error testing example input file: %v", err)
	}
	if result != 230 {
		t.Fatalf("wrong result for example input file, actual=%d, expected=%d", result, 230)
	}
}

func TestBinaryDiagnosticPowerConsumption(t *testing.T) {
	input := LoadInput("testdata/day3_input")
	result, err := BinaryDiagnosticPowerConsumption(input, false)
	if err != nil {
		t.Fatalf("error testing input file: %v", err)
	}
	if result != 587895 {
		t.Fatalf("wrong result for input file, actual=%d, expected=%d", result, 587895)
	}
}

/**
goos: linux
goarch: amd64
pkg: github.com/karlovskiy/aoc2021
cpu: Intel(R) Core(TM) i7-4770K CPU @ 3.50GHz
BenchmarkBinaryDiagnostic-8                   	   21386	     54882 ns/op
*/
func BenchmarkBinaryDiagnostic(b *testing.B) {
	input := LoadInput("testdata/day3_input")
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		_, err := BinaryDiagnostic(input, false)
		if err != nil {
			b.Fatalf("error testing input file: %v", err)
		}
	}
}

/**
goos: linux
goarch: amd64
pkg: github.com/karlovskiy/aoc2021
cpu: Intel(R) Core(TM) i7-4770K CPU @ 3.50GHz
BenchmarkBinaryDiagnosticPowerConsumption-8   	    1764	    733968 ns/op
*/
func BenchmarkBinaryDiagnosticPowerConsumption(b *testing.B) {
	input := LoadInput("testdata/day3_input")
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		_, err := BinaryDiagnosticPowerConsumption(input, false)
		if err != nil {
			b.Fatalf("error testing input file: %v", err)
		}
	}
}
