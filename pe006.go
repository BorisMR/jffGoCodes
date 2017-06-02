package main

import (
	"math"
	"fmt"
)

func main(){
	sumSqre := 0
	sqSum := 0
	const limit = 100

	for i := 1; i <= limit; i++ {
		sqSum += i
		sumSqre += i*i
	}

	sqSum = int(math.Pow(float64(sqSum), float64(2)))
	fmt.Printf("Square of sum - sum of squares: %d",	sqSum-sumSqre)
}

/*
The sum of the squares of the first ten natural numbers is,
1^2 + 2^2 + ... + 10^2 = 385

The square of the sum of the first ten natural numbers is,
(1 + 2 + ... + 10)^2 = 55^2 = 3025

Hence the difference between the sum of the squares of the
first ten natural numbers and the square of the sum is 3025 − 385 = 2640.

Find the difference between the sum of the squares of the
first one hundred natural numbers and the square of the sum.
 */