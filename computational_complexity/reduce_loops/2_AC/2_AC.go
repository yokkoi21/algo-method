package main

import (
	"fmt"
)

func main() {
	var N int
	var tmp int
	var A []int
	var sum int
	var max int = 0

	fmt.Scan(&N)

	for i := 0; i < N; i++ {
		fmt.Scan(&tmp)
		A = append(A, tmp)
	}

	for i := 0; i < N; i++ {
		sum += A[i]
	}

	for i := 0; i < N; i++ {
		if sum-A[i] > max {
			max = sum - A[i]
		}
	}

	fmt.Println(max)
}
