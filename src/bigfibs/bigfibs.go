package main

import (
	//	"os"
	"fmt"
	"math/big"
)

const Num_of_digits = 1000

func main() {

	// argsWithProg := os.Args
	//    argsWithoutProg := os.Args[1:]

	// Initialize two big ints with the first two numbers in the sequence.
	a := big.NewInt(0)
	b := big.NewInt(1)

	// Initialize limit as 10^999, the smallest integer with 100 digits.
	var limit big.Int
	limit.Exp(big.NewInt(10), big.NewInt(Num_of_digits-1), nil)

	// Loop while a is smaller than 1e100.
	for a.Cmp(&limit) < 0 {
		// Compute the next Fibonacci number, storing it in a.
		a.Add(a, b)
		// Swap a and b so that b is the next number in the sequence.
		a, b = b, a
	}

	fmt.Println(a) // Really big Fibonacci number

	// Test a for primality.
	// (ProbablyPrimes' argument sets the number of Miller-Rabin
	// rounds to be performed. 20 is a good value.)
	//    fmt.Print("\nProbability of being Prime: ")
	// fmt.Println(a.ProbablyPrime(20))

}
