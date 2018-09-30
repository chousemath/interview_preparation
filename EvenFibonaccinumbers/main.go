package main

import "fmt"

func main() {
	fmt.Println("v0.0.1")
	prev := 1
	curr := 2
	temp := 0
	sum := 2
	for prev+curr < 4000000 {
		temp = prev
		prev = curr
		curr += temp
		if curr%2 == 0 {
			sum += curr
		}
	}
	fmt.Println("Answer:", sum)
}
