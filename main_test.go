package main

import (
	"testing"
)

func TestFactorial(t *testing.T) {
	tests := []struct {
		name     string
		input    int
		expected int
	}{
		{"Factorial of 0", 0, 1},
		{"Factorial of 1", 1, 1},
		{"Factorial of 2", 2, 2},
		{"Factorial of 3", 3, 6},
		{"Factorial of 4", 4, 24},
		{"Factorial of 5", 5, 120},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Factorial(tt.input)
			if result != tt.expected {
				t.Errorf("Factorial(%d) = %d; expected %d", tt.input, result, tt.expected)
			}
		})
	}
}
