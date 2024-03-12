package main

import (
	"fmt"
	"math"
)

func main() {
	var P, V, Q, M int
	fmt.Scan(&P, &V)
	fmt.Scan(&Q, &M)

	distance := func(a, b int) float64 {
		return math.Abs(float64(a - b))
	}

	count := 0
	for i := 1; i <= 100; i++ {
		if distance(P, i) <= float64(V) && distance(Q, i) <= float64(M) {
			count++
		}
	}

	fmt.Println(count)
}
