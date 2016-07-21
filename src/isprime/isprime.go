package main

import (
	"fmt"
	"math/big"
	"os"
)

func main() {
	n := big.NewInt(0)

	if len(os.Args) < 2 {
		fmt.Printf("Usage: %v <Number>\n", os.Args[0])
		return
	}

	n.SetString(os.Args[1], 10)

	fmt.Println(n.ProbablyPrime(20))
}
