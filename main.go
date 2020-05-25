package main

import (
	"fmt"

	"github.com/kpteja/go-project-euler/fibonacci"
	"github.com/kpteja/go-project-euler/multiples"
)

func main() {
	// Problem 1
	p1Sum := multiples.CalculateSumOfMultiples([]int{3, 5}, 1, 1000)
	fmt.Println(p1Sum)

	// Problem 2
	p2Sum := fibonacci.SumOfEvenValuedTerms(4000000)
	fmt.Println(p2Sum)
}
