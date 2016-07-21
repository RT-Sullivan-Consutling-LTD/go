package main

import (
	"fmt"
	"math/big"
	"os"
)

func main() {
	n := big.NewInt(0)
	r := big.NewInt(0)
	var cnt int64

	if len(os.Args) < 2 {
		fmt.Printf("Usage: %v <Number>\n", os.Args[0])
		return
	}

	n.SetString(os.Args[1], 10)

	for e := big.NewInt(1); e.Cmp(n) <= 0; e.Add(e, big.NewInt(1)) {
		r.Mod(n, e)
		if r.Cmp(big.NewInt(0)) == 0 {
			if e.ProbablyPrime(20) {
				fmt.Printf("[%v] ", e)
			} else {
				fmt.Printf("%v ", e)
			}
			cnt++
		}
	}
	fmt.Printf("\nNumber of divisors: %v\n", cnt)
}
