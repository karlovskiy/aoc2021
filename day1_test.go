package aoc2021

import (
	"testing"
)

func TestSonarSweepExample(t *testing.T) {
	RunTest(t, SonarSweep, "testdata/day1_example", 7, false)
}

func TestSonarSweep(t *testing.T) {
	RunTest(t, SonarSweep, "testdata/day1_input", 1791, false)
}

func TestSonarSweepThreeMeasurementExample(t *testing.T) {
	RunTest(t, SonarSweepThreeMeasurement, "testdata/day1_example", 5, false)
}

func TestSonarSweepThreeMeasurement(t *testing.T) {
	RunTest(t, SonarSweepThreeMeasurement, "testdata/day1_input", 1822, false)
}

func BenchmarkSonarSweep(b *testing.B) {
	input := LoadInput("testdata/day1_input")
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		RunBench(b, SonarSweep, input, false)
	}
}

func BenchmarkSonarSweepThreeMeasurement(b *testing.B) {
	input := LoadInput("testdata/day1_input")
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		RunBench(b, SonarSweepThreeMeasurement, input, false)
	}
}
