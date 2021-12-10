package aoc2021

import (
	"testing"
)

func TestBinaryDiagnosticExample(t *testing.T) {
	RunTest(t, BinaryDiagnostic, "testdata/day3_example", 198, false)
}

func TestBinaryDiagnostic(t *testing.T) {
	RunTest(t, BinaryDiagnostic, "testdata/day3_input", 4001724, false)
}

func TestBinaryDiagnosticPowerConsumptionExample(t *testing.T) {
	RunTest(t, BinaryDiagnosticPowerConsumption, "testdata/day3_example", 230, false)
}

func TestBinaryDiagnosticPowerConsumption(t *testing.T) {
	RunTest(t, BinaryDiagnosticPowerConsumption, "testdata/day3_input", 587895, false)
}

func BenchmarkBinaryDiagnostic(b *testing.B) {
	input := LoadInput("testdata/day3_input")
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		RunBench(b, BinaryDiagnostic, input, false)
	}
}

func BenchmarkBinaryDiagnosticPowerConsumption(b *testing.B) {
	input := LoadInput("testdata/day3_input")
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		RunBench(b, BinaryDiagnosticPowerConsumption, input, false)
	}
}
