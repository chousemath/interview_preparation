package main

import (
	"fmt"
)

func main() {
	num := 600851475143
	// fmt.Println(sqrt)
	var maxPrime int
	for i := 3; i <= num; i += 2 {
		for num%i == 0 {
			num /= i
			maxPrime = i
		}
		// fmt.Println("maxPrime:", maxPrime)
	}
	fmt.Println("Answer:", maxPrime)
}
