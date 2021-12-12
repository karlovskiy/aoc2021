package aoc2021

import "testing"

func TestSmokeBasinExample(t *testing.T) {
	RunTest(t, SmokeBasin, "testdata/day9_example", 15, false)
}

func TestSmokeBasin(t *testing.T) {
	RunTest(t, SmokeBasin, "testdata/day9_input", 591, false)
}

func BenchmarkSmokeBasin(b *testing.B) {
	input := LoadInput("testdata/day9_input")
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		RunBench(b, SmokeBasin, input, false)
	}
}
