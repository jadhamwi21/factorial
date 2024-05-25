package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

// Factorial calculates the factorial of a non-negative integer n.
func Factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * Factorial(n-1)
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Please provide a number as an argument")
	}
	n, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatal("Invalid number")
	}
	fmt.Printf("Factorial of %d is %d\n", n, Factorial(n))
}
