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

func buildSquareOfSum(values []int) (result int) {
	result = 0

	for i := range values {
		result = result + values[i]
	}

	result = result * result
	return
}

func buildSumOfSquares(values []int) (result int) {
	result = 0

	for i := range values {
		result = result + (values[i] * values[i])
	}
	return
}

func main() {
	workingRange := genRange(1, 100)
	squareOfSum := buildSquareOfSum(workingRange)
	sumOfSquares := buildSumOfSquares(workingRange)
	result := squareOfSum - sumOfSquares
	fmt.Println("Result: ", result)
}
