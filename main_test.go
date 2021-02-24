package main

import "testing"

var TestCalculatorOldPrecedenceData = []struct {
	in string
	out int
}{
	{"1 + (2 * 3) + (4 * (5 + 6))", 51},
	{"1 + 2 * 3 + 4 * 5 + 6", 71},
	{"5 + (8 * 3 + 9 + 3 * 4 * 3)", 437},
	{"5 * 9 * (7 * 3 * 3 + 9 * 3 + (8 + 6 * 4))", 12240},
	{"((2 + 4 * 9) * (6 + 9 * 8 + 6) + 6) + 2 + 4 * 2", 13632},
}
func TestCalculatorOldPrecedence(t *testing.T) {
	for _, td := range TestCalculatorOldPrecedenceData {
		t.Run(td.in, func(t *testing.T) {
			r :=  calculate([]string{td.in},OldPrecedenceRules)
			if r != td.out {
				t.Errorf("got %d, expected %d", r, td.out)
			}
		})
	}
}

var TestCalculatorNewPrecedenceData = []struct {
	in string
	out int
}{
	{"1 + (2 * 3) + (4 * (5 + 6))", 51},
	{"1 + 2 * 3 + 4 * 5 + 6", 231},
	{"5 + (8 * 3 + 9 + 3 * 4 * 3)", 1445},
	{"5 * 9 * (7 * 3 * 3 + 9 * 3 + (8 + 6 * 4))", 669060},
	{"((2 + 4 * 9) * (6 + 9 * 8 + 6) + 6) + 2 + 4 * 2", 23340},
}
func TestCalculatorNewPrecedence(t *testing.T) {
	for _, td := range TestCalculatorNewPrecedenceData {
		t.Run(td.in, func(t *testing.T) {
			r :=  calculate([]string{td.in},NewPrecedenceRules)
			if r != td.out {
				t.Errorf("got %d, expected %d", r, td.out)
			}
		})
	}
}
