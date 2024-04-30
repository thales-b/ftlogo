// Package calculator does simple calculations.
package calculator

import (
	"errors"
	"math"
)

// Add takes two numbers and returns the result of adding
// them together.
func Add(a, b float64) float64 {
	return a + b
}

// Subtract takes two numbers a and b, and
// returns the result of subtracting b from a.
func Subtract(b, a float64) float64 {
	return b - a
}

// func Multiply(a, b float64) float64 {
// return 0
// }

// Multiply takes two numbers and returns the result of
// multiplying them together
func Multiply(a, b float64) float64 {
	return a * b
}

// Divide takes two numbers a and b, and returns the result of
// dividing a by b
func Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("division by zero not allowed")
	}
	return a / b, nil
}

// Sqrt takes a number and returns its square root
func Sqrt(a float64) (float64, error) {
	if a < 0 {
		return 0, errors.New("no square root for negative numbers")
	}
	return math.Sqrt(a), nil
}
