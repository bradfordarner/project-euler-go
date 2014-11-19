package main

import (
	"fmt"
)

func computeSum(n int) (value int) {
	value = 0
	for i := 0; i < n; i++ {
		if i%3 == 0 || i%5 == 0 {
			value += i
		}
	}
	return
}

func main() {
	val := computeSum(1000)
	fmt.Println("Solution: ", val)
}
