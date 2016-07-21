package main

import (
	//	"os"

	"fmt"
	"math/big"
)

func main() {

	MillsConstant := big.NewInt(0)
	exp := big.NewInt(0)
	exp.Div(big.NewInt(1), big.NewInt(9))
	fmt.Println(exp)

	MillsConstant.Exp(big.NewInt(11), exp, nil)

	//	n := big.NewInt(1)

	//startloop:
	fmt.Println(MillsConstant)

}
