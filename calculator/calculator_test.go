package calculator_test

import (
	"calculator"
	"testing"
)

func TestAdd(t *testing.T) {
	type testCase struct {
		a, b float64
		want float64
	}
	testCases := []testCase{
		{a: 2, b: 2, want: 4},
		{a: 1, b: 1, want: 2},
		{a: 5, b: 0, want: 5},
	}
	for _, tc := range testCases {
		got := calculator.Add(tc.a, tc.b)
		if tc.want != got {
			t.Errorf("Add(%f, %f): want %f, got %f", tc.a, tc.b,
				tc.want, got)
		}
	}
}

func TestSubtract(t *testing.T) {
	type testcase struct {
		a, b float64
		want float64
	}
	testcases := []testcase{
		{a: 4, b: 2, want: 2},
		{a: 1, b: 2, want: -1},
		{a: 5, b: 0, want: 5},
	}
	for _, tc := range testcases {
		got := calculator.Subtract(tc.a, tc.b)
		if tc.want != got {
			t.Errorf("Subtract(%f, %f): want %f, got %f", tc.a, tc.b,
				tc.want, got)
		}
	}
}

func TestMultiply(t *testing.T) {
	type testcase struct {
		a, b float64
		want float64
	}
	testcases := []testcase{
		{a: 3, b: 3, want: 9},
		{a: 1, b: -2, want: -2},
		{a: 5, b: 0, want: 0},
	}
	for _, tc := range testcases {
		got := calculator.Multiply(tc.a, tc.b)
		if tc.want != got {
			t.Errorf("Multiply(%f, %f): want %f, got %f", tc.a, tc.b,
				tc.want, got)
		}
	}
}
