package main

import "fmt"

func isPalindrome(n int) bool {
	if n < 0 {
		return false
	}
	iSlice := []int{}
	for n > 0 {
		iSlice = append(iSlice, n%10)
		n /= 10
	}
	limit := len(iSlice) / 2
	for i := 0; i < limit; i++ {
		if iSlice[i] != iSlice[len(iSlice)-1-i] {
			return false
		}
	}
	return true
}

func main() {
	var maxPalindrome int
	var product int
	for i := 999; i > 99; i-- {
		for j := 999; j > 99; j-- {
			product = i * j
			if product > maxPalindrome && isPalindrome(product) {
				maxPalindrome = product
			}
		}
	}
	fmt.Println("Answer:", maxPalindrome)
}
