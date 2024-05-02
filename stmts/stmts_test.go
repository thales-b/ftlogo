package stmts_test

import (
	"github.com/google/go-cmp/cmp"
	"stmts"
	"testing"
)

func TestBigger(t *testing.T) {
	t.Parallel()
	type TestCase struct {
		a, b int
		want int
	}
	testCases := []TestCase{
		{a: 5, b: 2, want: 5},
		{a: 4, b: 6, want: 6},
		{a: 3, b: 3, want: 3},
	}
	for _, tc := range testCases {
		got := stmts.Bigger(tc.a, tc.b)
		if tc.want != got {
			t.Errorf("a: %d, b: %d, want %d, got %d", tc.a, tc.b, tc.want, got)
		}
	}
}

func TestXor(t *testing.T) {
	t.Parallel()
	type TestCase struct {
		a, b bool
		want bool
	}
	testCases := []TestCase{
		{a: false, b: false, want: false},
		{a: true, b: false, want: true},
		{a: false, b: true, want: true},
		{a: true, b: true, want: false},
	}
	for _, tc := range testCases {
		got := stmts.Xor(tc.a, tc.b)
		if tc.want != got {
			t.Errorf("a: %t, b: %t, want %t, got %t", tc.a, tc.b, tc.want, got)
		}
	}
}

func TestGreet(t *testing.T) {
	t.Parallel()
	type TestCase struct {
		p    string
		want string
	}
	testCases := []TestCase{
		{p: "Alice", want: "Hey, Alice."},
		{p: "Bob", want: "What's up, Bob?"},
		{p: "abc", want: "Hello, stranger."},
	}
	for _, tc := range testCases {
		got := stmts.Greet(tc.p)
		if tc.want != got {
			t.Errorf("b: %q, want %q, got %q", tc.p, tc.want, got)
		}
	}
}

func TestTotal(t *testing.T) {
	t.Parallel()
	s := []int{1, 2, 3}
	want := 6
	got := stmts.Total(s)
	if want != got {
		t.Errorf("want %d, got %d", want, got)
	}
}

func TestEvens(t *testing.T) {
	t.Parallel()
	want := []int{0, 2, 4}
	var got []int
	result := stmts.Evens()
	for i := 0; i < 3; i++ {
		got = append(got, result[i])
	}
	if !cmp.Equal(want, got) {
		t.Errorf(cmp.Diff(want, got))
	}
}

func TestNonNegative(t *testing.T) {
	t.Parallel()
	s := []int{-1, 0, -2, 3}
	want := []int{0, 3}
	got := stmts.NonNegative(s)
	if !cmp.Equal(want, got) {
		t.Errorf(cmp.Diff(want, got))
	}
}
