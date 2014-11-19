package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func loadPrimes(path string) (file []byte) {
	file, err := ioutil.ReadFile(path)

	if err != nil {
		panic(err)
	}

	return
}

func prepPrimeFile(data []byte) (slice []string) {
	slice = strings.Split(string(data), "\r\n")
	slice = slice[4:]
	slice = slice[:len(slice)-2]

	return
}

func createPrimeSlice(data []string) (slice []int) {
	slice = []int{}
	for line := range data {
		trimmedLine := strings.Trim(data[line], " ")
		numberArray := strings.Split(trimmedLine, " ")
		for number := range numberArray {
			switch numberArray[number] {
			case "":
				continue
			default:
				convNum, err := strconv.ParseInt(numberArray[number], 0, 0)

				if err != nil {
					panic(err)
				}

				slice = append(slice, int(convNum))
			}
		}
	}

	return
}

func calculateLargestFactor(num int, primes []int) (largest int) {
	largest = 0
	for prime := range primes {
		if (num%primes[prime] == 0) && (primes[prime] > largest) {
			largest = primes[prime]
		}
	}
	return
}

func main() {
	file := loadPrimes("fixtures/primes.txt")
	slice := prepPrimeFile(file)
	primes := createPrimeSlice(slice)

	result := calculateLargestFactor(600851475143, primes)
	fmt.Println("The largest prime: ", result)
}
