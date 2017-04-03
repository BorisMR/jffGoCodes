package main

import (
	"fmt"
)

func main() {
	var ant, fib, sum int
	suc := 1

	for fib < 4000000 {
		fib = ant + suc
		ant = suc;
		suc = fib;
		if fib%2 == 0 {
			sum += fib
			fmt.Printf("%d\n", fib)
		}
	}

	fmt.Printf("\nSum of even val of fibonacci: %d", sum)
}

/*
Each new term in the Fibonacci sequence is generated
by adding the previous two terms. By starting with 1 and 2,
the first 10 terms will be:

1, 2, 3, 5, 8, 13, 21, 34, 55, 89, ...

By considering the terms in the Fibonacci sequence whose
values do not exceed four million, find the
sum of the even-valued terms.
*/
