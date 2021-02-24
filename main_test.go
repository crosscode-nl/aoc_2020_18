package main

import "testing"

var TestCalculatorTestData = []struct {
	in string
	out int
}{
	{"1 + 2 * 3 + 4 * 5 + 6", 71},
	{"5 + (8 * 3 + 9 + 3 * 4 * 3)", 437},
	{"5 * 9 * (7 * 3 * 3 + 9 * 3 + (8 + 6 * 4))", 12240},
	{"((2 + 4 * 9) * (6 + 9 * 8 + 6) + 6) + 2 + 4 * 2", 13632},
}
func TestCalculator(t *testing.T) {
	for _, td := range TestCalculatorTestData {
		t.Run(td.in, func(t *testing.T) {
			r :=  calculate1(td.in)
			if r != td.out {
				t.Errorf("got %d, expected %d", r, td.out)
			}
		})
	}
}
