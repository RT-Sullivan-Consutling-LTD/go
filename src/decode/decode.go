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
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"math/big"
	"os"
	"strconv"
)

// base encoding (number of ascii characters)
const messageBase = 256
const MaxMsgLen = 256 - 3

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	// Private Key Variables
	e, n := big.NewInt(0), big.NewInt(0)
	// Public Key Variables
	d, p, q := big.NewInt(0), big.NewInt(0), big.NewInt(0)

	var err error

	if len(os.Args) < 3 {
		fmt.Printf("Usage: %v <Key_Prefix> <Message_File>\n", os.Args[0])
		return
	}
	// Take key prefix from command line
	keyname := fmt.Sprintf("%v.private_keys.info", os.Args[1])
	fmt.Printf("Using %v\n", keyname)

	// Take Message File Name from command line
	fname := os.Args[2]
	fmt.Printf("Reading from %v\n", fname)

	c := big.NewInt(0)

	// ------------------------------------------
	// Read encrypted message from file.
	ctext, err := ioutil.ReadFile(fname)
	check(err)
	//fmt.Printf("Encrypted message from file: %v\n", ctext)
	c.SetString(string(ctext[3:]), 10)
	// ------------------------------------------

	// Get length of the message in the text
	msgLen, err := strconv.Atoi(string([]byte(ctext[0:3])))
	check(err)
	fmt.Println("mesg len: ", msgLen)

	// /////////////////////////////////////////
	// Read PUBLIC keys from file
	fpri, err := os.Open(keyname)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer fpri.Close()
	rpri := bufio.NewReaderSize(fpri, 4*1024)

	// Read the first line
	line, isPrefix, err := rpri.ReadLine()
	for err == nil && !isPrefix {
		s := string(line)
		switch 0 {
		case bytes.Compare(line[0:2], []byte("e:")):
			e.SetString(s[3:], 10)
		case bytes.Compare(line[0:2], []byte("n:")):
			n.SetString(s[3:], 10)
		case bytes.Compare(line[0:2], []byte("d:")):
			d.SetString(s[3:], 10)
		case bytes.Compare(line[0:2], []byte("p:")):
			p.SetString(s[3:], 10)
		case bytes.Compare(line[0:2], []byte("q:")):
			q.SetString(s[3:], 10)
		}

		// Read the next line
		line, isPrefix, err = rpri.ReadLine()
	}
	if isPrefix {
		fmt.Println("buffer size to small")
		return
	}

	// Decrypt the message
	// m = c^d mod n
	m := new(big.Int).Exp(c, d, n)
	//	fmt.Printf("Decrypted Message: %v\n\n", m)

	// Decode the Base 256 message number
	exp := MaxMsgLen - 1
	tmp := big.NewInt(0)
	fmt.Printf("Decoded Message: ") //
	for cnt := 0; cnt < msgLen; cnt++ {
		tmp.Div(m, new(big.Int).Exp(big.NewInt(messageBase), big.NewInt(int64(exp)), nil))
		fmt.Printf("%c", tmp.Int64()) //
		m.Mod(m, new(big.Int).Exp(big.NewInt(messageBase), big.NewInt(int64(exp)), nil))
		exp--
	}
	fmt.Printf("\nEND\n")

}
