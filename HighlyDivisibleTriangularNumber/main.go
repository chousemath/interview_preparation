package main

import (
	"fmt"
	"math"
)

func factorsAbove500(n int) bool {
	hm := map[int]bool{}
	hm[1] = true
	hm[n] = true
	if n > 1 {
		count := 2
		for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
			if val := hm[i]; val {
				// for k := range hm {
				// 	fmt.Println(k)
				// }
				return false
			}
			if n%i == 0 {
				count += 2
				if count >= 500 {
					return true
				}
				hm[i] = true
				hm[n/i] = true
			}
		}

	}
	// for k := range hm {
	// 	fmt.Println(k)
	// }
	return false
}

func main() {
	fmt.Println("v0.0.1")

	counter := 0
	triangular := 0

	for {
		counter++
		triangular += counter
		if factorsAbove500(triangular) {
			fmt.Println(triangular)
			break
		}
	}
}
