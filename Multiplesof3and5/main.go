package main

import (
	"fmt"
)

func main() {
	fmt.Println("v0.0.1")
	sum := 0
	for i := 0; i < 1000; i++ {
		if i%3 == 0 || i%5 == 0 {
			sum += i
		}
	}
	fmt.Println("Answer:", sum)
}