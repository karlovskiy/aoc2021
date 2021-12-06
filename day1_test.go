package aoc2021

import (
	"testing"
)

func TestSonarSweepExample(t *testing.T) {
	input := LoadInput("testdata/day1_example")
	result, err := SonarSweep(input, false)
	if err != nil {
		t.Fatalf("error testing example input file: %v", err)
	}
	if result != 7 {
		t.Fatalf("wrong result for example input file, actual=%d, expected=%d", result, 7)
	}
}

func TestSonarSweep(t *testing.T) {
	input := LoadInput("testdata/day1_input")
	result, err := SonarSweep(input, false)
	if err != nil {
		t.Fatalf("error testing input file: %v", err)
	}
	if result != 1791 {
		t.Fatalf("wrong result for input file, actual=%d, expected=%d", result, 1791)
	}
}

func TestSonarSweepThreeMeasurementExample(t *testing.T) {
	input := LoadInput("testdata/day1_example")
	result, err := SonarSweepThreeMeasurement(input, false)
	if err != nil {
		t.Fatalf("error testing example input file: %v", err)
	}
	if result != 5 {
		t.Fatalf("wrong result for example input file, actual=%d, expected=%d", result, 5)
	}
}

func TestSonarSweepThreeMeasurement(t *testing.T) {
	input := LoadInput("testdata/day1_input")
	result, err := SonarSweepThreeMeasurement(input, false)
	if err != nil {
		t.Fatalf("error testing input file: %v", err)
	}
	if result != 1822 {
		t.Fatalf("wrong result for input file, actual=%d, expected=%d", result, 1822)
	}
}

/**
goos: linux
goarch: amd64
pkg: github.com/karlovskiy/aoc2021
cpu: Intel(R) Core(TM) i7-4770K CPU @ 3.50GHz
BenchmarkSonarSweep-8                   	   20949	     56673 ns/op
*/
func BenchmarkSonarSweep(b *testing.B) {
	input := LoadInput("testdata/day1_input")
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		_, err := SonarSweep(input, false)
		if err != nil {
			b.Fatalf("error testing input file: %v", err)
		}
	}
}

/*
goos: linux
goarch: amd64
pkg: github.com/karlovskiy/aoc2021
cpu: Intel(R) Core(TM) i7-4770K CPU @ 3.50GHz
BenchmarkSonarSweepThreeMeasurement-8   	   18062	     67118 ns/op

*/
func BenchmarkSonarSweepThreeMeasurement(b *testing.B) {
	input := LoadInput("testdata/day1_input")
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		_, err := SonarSweepThreeMeasurement(input, false)
		if err != nil {
			b.Fatalf("error testing input file: %v", err)
		}
	}
}
