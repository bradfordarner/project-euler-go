package main

import (
	"fmt"
	"strconv"
)

func genNums(start, end int) (numbers []int) {
	numbers = []int{}

	for i := start; i < end+1; i++ {
		numbers = append(numbers, i)
	}

	return
}

func genMultiples(numbers []int) (multiples []int) {
	multiples = []int{}

	for i := 0; i < len(numbers); i++ {
		for j := 0; j < len(numbers); j++ {
			multiples = append(multiples, numbers[i]*numbers[j])
		}
	}

	return
}

func reverseInt(number int64) (reverse int) {
	stringNumber := strconv.FormatInt(number, 10)

	runes := make([]rune, len(stringNumber))
	n := 0
	for _, r := range stringNumber {
		runes[n] = r
		n++
	}

	runes = runes[0:n]

	for i := 0; i < n/2; i++ {
		runes[i], runes[n-1-i] = runes[n-1-i], runes[i]
	}

	reverse64, _ := strconv.ParseInt(string(runes), 0, 0)

	reverse = int(reverse64)
	return
}

func findLargestPalindrome(numbers []int) (largest int) {
	largest = 0

	for i := 0; i < len(numbers); i++ {
		if (numbers[i] == reverseInt(int64(numbers[i]))) && (numbers[i] > largest) {
			largest = numbers[i]
		}
	}

	return
}

func main() {
	numbers := genNums(100, 999)
	multiples := genMultiples(numbers)
	result := findLargestPalindrome(multiples)

	// result := reverseInt(105)
	fmt.Println("Largest Palindrome: ", result)
}
