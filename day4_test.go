package aoc2021

import "testing"

func TestGiantSquidExample(t *testing.T) {
	RunTest(t, GiantSquid, "testdata/day4_example", 4512, false, false)
}

func TestGiantSquid(t *testing.T) {
	RunTest(t, GiantSquid, "testdata/day4_input", 45031, false, false)
}

func TestGiantSquidLastExample(t *testing.T) {
	RunTest(t, GiantSquid, "testdata/day4_example", 1924, false, true)
}

func TestGiantSquidLast(t *testing.T) {
	RunTest(t, GiantSquid, "testdata/day4_input", 2568, false, true)
}

func BenchmarkGiantSquid(b *testing.B) {
	input := LoadInput("testdata/day4_input")
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		RunBench(b, GiantSquid, input, false, false)
	}
}

func BenchmarkGiantSquidLast(b *testing.B) {
	input := LoadInput("testdata/day4_input")
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		RunBench(b, GiantSquid, input, false, true)
	}
}
