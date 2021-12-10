package aoc2021

import (
	"testing"
)

func TestDiveExample(t *testing.T) {
	RunTest(t, Dive, "testdata/day2_example", 150, false)
}

func TestDive(t *testing.T) {
	RunTest(t, Dive, "testdata/day2_input", 1507611, false)
}

func TestDiveWithAimExample(t *testing.T) {
	RunTest(t, Dive, "testdata/day2_example", 900, true)
}

func TestDiveWithAim(t *testing.T) {
	RunTest(t, Dive, "testdata/day2_input", 1880593125, true)
}

func BenchmarkDive(b *testing.B) {
	input := LoadInput("testdata/day2_input")
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		RunBench(b, Dive, input, false)
	}
}

func BenchmarkDiveWithAim(b *testing.B) {
	input := LoadInput("testdata/day2_input")
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		RunBench(b, Dive, input, true)
	}
}
