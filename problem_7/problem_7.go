package main

import (
	"fmt"
)

func genPrimes(limit int) (primes []int) {
	primes = []int{}
	primes = append(primes, 2)
	for i := 3; len(primes) < limit; i++ {
		divisible := false
		for j := range primes {
			if i%primes[j] == 0 {
				divisible = true
				break
			}
		}
		if divisible != true {
			primes = append(primes, i)
		}
	}

	return
}

func main() {
	primes := genPrimes(10001)
	fmt.Println("Result: ", primes[len(primes)-1])
}
