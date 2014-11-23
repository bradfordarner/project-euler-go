package main

import "fmt"

type Triplet struct {
	A int
	B int
	C int
}

func tripletSum(goal int, triplets []Triplet) (result Triplet) {
	result = Triplet{}
	for i := 0; i < len(triplets); i++ {
		triplet := triplets[i]
		if triplet.A+triplet.B+triplet.C == goal {
			result = triplets[i]
		}
	}
	return
}

func tripletProduct(triplet Triplet) (result int) {
	result = triplet.A * triplet.B * triplet.C
	return
}

func tripletCheck(a, b, c int) bool {
	a_squared := a * a
	b_squared := b * b
	c_squared := c * c

	return a_squared+b_squared == c_squared
}

func genTriplets(max int) (values []Triplet) {
	values = []Triplet{}

	for c := max; c > 0; c-- {
		for b := c - 1; b > 0; b-- {
			for a := b - 1; a > 0; a-- {
				if tripletCheck(a, b, c) {
					triplet := Triplet{a, b, c}
					values = append(values, triplet)
				}
			}
		}
	}

	return
}

func main() {
	triplets := genTriplets(1000)
	goalTriplet := tripletSum(1000, triplets)
	result := tripletProduct(goalTriplet)
	fmt.Println("Result: ", result)
}
