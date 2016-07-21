package main

import (
	"fmt"
	"os"
	"strconv"
)

func FactorCnt(n int64) int64 {
	var cnt int64
	var e int64
	cnt = 0

	for e = 1; e <= n; e++ {
		if n%e == 0 {
			cnt++
		}
	}

	return cnt
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	var max int64
	var e int64
	var maxdiv int64
	var fac int64

	if len(os.Args) < 2 {
		fmt.Printf("Description: %v calculates numbers that are highly divisible up the number passed on the commandline.\n", os.Args[0])
		fmt.Printf("Usage: %v <Max_Number>\n", os.Args[0])
		return
	}
	max, err := strconv.ParseInt(os.Args[1], 10, 64)
	check(err)
	maxdiv = 0
	fmt.Println("working with: ", max)

	// Check factory count for number 1 - max.
	for e = 1; e <= max; e++ {
		fac = FactorCnt(e)
		if maxdiv < fac {
			maxdiv = fac
			fmt.Printf("%v has %v divisors.\n", e, FactorCnt(e))
		}
	}

}
