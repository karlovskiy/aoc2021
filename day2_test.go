package aoc2021

import (
	"testing"
)

func TestDiveExample(t *testing.T) {
	result, err := Dive("testdata/day2_example")
	if err != nil {
		t.Fatalf("error testing example input file: %v", err)
	}
	if result != 150 {
		t.Fatalf("wrong result for example input file, actual=%d, expected=%d", result, 150)
	}
}

func TestDive(t *testing.T) {
	result, err := Dive("testdata/day2_input")
	if err != nil {
		t.Fatalf("error testing input file: %v", err)
	}
	if result != 1507611 {
		t.Fatalf("wrong result for input file, actual=%d, expected=%d", result, 1507611)
	}
}

/*
goos: linux
goarch: amd64
pkg: github.com/karlovskiy/aoc2021
cpu: Intel(R) Core(TM) i7-4770K CPU @ 3.50GHz
BenchmarkDive-8               	    9038	    136063 ns/op
*/
func BenchmarkDive(b *testing.B) {
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		_, err := Dive("testdata/day2_input")
		if err != nil {
			b.Fatalf("error testing input file: %v", err)
		}
	}
}
