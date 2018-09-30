package main

import (
	"fmt"
)

// time complexity is O(sqrt(n))
// space complexity is 1
func main() {
	num := 600851475143
	var maxPrime int
	for i := 3; i <= num; i += 2 {
		for num%i == 0 {
			num /= i
			maxPrime = i
		}
	}
	fmt.Println("Answer:", maxPrime)
}
