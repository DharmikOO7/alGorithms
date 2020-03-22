package main

/*
	2.1-4
	Consider the problem of adding two n-bit binary integers, stored in two n-element
	arrays A and B. The sum of the two integers should be stored in binary form in
	an n + 1 -element array C. State the problem formally and write pseudocode for
	adding the two integers.
*/
import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var n int
	fmt.Print("Enter the number of bits: ")
	fmt.Scan(&n)
	rand.Seed(time.Now().UnixNano())
	fmt.Println("Generating two random binary numbers...")
	b1 := randomBinary(n)
	fmt.Printf("Number 1 is  %v, in decimal %d: \n", b1, binToInt(b1))
	b2 := randomBinary(n)
	fmt.Printf("Number 2 is  %v, in decimal %d: \n", b2, binToInt(b2))
	sum := addBinary(b1, b2)
	fmt.Printf("The sum is %v, in decimal %d: \n", sum, binToInt(sum))
}

func addBinary(b1, b2 []byte) []byte {
	n := len(b1)
	res := make([]byte, n+1)
	carry := byte(0)
	for i := n - 1; i > -1; i-- {
		res[i+1] = (b1[i] + b2[i] + carry) % 2
		carry = (b1[i] + b2[i] + carry) / 2
	}
	res[0] = carry
	return res
}

func binToInt(b []byte) int {
	n := len(b)
	decnum := 0
	for i := 0; i < n; i++ {
		decnum += int(b[n-i-1]) * (1 << i)
	}
	return decnum
}

func randomBinary(n int) []byte {
	b := make([]byte, n)
	for i := 0; i < n; i++ {
		b[i] = byte(rand.Intn(2))
	}
	return b
}
