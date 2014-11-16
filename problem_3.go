package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	file, err := ioutil.ReadFile("fixtures/primes.txt")
	if err != nil {
		panic(err)
	}

	// slice := []string{}
	slice := strings.Split(string(file), "\r\n")

	// fmt.Println(slice)
	// fmt.Println("First Line: ", slice[0])
	// fmt.Println("Second Line: ", slice[1])
	// fmt.Println("Third Line: ", slice[2])
	// fmt.Println("Fourth Line: ", slice[3])
	// fmt.Println("Fifth Line: ", strings.Trim(slice[4], " "))
	// fmt.Println("Last Line: ", slice[len(slice)-2])

	slice = slice[4:]
	slice = slice[:len(slice)-2]

	for i := 0; i < len(slice); i++ {
		line := strings.Trim(slice[i], " ")
		fmt.Println(strings.Split(line, " "))
	}
}
