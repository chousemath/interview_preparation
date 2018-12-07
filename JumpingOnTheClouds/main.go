// https://www.hackerrank.com/challenges/jumping-on-the-clouds/problem?h_l=interview&playlist_slugs%5B%5D=interview-preparation-kit&playlist_slugs%5B%5D=warmup
package main

import "fmt"

func cloudCount(c []int32) int32 {
	var steps int32
	limit := len(c) - 1
	for i := 0; i < len(c); i++ {
		reach := i + 2
		if reach <= limit && c[reach] == 0 {
			i++
		}
		steps++
	}
	return steps - 1
}

func main() {
	fmt.Println("v0.0.1")
	fmt.Println(cloudCount([]int32{0, 1, 0, 0, 0, 1, 0}))
	fmt.Println(cloudCount([]int32{0, 0, 0, 0, 1, 0}))
}
