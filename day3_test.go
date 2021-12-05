package aoc2021

import (
	"testing"
)

func TestBinaryDiagnosticExample(t *testing.T) {
	pwrCnsmpt, err := BinaryDiagnostic("testdata/day3_example", false)
	if err != nil {
		t.Fatalf("error testing example input file: %v", err)
	}
	if pwrCnsmpt != 198 {
		t.Fatalf("wrong result for example input file, actual=%d, expected=%d", pwrCnsmpt, 198)
	}
}

func TestBinaryDiagnostic(t *testing.T) {
	pwrCnsmpt, err := BinaryDiagnostic("testdata/day3_input", false)
	if err != nil {
		t.Fatalf("error testing input file: %v", err)
	}
	if pwrCnsmpt != 4001724 {
		t.Fatalf("wrong result for input file, actual=%d, expected=%d", pwrCnsmpt, 198)
	}
}

/**
goos: linux
goarch: amd64
pkg: github.com/karlovskiy/aoc2021
cpu: Intel(R) Core(TM) i7-4770K CPU @ 3.50GHz
BenchmarkInput-8   	   17738	     70342 ns/op
*/
func BenchmarkBinaryDiagnostic(b *testing.B) {
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		_, err := BinaryDiagnostic("testdata/day3_input", false)
		if err != nil {
			b.Fatalf("error testing input file: %v", err)
		}
	}
}
