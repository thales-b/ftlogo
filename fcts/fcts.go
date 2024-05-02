package fcts

import (
	"fmt"
)

func Withdraw(balance, amount int) (int, error) {
	if amount > balance {
		return 0, fmt.Errorf("withdrawal amount %d superior to current balance %d", amount, balance)
	}
	return balance - amount, nil
}

func Apply(x int, f func(int) int) int {
	return f(x)
}

func AddMany(inputs ...float64) float64 {
	var result float64
	for _, input := range inputs {
		result += input
	}
	return result
}

func ComputeMany(f func(float64, float64) float64, inputs ...float64) float64 {
	result := inputs[0]
	for i := 1; i < len(inputs); i++ {
		result = f(result, inputs[i])
	}
	return result
}
