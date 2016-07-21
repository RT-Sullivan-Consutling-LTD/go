/*
n = pq, where p and q are distinct primes.
phi, φ = (p-1)(q-1)
e < n such that gcd(e, phi)=1
d = e-1 mod phi.
c = me mod n, 1<m<n.
m = cd mod n.
*/

package main

import (
	"crypto/rand"
	"os"
	//	"encoding/hex"
	//	"encoding/binary"
	"fmt"
	"math/big"
)

// Initialize key size
const k = 2048 // k = 1024, 2048, 3072, 4096,...

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	p_1 := big.NewInt(0) // will hold the value of p-1
	q_1 := big.NewInt(0) // will hold the value of q-1
	n := big.NewInt(0)
	phi := big.NewInt(0)
	d := big.NewInt(0)
	var err error

	if len(os.Args) < 2 {
		fmt.Printf("Usage: %v <Key_Prefix>\n", os.Args[0])
		return
	}
	// Take key prefix from command line
	privatefilename := fmt.Sprintf("%v.private_keys.info", os.Args[1])
	publicfilename := fmt.Sprintf("%v.public_keys.info", os.Args[1])
	fmt.Printf("Using %v, %v\n", privatefilename, publicfilename)

	// /////////////////////////////////////////
	// Generate base pair of Prime numbers p & q
generate_p:
	p := big.NewInt(0)
	p, err = rand.Prime(rand.Reader, k/2)
	if err != nil {
		fmt.Println(err)
		goto generate_p
	}

generate_q:
	q := big.NewInt(0)
	q, err = rand.Prime(rand.Reader, k/2)
	if err != nil {
		fmt.Println(err)
		goto generate_q
	}

	// swap if p < q
	if p.Cmp(q) < 0 {
		p, q = q, p
	}

	// /////////////////////////////////////////
	// Common selections for e are 3, 5, 17, 257 or 65537
	e := big.NewInt(65537)

	// /////////////////////////////////////////
	// Generate n = pq, phi = (p-1)(q-1)
	n.Mul(p, q)
	phi.Mul(p_1.Sub(p, big.NewInt(1)), q_1.Sub(q, big.NewInt(1)))
	// Test to see that e and phi have no common factors. If so, regenerate primes again.
	if d.ModInverse(e, phi) == nil {
		goto generate_p
	}

	// ////////////////////////////////////////////
	// Save key results to file.
	f, err := os.Create(publicfilename)
	check(err)
	defer f.Close()

	output := []byte(fmt.Sprintf("[Public Key Components]\n")) //
	f.Write(output)
	output = []byte(fmt.Sprintf("e: %v\nn: %v\n", e, n)) //
	f.Write(output)

	fp, err := os.Create(privatefilename)
	check(err)
	defer fp.Close()

	output = []byte(fmt.Sprintf("[Private Key Components]\n")) //
	fp.Write(output)
	output = []byte(fmt.Sprintf("d: %v\n", d)) //
	fp.Write(output)
	output = []byte(fmt.Sprintf("p: %v\n", p)) //
	fp.Write(output)
	output = []byte(fmt.Sprintf("q: %v\n", q)) //
	fp.Write(output)
	output = []byte(fmt.Sprintf("n: %v\n", n)) //
	fp.Write(output)

}
