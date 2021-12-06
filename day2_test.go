package aoc2021

import (
	"testing"
)

func TestDiveExample(t *testing.T) {
	input := LoadInput("testdata/day2_example")
	result, err := Dive(input, false)
	if err != nil {
		t.Fatalf("error testing example input file: %v", err)
	}
	if result != 150 {
		t.Fatalf("wrong result for example input file, actual=%d, expected=%d", result, 150)
	}
}

func TestDive(t *testing.T) {
	input := LoadInput("testdata/day2_input")
	result, err := Dive(input, false)
	if err != nil {
		t.Fatalf("error testing input file: %v", err)
	}
	if result != 1507611 {
		t.Fatalf("wrong result for input file, actual=%d, expected=%d", result, 1507611)
	}
}

func TestDiveWithAimExample(t *testing.T) {
	input := LoadInput("testdata/day2_example")
	result, err := Dive(input, true)
	if err != nil {
		t.Fatalf("error testing example input file: %v", err)
	}
	if result != 900 {
		t.Fatalf("wrong result for example input file, actual=%d, expected=%d", result, 900)
	}
}

func TestDiveWithAim(t *testing.T) {
	input := LoadInput("testdata/day2_input")
	result, err := Dive(input, true)
	if err != nil {
		t.Fatalf("error testing input file: %v", err)
	}
	if result != 1880593125 {
		t.Fatalf("wrong result for input file, actual=%d, expected=%d", result, 1880593125)
	}
}

/*
goos: linux
goarch: amd64
pkg: github.com/karlovskiy/aoc2021
cpu: Intel(R) Core(TM) i7-4770K CPU @ 3.50GHz
BenchmarkDive-8                         	   10000	    118782 ns/op
*/
func BenchmarkDive(b *testing.B) {
	input := LoadInput("testdata/day2_input")
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		_, err := Dive(input, false)
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
BenchmarkDiveWithAim-8                  	   10000	    121681 ns/op
*/
func BenchmarkDiveWithAim(b *testing.B) {
	input := LoadInput("testdata/day2_input")
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		_, err := Dive(input, true)
		if err != nil {
			b.Fatalf("error testing input file: %v", err)
		}
	}
}
