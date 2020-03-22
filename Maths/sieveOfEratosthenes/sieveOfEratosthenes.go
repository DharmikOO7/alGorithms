package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	fmt.Print("Enter the range till which to find primes: ")
	fmt.Scan(&n)
	a := make([]bool, n+1)
	a[0] = false
	a[1] = false
	a[2] = true
	var primes []int
	//Initialize array
	for i := 3; i < len(a); i += 2 {
		a[i] = true
	}
	nsqrt := math.Sqrt(float64(n))
	for i := 3; i <= int(nsqrt+1); i++ {
		if a[i] {
			//Set all multiples of current non-zero element to 0
			for j := i * i; j < len(a); j += i {
				a[j] = false
			}
		}
	}
	//Filter all primes and store them in an array
	for i, v := range a {
		if v {
			primes = append(primes, i)
		}
	}
	fmt.Printf("Primes between 2 and %d: %v", n, primes)
}
