// Given an array of int32egers representing the color of each sock,
// determine how many pairs of socks with matching colors there are.
package main

import "fmt"

func main() {
	socks := []int32{1, 2, 1, 2, 1, 3, 2}
	sockMap := make(map[int32]int32)
	// O(n)
	for _, sock := range socks {
		sockMap[sock]++
	}
	var sum int32
	// O(n) in the worst case, e.g. []int32{1,2,3,4,5}
	for _, v := range sockMap {
		sum += v / 2
	}
	fmt.Println(sum)
	// worst case this would be 2*O(n) => O(n)
}
