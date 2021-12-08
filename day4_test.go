package aoc2021

import "testing"

func TestGiantSquidExample(t *testing.T) {
	input := LoadInput("testdata/day4_example")
	result, err := GiantSquid(input, false, false)
	if err != nil {
		t.Fatalf("error testing example input file: %v", err)
	}
	if result != 4512 {
		t.Fatalf("wrong result for example input file, actual=%d, expected=%d", result, 4512)
	}
}

func TestGiantSquid(t *testing.T) {
	input := LoadInput("testdata/day4_input")
	result, err := GiantSquid(input, false, false)
	if err != nil {
		t.Fatalf("error testing input file: %v", err)
	}
	if result != 45031 {
		t.Fatalf("wrong result for input file, actual=%d, expected=%d", result, 45031)
	}
}

func TestGiantSquidLastExample(t *testing.T) {
	input := LoadInput("testdata/day4_example")
	result, err := GiantSquid(input, true, false)
	if err != nil {
		t.Fatalf("error testing example input file: %v", err)
	}
	if result != 1924 {
		t.Fatalf("wrong result for example input file, actual=%d, expected=%d", result, 1924)
	}
}

func TestGiantSquidLast(t *testing.T) {
	input := LoadInput("testdata/day4_input")
	result, err := GiantSquid(input, true, false)
	if err != nil {
		t.Fatalf("error testing input file: %v", err)
	}
	if result != 2568 {
		t.Fatalf("wrong result for input file, actual=%d, expected=%d", result, 2568)
	}
}

/**
goos: linux
goarch: amd64
pkg: github.com/karlovskiy/aoc2021
cpu: Intel(R) Core(TM) i7-4770K CPU @ 3.50GHz
BenchmarkGiantSquid-8   	    8078	    150464 ns/op
*/
func BenchmarkGiantSquid(b *testing.B) {
	input := LoadInput("testdata/day4_input")
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		_, err := GiantSquid(input, false, false)
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
BenchmarkGiantSquidLast-8   	    2598	    437367 ns/op
*/
func BenchmarkGiantSquidLast(b *testing.B) {
	input := LoadInput("testdata/day4_input")
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		_, err := GiantSquid(input, true, false)
		if err != nil {
			b.Fatalf("error testing input file: %v", err)
		}
	}
}
