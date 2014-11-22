package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func loadDigits(filename string) (digits []int) {
	file, err := ioutil.ReadFile("data")
	if err != nil {
		panic(err)
	}
	newString := string(file)
	trimmedString := strings.TrimSpace(newString)
	stringDigits := strings.Split(trimmedString, "")
	digits = digitsToInts(stringDigits)
	return
}

func digitsToInts(strDigs []string) (digits []int) {
	digits = []int{}
	for i := 0; i < len(strDigs); i++ {
		number, err := strconv.ParseInt(strDigs[i], 0, 32)
		if err != nil {
			panic(err)
		}
		digits = append(digits, int(number))
	}
	return
}

func computeGreatestProduct(digits []int, number int) (value int) {
	value = 0

	for i := number - 1; i < len(digits); i++ {
		product := digits[i]
		for j := i - 1; j > i-number; j-- {
			product = product * digits[j]
		}
		if product > value {
			value = product
		}
	}

	return
}

func main() {
	digits := loadDigits("data")
	result := computeGreatestProduct(digits, 13)
	fmt.Println(result)
}
