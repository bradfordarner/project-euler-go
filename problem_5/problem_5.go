package main

import "fmt"

func genRange(start, end int) (values []int) {
	values = []int{}
	i := start
	for {
		values = append(values, i)
		if i == end {
			return
		}
		i++
	}
}

func maxEvenlyDiv(testRange []int) (value int) {
	value = 1

	for {
		completed := true
		for i := range testRange {
			if value%testRange[i] != 0 {
				completed = false
				break
			}
		}
		if completed {
			return
		}
		value++
	}
}

func main() {
	testRange := genRange(1, 20)
	result := maxEvenlyDiv(testRange)
	fmt.Println("Range: ", testRange)
	fmt.Println("Results: ", result)
}
