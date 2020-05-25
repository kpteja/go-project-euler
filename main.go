package main

import (
	"fmt"

	"github.com/kpteja/go-project-euler/multiples"
)

func main() {
	// Problem 1
	p1Sum := multiples.CalculateSumOfMultiples([]int{3, 5}, 1, 1000)
	fmt.Println(p1Sum)
}
