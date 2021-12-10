package aoc2021

import (
	"testing"
)

func TestLanternfishExample(t *testing.T) {
	RunTest(t, Lanternfish, "testdata/day6_example", 5934, false, 80)
}

func TestLanternfish(t *testing.T) {
	RunTest(t, Lanternfish, "testdata/day6_input", 380758, false, 80)
}

func TestLanternfish256Example(t *testing.T) {
	RunTest(t, Lanternfish, "testdata/day6_example", 26984457539, false, 256)
}

func TestLanternfish256(t *testing.T) {
	RunTest(t, Lanternfish, "testdata/day6_input", 1710623015163, false, 256)
}

func BenchmarkLanternfish(b *testing.B) {
	input := LoadInput("testdata/day6_input")
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		RunBench(b, Lanternfish, input, false, 80)
	}
}

func BenchmarkLanternfish256(b *testing.B) {
	input := LoadInput("testdata/day6_input")
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		RunBench(b, Lanternfish, input, false, 256)
	}
}
