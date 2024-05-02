package fcts_test

import (
	"fcts"
	"testing"
)

func TestWithdrawValid(t *testing.T) {
	t.Parallel()
	balance := 100
	amount := 20
	want := 80
	got, err := fcts.Withdraw(balance, amount)
	if err != nil {
		t.Fatal(err)
	}
	if want != got {
		t.Errorf("want %d, got %d", want, got)
	}
}
func TestWithdrawInvalid(t *testing.T) {
	t.Parallel()
	balance := 20
	amount := 100
	_, err := fcts.Withdraw(balance, amount)
	if err == nil {
		t.Fatal("want error for invalid withdrawal, got nil")
	}
}

func TestApply(t *testing.T) {
	t.Parallel()
	want := 2
	got := fcts.Apply(1, func(x int) int {
		return x * 2
	})
	if want != got {
		t.Fatalf("want %d, got %d", want, got)
	}
}

func add(x, y float64) float64 {
	return x + y
}

func sub(x, y float64) float64 {
	return x - y
}

func mul(x, y float64) float64 {
	return x * y
}

func div(x, y float64) float64 {
	return x / y
}

func TestComputeMany(t *testing.T) {
	t.Parallel()
	input := []float64{12, 4, 3}
	ops := []func(float64, float64) float64{
		add,
		sub,
		mul,
		div,
	}
	want := []float64{19, 5, 144, 1}
	for i, op := range ops {
		got := fcts.ComputeMany(op, input...)
		if want[i] != got {
			t.Fatalf("fct %T, input %f, want %f, got %f", op, input, want[i], got)
		}
	}
}
