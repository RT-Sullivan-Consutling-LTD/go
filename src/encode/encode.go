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
	"math/rand"
	"os"
	"time"
)

// base encoding (number of ascii characters)
const messageBase = 256
const MaxMsgLen = 256 - 3

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func RandomString(strlen int) []byte {
	rand.Seed(time.Now().UTC().UnixNano())
	const chars = " abcdefghijklmnopqrstuvwxyz0123456789"
	result := make([]byte, strlen)
	for i := 0; i < strlen; i++ {
		result[i] = chars[rand.Intn(len(chars))]
	}
	return result
}

func main() {
	// Private Key Variables
	e, n := big.NewInt(0), big.NewInt(0)
	var err error

	mtext := RandomString(MaxMsgLen)
	if len(os.Args) < 3 {
		fmt.Printf("Usage: %v <Key_Prefix> <Input_File> <optional: Output_File>\n", os.Args[0])
		return
	}
	// Take key prefix from command line
	keyname := fmt.Sprintf("%v.public_keys.info", os.Args[1])
	fmt.Printf("Using %v\n", keyname)

	// Take Message File Name from command line
	fname := os.Args[2]
	fmt.Printf("Input file %v\n", fname)

	outputfilename := fmt.Sprintf("%v.crypto", fname)
	if len(os.Args) == 4 {
		// Take Message File Name from command line
		outputfilename := os.Args[3]
		fmt.Printf("Writing code to %v\n", outputfilename)
	}

	mfiletxt, err := ioutil.ReadFile(fname)
	check(err)
	//	fmt.Printf("file contents: %v\n", mfiletxt)

	// Copy the message into the random string
	copy(mtext[:], mfiletxt)

	// Encode the data string
	tmp := big.NewInt(0)
	m := big.NewInt(0)
	exp := len(mtext) - 1
	for cnt := 0; cnt < len(mtext); cnt++ {
		tmp.Mul(big.NewInt(int64(mtext[cnt])), new(big.Int).Exp(big.NewInt(messageBase), big.NewInt(int64(exp)), nil))
		exp--
		m.Add(m, tmp)
	}
	//fmt.Println("Message in Base256: ", m)

	// /////////////////////////////////////////
	// Read PUBLIC keys from file
	fpub, err := os.Open(keyname)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer fpub.Close()
	r := bufio.NewReaderSize(fpub, 4*1024)

	// Read the first line
	line, isPrefix, err := r.ReadLine()
	for err == nil && !isPrefix {
		s := string(line)
		switch 0 {
		case bytes.Compare(line[0:2], []byte("e:")):
			e.SetString(s[3:], 10)
		case bytes.Compare(line[0:2], []byte("n:")):
			n.SetString(s[3:], 10)
		}

		// Read the next line
		line, isPrefix, err = r.ReadLine()
	}
	if isPrefix {
		fmt.Println("buffer size to small")
		return
	}

	// Make sure the message isn't too big to handle
	if m.Cmp(n) >= 0 {
		fmt.Println("Message too long.")
		return
	}

	// Encrypt the message
	// c = m^e mod n
	c := new(big.Int).Exp(m, e, n)
	//fmt.Printf("\nEncrypted Message: %v\n", c)

	// ------------------------------------------
	// Write the encrypted data to a file.
	codedMsg := fmt.Sprintf("%03d%v", len(mfiletxt), c)
	ioutil.WriteFile(outputfilename, []byte(codedMsg), 0644)
	// ------------------------------------------
	fmt.Println("END - Message Encrypted")

}
