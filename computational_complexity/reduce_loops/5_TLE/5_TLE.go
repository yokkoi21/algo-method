package main

import (
	"fmt"
)

func main() {
	var N int
	var tmp int
	var A []int
	var sum int

	fmt.Scan(&N)

	for i := 0; i < N; i++ {
		fmt.Scan(&tmp)
		A = append(A, tmp)
	}

	for i := 0; i < N; i++ {
		for j := i + 1; j < N; j++ {
			sum += A[i] * A[j]
		}
	}

	fmt.Println(sum)
}
