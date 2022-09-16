package main

import (
	"fmt"
)

func main() {
	var N int
	var tmp int
	var A []int
	var min int = 1000000000
	var max int = 0

	fmt.Scan(&N)

	for i := 0; i < N; i++ {
		fmt.Scan(&tmp)
		A = append(A, tmp)
	}

	for i := 0; i < N; i++ {
		if A[i] > max {
			max = A[i]
		}
		if A[i] < min {
			min = A[i]
		}
	}

	fmt.Println(max - min)
}
