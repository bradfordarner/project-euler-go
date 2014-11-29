package main

import (
	"fmt"
)

func genPrimes(limit int) (primes []int) {
	primes = []int{}
	primes = append(primes, 2)
	for i := 3; i < limit; i++ {
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

func sumPrimes(primes []int) (value int) {
	for i := range primes {
		value += primes[i]
	}
	return
}

func main() {
	primes := genPrimes(2000000)
	result := sumPrimes(primes)
	fmt.Println("Result: ", result)
}
