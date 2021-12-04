package aoc2021

import (
	"testing"
)

func TestExample(t *testing.T) {
	pwrCnsmpt, err := BinaryDiagnostic("testdata/day3_example", false)
	if err != nil {
		t.Fatalf("error testing example input file: %v", err)
	}
	if pwrCnsmpt != 198 {
		t.Fatalf("wrong result for example input file, actual=%d, expected=%d", pwrCnsmpt, 198)
	}
}

func TestInput(t *testing.T) {
	pwrCnsmpt, err := BinaryDiagnostic("testdata/day3_input", false)
	if err != nil {
		t.Fatalf("error testing example input file: %v", err)
	}
	if pwrCnsmpt != 4001724 {
		t.Fatalf("wrong result for example input file, actual=%d, expected=%d", pwrCnsmpt, 198)
	}
}
