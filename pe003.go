package main

import "fmt"

func main()  {
	largestFact := 2
	valToFound := 600851475143
	divValue := valToFound
	factTotal := 1
	cont := 2

	for cont <= divValue {
		if divValue % cont == 0 {
			largestFact = cont
			fmt.Printf("factor->%d\n", largestFact)
			divValue /= cont
			factTotal *= cont
			if factTotal == valToFound {
				fmt.Printf("valToFound=>%d\n",factTotal)
				break
			}
		}else {
			cont++
		}
	}
	fmt.Printf("Largest Prime Factor is %d", largestFact)
}

/*
The prime factors of 13195 are 5, 7, 13 and 29.
What is the largest prime factor of the number 600851475143 ?
*/