package aoc2021

import (
	"testing"
)

func TestSonarSweepExample(t *testing.T) {
	result, err := SonarSweep("testdata/day1_example", false)
	if err != nil {
		t.Fatalf("error testing example input file: %v", err)
	}
	if result != 7 {
		t.Fatalf("wrong result for example input file, actual=%d, expected=%d", result, 7)
	}
}

func TestSonarSweep(t *testing.T) {
	result, err := SonarSweep("testdata/day1_input", false)
	if err != nil {
		t.Fatalf("error testing input file: %v", err)
	}
	if result != 1791 {
		t.Fatalf("wrong result for input file, actual=%d, expected=%d", result, 1791)
	}
}

/**
goos: linux
goarch: amd64
pkg: github.com/karlovskiy/aoc2021
cpu: Intel(R) Core(TM) i7-4770K CPU @ 3.50GHz
BenchmarkSonarSweep-8         	   16154	     74222 ns/op
*/
func BenchmarkSonarSweep(b *testing.B) {
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		_, err := SonarSweep("testdata/day1_input", false)
		if err != nil {
			b.Fatalf("error testing input file: %v", err)
		}
	}
}
