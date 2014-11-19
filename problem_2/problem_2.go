package main

import "fmt"

func genFibUntil(n int) (fibSlice []int) {
	fibSlice = []int{1, 2}

	i := 2

	for {
		tempValue := fibSlice[i-2] + fibSlice[i-1]

		if tempValue > n {
			return
		}

		fibSlice = append(fibSlice, tempValue)
		i++
	}
}

func getEven(slice []int) (newSlice []int) {
	newSlice = []int{}

	for i := 0; i < len(slice); i++ {
		if (slice[i] % 2) == 0 {
			newSlice = append(newSlice, slice[i])
		}
	}
	return
}

func sumFib(slice []int) (value int) {
	for i := 0; i < len(slice); i++ {
		value += slice[i]
	}
	return
}

func main() {
	values := genFibUntil(4000000)
	evens := getEven(values)
	result := sumFib(evens)
	fmt.Println("Solution: ", result)
}
