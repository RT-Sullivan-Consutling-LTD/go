package main

import (
	"fmt"
	"math/big"
	"os"
)

func main() {
	n := big.NewInt(0)
	w := big.NewInt(0)

	if len(os.Args) < 2 {
		fmt.Printf("Usage: %v <Number>\n", os.Args[0])
		return
	}

	n.SetString(os.Args[1], 10)

	fmt.Printf("Prime factors of %v are: ", n)
	div := big.NewInt(2)

	for n.Cmp(big.NewInt(1)) > 0 {
		w.Mod(n, div)
		if w.Cmp(big.NewInt(0)) != 0 && div.Cmp(n) <= 0 {
			//			fmt.Printf("%v %v   ", w, div)
			if div.Cmp(big.NewInt(2)) == 0 {
				div.Add(div, big.NewInt(1))
			} else {
				div.Add(div, big.NewInt(2))
			}
		} else {
			n.Div(n, div)
			fmt.Printf("[%v] ", div)
		}
	}
	fmt.Printf("\n")
}
