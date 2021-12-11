package aoc2021

import "testing"

func TestTreacheryOfWhalesExample(t *testing.T) {
	RunTest(t, TreacheryOfWhales, "testdata/day7_example", 37, false, false)
}

func TestTreacheryOfWhales(t *testing.T) {
	RunTest(t, TreacheryOfWhales, "testdata/day7_input", 340056, false, false)
}

func TestTreacheryOfWhalesExpensiveExample(t *testing.T) {
	RunTest(t, TreacheryOfWhales, "testdata/day7_example", 168, false, true)
}

func TestTreacheryOfWhalesExpensive(t *testing.T) {
	RunTest(t, TreacheryOfWhales, "testdata/day7_input", 96592275, false, true)
}

func BenchmarkTreacheryOfWhales(b *testing.B) {
	input := LoadInput("testdata/day7_input")
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		RunBench(b, TreacheryOfWhales, input, false, false)
	}
}

func BenchmarkTreacheryOfWhalesExpensive(b *testing.B) {
	input := LoadInput("testdata/day7_input")
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		RunBench(b, TreacheryOfWhales, input, false, true)
	}
}
